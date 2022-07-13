package Routes

import (
	"customerproblem/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/shop")
	{
		grp1.GET("products", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.GET("transactions/:id", Controllers.GetTransactionsByID)

		grp1.POST("order", Controllers.CreateOrder)
		grp1.GET("order/:id", Controllers.GetOrderByID)
		grp1.GET("history/:id", Controllers.GetOrdersByID)

	}
	return r
}
