package Models

import (
	"customerproblem/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
)

func CreateOrder(order *Order, product *Product) (err error) {
	//fmt.Println(order)

	userdata := map[string]string{
		"ram" : "ram123",
		"shyam" : "shyam1234",
	}

    r := http.Request{}

	username,password,ok := r.BasicAuth()
    fmt.Println(username,password,ok,"authentication")
	if ok {
		if userdata[username] == password {
			fmt.Println("auth done succesfully")
		}

		return nil

	}



	var lastorder Order
	Config.DB.Where("customer_id=?", order.CustomerId).Last(&lastorder)
	//fmt.Println(lastorder)


	err = Config.DB.Create(order).Error
	if err != nil {
		return err
	}
	err = Config.DB.Where("id=?", order.ProductId).First(product).Error
	if err != nil {
		return err
	}
	//fmt.Println(product)

	if time.Now().Sub(lastorder.CreatedAt).Minutes() < 5 {
		order.Status = "order failed due to time constraint"
		order.CreatedAt=lastorder.CreatedAt
		Config.DB.Save(order)
		return nil
	}
	if product.Quantity < order.Quantity {
		order.Status = "order failed due to less availability"
		order.CreatedAt=lastorder.CreatedAt
		Config.DB.Save(order)
		//err = errors.New("order quantity is more than available")
		return nil
	}

	product.Quantity -= order.Quantity
	//fmt.Println(product)
	err = UpdateProduct(product, string(1))

	//fmt.Println(order)
	if err != nil {
		return err
	}
	order.Status = "order placed"
	err = Config.DB.Save(order).Error
	return err
}

func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrdersByID(orders *[]Order, id string) (err error) {
	if err = Config.DB.Where("customer_id = ?", id).Find(orders).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactionsByID(orders *[]Order, id string) (err error) {

	var products []Product

	if err = Config.DB.Where("retailer_id = ?", id).Find(&products).Error; err != nil {
		return err
	}
	fmt.Println(products)

	return nil
}
