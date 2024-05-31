package productRepository

import (
	"fmt"
	"log"

	database "github.com/ecommerce-api-services/src/database/sql"
	"github.com/ecommerce-api-services/src/model/schema"
	"gorm.io/gorm"
)


func CreateProductAndImages(product *schema.Product, images *[]schema.ProductImage) error{
	conn,err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	err=conn.Transaction(func(tx *gorm.DB) error{
		err := tx.Create(&product).Error
		if err != nil {
			return err
		}
		if len(*images) > 0 {
			err = tx.Create(&images).Error
			if(err != nil){
				return err
			}
		}

		return nil
	})
	if err != nil{
		return err
	}
	return nil
}


func GetProductById(productId string) (*schema.Product,error){
	conn,err := database.GetConnection()
	if err != nil{
		return nil,err
	}
	product := schema.Product{
		Id: productId,
		IsActive: true,
	}

	conn.Preload("Category").Preload("SubCategory").Preload("Gender").Preload("Status").Preload("Images").First(&product)

	fmt.Println(product)
	return &product, nil
}