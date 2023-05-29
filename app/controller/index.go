package controller

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (IndexController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Index successfully",
	})
}
