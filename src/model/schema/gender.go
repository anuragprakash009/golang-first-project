package schema

import "time"

type Gender struct {
	Id          string `json:"id" gorm: "primaryKey"`
	Name        string `json:"name" gorm: "index"`
	IsActive    bool   `json:"isActive"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json: "updatedDate"`
	DeletedDate *time.Time `json: "deletedDate"`
}
