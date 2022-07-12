package Models

import (
	"customerproblem/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateOrder(order *Order) (err error) {
	fmt.Println(order)

	err = Config.DB.Create(order).Error
	//user.FirstName = "kfefe"
	//fmt.Println(order)
	if err != nil {
		return err
	}
	order.Status = "order placed"
	Config.DB.Save(order)
	return nil
}

func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
