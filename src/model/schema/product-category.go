package schema

import "time"

type ProductCategory struct {
	Id          string `json:"id" gorm: "primaryKey"`
	Name        string `json:"name"`
	IsActive bool `json:"isActive" gorm:"default:false"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json: "updatedDate"`
	DeletedDate *time.Time	`json: "deletedDate"`
}
