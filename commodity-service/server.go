package main

import (
	"commodity-service/conf"
	"commodity-service/dao"
	"commodity-service/wire"
	"flag"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lgangkai/golog"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"protos/commodity"
)

type Server struct {
	service micro.Service
}

func (s *Server) Init() error {
	log.Println("Init server...")
	// 1. load config file.
	var confPath string
	flag.StringVar(&confPath, "config", "conf/commodity_live.yaml", "define config file")
	flag.Parse()
	config, err := conf.LoadConfig(confPath)
	if err != nil {
		log.Println("load config file error, err: ", err)
		return err
	}
	log.Println("config file loaded, config: ", config)

	mysqlMasterConf := config.MysqlMaster
	mysqlSlaveConf := config.MysqlSlave
	etcdConf := config.Etcd
	microConf := config.Micro
	redisConf := config.Redis

	// 2. register service.
	etcdReg := etcd.NewRegistry(
		registry.Addrs(etcdConf.Addr),
	)
	s.service = micro.NewService(
		micro.Name(microConf.Name),
		micro.Address(microConf.Addr),
		micro.Registry(etcdReg),
	)

	// 3. init basic dependencies.
	sqlMaster, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", mysqlMasterConf.Name,
		mysqlMasterConf.Password, mysqlMasterConf.Host, mysqlMasterConf.Port, mysqlMasterConf.DB)), &gorm.Config{})
	if err != nil {
		log.Println("init sqlDB master failed, err: ", err.Error())
	}

	sqlSlave, err := gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", mysqlSlaveConf.Name,
		mysqlSlaveConf.Password, mysqlSlaveConf.Host, mysqlSlaveConf.Port, mysqlSlaveConf.DB)), &gorm.Config{})
	if err != nil {
		log.Println("init sqlDB slave failed, err: ", err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisConf.Addr,
	})

	lgr := golog.Default()

	// 4. injection.
	commodityHandler := wire.InitCommodityHandler(
		&dao.DBMaster{DB: sqlMaster},
		&dao.DBSlave{DB: sqlSlave},
		rdb,
		lgr,
	)

	// 5. init service
	s.service.Init()
	err = commodity.RegisterCommodityHandler(s.service.Server(), commodityHandler)
	if err != nil {
		log.Println("register UserinfoHandler failed, err: ", err.Error())
		return err
	}

	return nil
}

func (s *Server) Run() error {
	if err := s.service.Run(); err != nil {
		log.Println("run server failed, err: ", err.Error())
		return err
	}
	return nil
}
