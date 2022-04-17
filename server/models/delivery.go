package models

type Delivery struct {
	Id                string `csv:"id" json:"id,omitempty" gorm:"primary_key"`
	OrderItemId       string `csv:"order_item_id" json:"orderItemId,omitempty"`
	DeliveredQuantity int `csv:"delivered_quantity" json:"deliveredQuantity,omitempty"`
}
