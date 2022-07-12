package Models

type Order struct {
	Id      	uint    `json:"id"`
	CustomerId	uint	`json:"customer_id"`
	ProductId 	uint   	`json:"product_id"`
	Quantity   	uint   	`json:"quantity"`
	Status    	string  `json:"status"`
}


func (b *Order) TableName() string {
	return "order"
}
