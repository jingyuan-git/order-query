package models

type OrderItem struct {
	Id           int        `csv:"id"`
	OrderId      int        `csv:"order_id" `
	PricePerUnit int        `csv:"price_per_unit"`
	Quantity     int        `csv:"quantity"`
	Product      string     `csv:"product"`
	Delivery     []Delivery `gorm:"foreignKey:OrderItemId;references:Id"`
}
