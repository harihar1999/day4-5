package Controllers

import (
	"customerproblem/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var order Models.Order
	//fmt.Println(user)
	c.BindJSON(&order)
	fmt.Println(order)
	err := Models.CreateOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)

	}
}

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//var output string
		c.JSON(http.StatusOK, order)
	}
}
