package Models

type Product struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	RetailerId  uint   `json:"retailer_id"`
}

func (b *Product) TableName() string {
	return "product"
}
