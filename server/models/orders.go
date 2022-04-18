package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         int         `gorm:"primarykey"`
	CreatedAt  time.Time   `csv:"created_at" json:"createdAt,omitempty"`
	OrderName  string      `csv:"order_name" json:"orderName,omitempty"`
	CustomerId string      `csv:"customer_id" json:"customerId,omitempty"`
	Customer   Customer    `gorm:"foreignKey:CustomerId" json:"customer,omitempty"`
	OrderItem  []OrderItem `gorm:"foreignKey:OrderId;references:ID"`
}

func dbSearchOrder(db *gorm.DB, keyWord string, pagingCconstraints bool, pageNum int, pageSize int) *gorm.DB {
	db.Preload("Customer").Preload("Customer.CustomerCompany").Preload("OrderItem").Preload("OrderItem.Delivery")
	db.Joins("JOIN customers on customers.id = orders.customer_id")
	db.Joins("JOIN order_items ON orders.id = order_items.order_id")
	db.Distinct("orders.id", "orders.created_at", "orders.order_name", "orders.customer_id")
	db.Where("orders.order_name ilike ? or order_items.product ilike ?", "%"+keyWord+"%", "%"+keyWord+"%")
	db.Distinct("orders.id", "orders.created_at", "orders.order_name", "orders.customer_id")
	if pagingCconstraints {
		db.Offset(pageNum).Limit(pageSize)
	}
	return db
}

func dbFilterOrderDate(db *gorm.DB, startDate string, endDate string, pagingCconstraints bool, pageNum int, pageSize int) *gorm.DB {
	db.Preload("Customer").Preload("Customer.CustomerCompany").Preload("OrderItem").Preload("OrderItem.Delivery")
	db.Where("orders.created_at > to_date(?,'yyyy-MM-dd hh24:mi:ss') and orders.created_at < to_date(?,'yyyy-MM-dd hh24:mi:ss')", startDate, endDate)
	if pagingCconstraints {
		db.Offset(pageNum).Limit(pageSize)
	}
	return db
}

func dbListAllOrders(db *gorm.DB, pagingCconstraints bool, pageNum int, pageSize int) *gorm.DB {
	db.Preload("Customer").Preload("Customer.CustomerCompany").Preload("OrderItem").Preload("OrderItem.Delivery")
	if pagingCconstraints {
		db.Offset(pageNum).Limit(pageSize)
	}
	return db
}

// GetOrdersTotal gets the total number of orders based on the constraints
func GetOrdersTotal(maps map[string]interface{}) ([]*Order, error) {
	var (
		orders []*Order
	)

	if v, ok := maps["searchOrder"]; ok {
		db := db.Model(orders)
		db = dbSearchOrder(db, v.(string), false, 0, 0)
		err := db.Find(&orders).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		startDate, ok1 := maps["startOrderDate"]
		endDate, ok2 := maps["endOrderDate"]
		if ok1 && ok2 {
			db := db.Model(orders)
			db = dbFilterOrderDate(db, startDate.(string), endDate.(string), false, 0, 0)
			err := db.Find(&orders).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return nil, err
			}
		} else {
			db := db.Model(orders)
			db = dbListAllOrders(db, false, 0, 0)
			err := db.Find(&orders).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
	}
	return orders, nil
}

// GetOrders gets a list of orders based on paging constraints
func GetOrders(pageNum int, pageSize int, maps map[string]interface{}) ([]*Order, error) {
	var (
		orders []*Order
	)

	// search orders by part of the order or product name
	if v, ok := maps["searchOrder"]; ok {
		db := db.Model(orders)
		db = dbSearchOrder(db, v.(string), true, pageNum, pageSize)
		err := db.Find(&orders).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
	} else {
		// filter orders by date range
		startDate, ok1 := maps["startOrderDate"]
		endDate, ok2 := maps["endOrderDate"]
		if ok1 && ok2 {
			db := db.Model(orders)
			db = dbFilterOrderDate(db, startDate.(string), endDate.(string), true, pageNum, pageSize)
			err := db.Find(&orders).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return nil, err
			}
		} else {
			db := db.Model(orders)
			db = dbListAllOrders(db, true, pageNum, pageSize)
			err := db.Find(&orders).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
	}
	return orders, nil
}
