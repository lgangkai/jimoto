package main

import (
	"flag"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"jimotoapi/conf"
	"log"
	"protos/account"
	"protos/commodity"
)

type Server struct {
	Config          *conf.Config
	CommodityClient commodity.CommodityService
	AccountClient   account.AccountService
}

func (s *Server) Init() error {
	log.Println("Init server...")
	// 1. load config file.
	var confPath string
	flag.StringVar(&confPath, "config", "conf/jimotoapi_live.yaml", "define config file")
	flag.Parse()
	config, err := conf.LoadConfig(confPath)
	if err != nil {
		log.Println("load config file error, err: ", err)
		return err
	}
	log.Println("config file loaded, config: ", config)

	microConf := config.Micro
	etcdConf := config.Etcd

	s.Config = config

	// 2. init microservice client
	etcdReg := etcd.NewRegistry(
		registry.Addrs(etcdConf.Addr),
	)
	mService := micro.NewService(micro.Registry(etcdReg))
	mService.Init()
	s.CommodityClient = commodity.NewCommodityService(microConf.Commodity.Name, mService.Client())
	s.AccountClient = account.NewAccountService(microConf.Account.Name, mService.Client())

	return nil
}
