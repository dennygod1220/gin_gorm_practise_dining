package controller

import (
	"denny/gin_gorm_practise_dining/app/db"
	"denny/gin_gorm_practise_dining/app/query"
	"fmt"

	"github.com/gin-gonic/gin"
)

type StoreBusinessTimeController struct {
}

func (StoreBusinessTimeController) Index(c *gin.Context) {

	q := query.Use(db.DB)
	t, err := q.StoreBusinessTime.First()

	if err != nil {
		fmt.Printf("%#+v", err)
		panic("Query StorebusinessTime failed")
	}

	fmt.Printf("%+v\n", t)

	c.JSON(200, gin.H{
		"message": "StoreBusinessTimeController successfully",
	})
}
