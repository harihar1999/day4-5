package Models

import (
	"customerproblem/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateOrder(order *Order, product *Product) (err error) {
	//fmt.Println(order)

	err = Config.DB.Create(order).Error
	if err != nil {
		return err
	}
    err = Config.DB.Where("id=?", order.ProductId).First(product).Error
	if err != nil {
		return err
	}
	fmt.Println(product)

	if product.Quantity < order.Quantity {
		order.Status="order failed"
		Config.DB.Save(order)
		//err = errors.New("order quantity is more than available")
		return nil
	}

	product.Quantity -= order.Quantity
	fmt.Println(product)
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