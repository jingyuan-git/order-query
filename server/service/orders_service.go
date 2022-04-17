package service

import (
	"server/models"
)

type Order struct {
	SearchOrder string `json:"searchOrder,omitempty"`

	StartOrderDate string `json:"startOrderDate,omitempty"`
	EndOrderDate   string `json:"endOrderDate,omitempty"`

	PageNum  int `json:"pageNum,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

func (a *Order) GetAll() ([]*models.OrderDto, error) {

	orders, err := models.GetOrders(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (a *Order) CountAndTotalAmount() (int, int, error) {
	return models.GetOrdersTotal(a.getMaps())
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
