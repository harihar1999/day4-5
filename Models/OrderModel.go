package Models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	CustomerId uint   `json:"customer_id"`
	ProductId  uint   `json:"product_id"`
	Quantity   uint   `json:"quantity"`
	Status     string `json:"status"`
}

func (b *Order) TableName() string {
	return "order"
}
