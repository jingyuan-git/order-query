package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/pkg/setting"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var data_source string
	var err error

	if setting.DatabaseSetting.Type == "postgres" {
		data_source = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Name,
			setting.DatabaseSetting.Port,
			setting.DatabaseSetting.TimeZone,
		)
	} else {
		log.Fatalf("[error] sorry, the current database type is not yet supported")
	}

	db, err = gorm.Open(postgres.Open(data_source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	// Clean old data in database
	db.Migrator().DropTable(
		[]Order{},
		[]OrderItem{},
		[]Delivery{},
		[]Customer{},
		[]CustomerCompany{},
	)

	// Migrate order associated table
	db.AutoMigrate(
		[]Order{},
		[]OrderItem{},
		[]Delivery{},
		[]Customer{},
		[]CustomerCompany{},
	)

	orders := csvOutput(setting.DatabaseSetting.DataPath+"orders.csv", "Order").([]Order)
	orderItems := csvOutput(setting.DatabaseSetting.DataPath+"order_items.csv", "OrderItem").([]OrderItem)
	deliveries := csvOutput(setting.DatabaseSetting.DataPath+"deliveries.csv", "Delivery").([]Delivery)
	customers := csvOutput(setting.DatabaseSetting.DataPath+"customers.csv", "Customer").([]Customer)
	customerCompanies := csvOutput(setting.DatabaseSetting.DataPath+"customer_companies.csv", "CustomerCompany").([]CustomerCompany)

	result := db.Create(&customerCompanies)
	if result.Error != nil {
		panic(result.Error)
	}
	result = db.Create(&customers)
	if result.Error != nil {
		panic(result.Error)
	}
	result = db.Create(&orders)
	if result.Error != nil {
		panic(result.Error)
	}
	result = db.Create(&orderItems)
	if result.Error != nil {
		panic(result.Error)
	}
	result = db.Create(&deliveries)
	if result.Error != nil {
		panic(result.Error)
	}

	log.Default().Printf("the databases are successfully loaded")
}

func csvOutput(path string, structType string) interface{} {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	switch structType {
	case "Order":
		{
			var orders []Order
			err = gocsv.Unmarshal(file, &orders)
			if err != nil {
				panic(err)
			}
			return orders
		}
	case "OrderItem":
		{
			var orderItems []OrderItem
			err = gocsv.Unmarshal(file, &orderItems)
			if err != nil {
				panic(err)
			}
			return orderItems
		}
	case "Customer":
		{
			var customers []Customer
			err = gocsv.Unmarshal(file, &customers)
			if err != nil {
				panic(err)
			}
			return customers
		}
	case "Delivery":
		{
			var deliveries []Delivery
			err = gocsv.Unmarshal(file, &deliveries)
			if err != nil {
				panic(err)
			}
			return deliveries
		}
	case "CustomerCompany":
		{
			var customerCompanies []CustomerCompany
			err = gocsv.Unmarshal(file, &customerCompanies)
			if err != nil {
				panic(err)
			}
			return customerCompanies
		}
	}
	return nil
}
