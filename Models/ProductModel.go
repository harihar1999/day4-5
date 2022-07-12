package Models

type Product struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}

func (b *Product) TableName() string {
	return "product"
}
