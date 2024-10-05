package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lgangkai/golog"
	"jimotoapi/handler"
	"net/http"
)

func main() {
	server := &Server{}
	if err := server.Init(); err != nil {
		panic(err)
	}

	client := handler.NewClient(context.Background(), server.CommodityClient, server.AccountClient, server.Config, golog.Default())

	r := gin.Default()
	r.Use(client.GenRequestId, client.Cors)

	r.StaticFS("static", http.Dir("static"))
	api := r.Group("api")

	// init router
	accountApi := api.Group("account")
	{
		accountApi.POST("login", client.Login)
		accountApi.POST("logout", client.Logout)
		accountApi.POST("register", client.Register)
		accountApi.GET("/user/id", client.Authenticate, client.GetUserId)

		profileApi := accountApi.Group("profile", client.Authenticate)
		{
			profileApi.GET("", client.GetProfile)
			profileApi.POST("", client.CreateProfile)
			profileApi.PUT("", client.UpdateProfile)
			profileApi.DELETE("", client.DeleteProfile)
		}
	}

	commodityApi := api.Group("commodity")
	//commodityApi.Use(client.Authenticate)
	{
		commodityApi.GET(":commodity_id", client.GetCommodity)
		commodityApi.POST("", client.Authenticate, client.PublishCommodity)
		commodityApi.GET(":commodity_id/images", client.GetCommodityImages)
		commodityApi.GET(":commodity_id/liked-users", client.GetLikedUsers)

		likeApi := commodityApi.Group(":commodity_id/action")
		likeApi.Use(client.Authenticate)
		{
			likeApi.POST("like", client.LikeCommodity)
			likeApi.POST("unlike", client.UnlikeCommodity)
		}
	}

	commoditiesApi := api.Group("commodities")
	{
		commoditiesApi.GET("latest", client.GetLatestCommodityList)
		commoditiesApi.GET("liked", client.Authenticate, client.GetUserLikeCommodities)
	}

	publicServiceApi := api.Group("services")
	{
		publicServiceApi.POST("image/action/upload", client.UploadImage)
	}

	if err := r.Run(server.Config.Server.Addr); err != nil {
		panic(err)
	}
}
