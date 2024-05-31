package database

import (
	"errors"
	"fmt"

	schema "github.com/ecommerce-api-services/src/model/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var postgresConnection *gorm.DB 


func Connect() error {
	fmt.Print(postgresConnection)
	dsn := "host=localhost user=postgres password=postgres dbname=e_commerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return errors.New("unable to connect to database")
	}
	postgresConnection = db
	return nil
}


func GetConnection() ( *gorm.DB , error){
	fmt.Print(postgresConnection)
	if postgresConnection == nil {
		return nil, errors.New("postgress connection not established")
	}
	return postgresConnection, nil
}


func Migrate() error {
	postgresConnection.AutoMigrate(&schema.Gender{},&schema.ProductSubCategory{},&schema.ProductCategory{},&schema.ProductStatus{},&schema.Product{},&schema.ProductImage{})
	return nil
}


