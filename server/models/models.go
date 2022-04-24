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

	errs := make(chan error, 1)

	go func() {
		errs <- importDB(setting.DatabaseSetting.DataPath+"orders.csv", []Order{})
	}()

	go func() {
		errs <- importDB(setting.DatabaseSetting.DataPath+"order_items.csv", []OrderItem{})
	}()

	go func() {
		errs <- importDB(setting.DatabaseSetting.DataPath+"deliveries.csv", []Delivery{})
	}()

	go func() {
		errs <- importDB(setting.DatabaseSetting.DataPath+"customers.csv", []Customer{})
	}()

	go func() {
		errs <- importDB(setting.DatabaseSetting.DataPath+"customer_companies.csv", []CustomerCompany{})
	}()

	if err := <-errs; err != nil {
		log.Default().Panicf("the database is failed to import. Error: %+v", err)
	}

	log.Default().Printf("the databases are successfully loaded")
}

// importDB parses the CSV from the file in the interface
// and insert the value into database.
func importDB[V any](path string, s V) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gocsv.Unmarshal(file, &s)
	if err != nil {
		return err
	}

	result := db.Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
