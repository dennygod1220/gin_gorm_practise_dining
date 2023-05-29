package router

import (
	"denny/gin_gorm_practise_dining/app/controller"

	"github.com/gin-gonic/gin"
)

type ControllerList struct {
	controller.IndexController
	controller.StoreBusinessTimeController
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	c := ControllerList{}

	r.GET("/", c.IndexController.Index)

	store := r.Group("/store")
	store.GET("/business_time", c.StoreBusinessTimeController.Index)

	return r
}
