package service

import (
	"log"
	"time"

	"server/models"
)

type Order struct {
	SearchOrder string `json:"searchOrder,omitempty"`

	StartOrderDate string `json:"startOrderDate,omitempty"`
	EndOrderDate   string `json:"endOrderDate,omitempty"`

	PageNum  int `json:"pageNum,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type OrderDto struct {
	OrderName       string    `json:"orderName,omitempty"`
	CompanyName     string    `json:"companyName,omitempty"`
	CustomerName    string    `json:"customerName,omitempty"`
	OrderData       time.Time `json:"orderData,omitempty"`
	DeliveredAmount int       `json:"deliveredAmount,omitempty"`
	TotalAmount     int       `json:"totalAmount,omitempty"`
}

func (a *Order) GetAll() ([]*OrderDto, error) {
	orders, err := models.GetOrders(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		log.Default().Printf("fail to list all orders, error: %+v \n", err)
		return nil, err
	}

	return orderListResp(orders), nil
}

func orderListResp(orders []*models.Order) []*OrderDto {
	var orderDtos []*OrderDto
	for _, v := range orders {
		var totalAmount, deliveredAmount int
		for _, orderItem := range v.OrderItem {
			var deliveredQuantity int
			totalAmount += orderItem.PricePerUnit * orderItem.Quantity
			for _, orderDeliveredQuantity := range orderItem.Delivery {
				deliveredQuantity += orderDeliveredQuantity.DeliveredQuantity
			}
			deliveredAmount += orderItem.PricePerUnit * deliveredQuantity
		}
		orderDto := OrderDto{
			OrderName:       v.OrderName,
			CompanyName:     v.Customer.CustomerCompany.CompanyName,
			CustomerName:    v.Customer.Name,
			OrderData:       v.CreatedAt,
			DeliveredAmount: deliveredAmount,
			TotalAmount:     totalAmount,
		}
		orderDtos = append(orderDtos, &orderDto)
	}
	return orderDtos
}

func (a *Order) CountAndTotalAmount() (int, int, error) {
	orders, err := models.GetOrdersTotal(a.getMaps())
	if err != nil {
		log.Default().Printf("fail to list all orders, error: %+v \n", err)
		return 0, 0, err
	}

	return orderTotalResp(orders)
}

func orderTotalResp(orders []*models.Order) (int, int, error) {
	var totalAmount int
	for _, v := range orders {
		for _, orderItem := range v.OrderItem {
			totalAmount += orderItem.PricePerUnit * orderItem.Quantity
		}
	}
	return len(orders), totalAmount, nil
}

func (a *Order) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if a.SearchOrder != "" {
		maps["searchOrder"] = a.SearchOrder
	}
	if a.StartOrderDate != "" {
		maps["startOrderDate"] = a.StartOrderDate
	}
	if a.EndOrderDate != "" {
		maps["endOrderDate"] = a.EndOrderDate
	}
	return maps
}
