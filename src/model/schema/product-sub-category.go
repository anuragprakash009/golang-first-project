package schema

import "time"

type ProductSubCategory struct {
	Id          string `json:"id" gorm: "primaryKey"`
	Name        string `json:"name"`
	ParentCategoryId string `json:"parentCategoryId"`
	ParentCategory ProductCategory `grom "references:Id"`
	IsActive bool `json:"isActive" gorm:"default:false"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate"`
}