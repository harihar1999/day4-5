package Controllers

import (
	"customerproblem/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type newdata struct {
	Id  uint
	ProductId uint
	Quantity uint
	Status string
}



func CreateOrder(c *gin.Context) {
	var order Models.Order
	var product Models.Product

	//fmt.Println(user)
	c.BindJSON(&order)
	fmt.Println(order)
	err := Models.CreateOrder(&order,&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		new := newdata{
			Id: order.Id,
			ProductId: order.ProductId,
			Quantity: order.Quantity,
			Status: order.Status,
		}
		fmt.Println(new)
		c.JSON(http.StatusOK, new)

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

func GetOrdersByID(c *gin.Context){
	id := c.Params.ByName("id")
	var orders []Models.Order
	err := Models.GetOrdersByID(&orders,id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//var output string
		c.JSON(http.StatusOK, orders)
	}
}
