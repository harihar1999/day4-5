package Models

import (
	"customerproblem/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Product) (err error) {
	fmt.Println(product)

	err = Config.DB.Create(product).Error
	//user.FirstName = "kfefe"
	fmt.Println(product)
	if  err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	//fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

//func DeleteUser(user *User, id string) (err error) {
//	Config.DB.Where("id = ?", id).Delete(user)
//	return nil
//}

