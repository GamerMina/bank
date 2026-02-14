package main

import (
	"bank/config"
	"bank/handler"
	"bank/repository"
	"bank/router"
	"bank/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	appConfig := config.LoadAppCfg()

	Db := config.DbCon(appConfig.Db)

	newRepository := repository.NewRepository(Db)
	newService := service.NewServices(newRepository)
	newHandler := handler.NewHandler(r, newService)

	router.InitRoutes(r, newHandler)

	r.Run("localhost:8081")
}
