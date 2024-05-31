package schema

import (
	"time"
)

type Product struct {
	Id          string `json: "id" gorm: "primaryKey"`
	Name        string `json: "name" gorm:"index"`
	Description string `json: "description"`
	CategoryId  string `json: "categoryId"`
 	Category  ProductCategory `json:"productCategory" gorm:"references:Id"`
	SubCategoryId string `json: "subCategoryId"`
	SubCategory ProductSubCategory `json:"subCategory" gorm:"references:Id"`
	GenderId string `json: "genderId"`
	Gender Gender `json:"gender" gorm:"references:Id"`
	Stock       int32 `json: "stock"`
	IsActive    bool `json: "isActive"`
	StatusId	string `json:"statusId"`
	Status ProductStatus `json:"status" gorm:"references:Id"`
	Images [] ProductImage `gorm:"references:Id`
	CreatedDate time.Time `json:"createdDate" gorm:"index"`
	UpdatedDate time.Time`json:"updatedDate" gorm:"index"`
	DeletedDate *time.Time `json:"deletedDate"`
}
 