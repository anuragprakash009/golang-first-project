package Datatype

import "time"

type CreateProductBody struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	CategoryId    string   `json:"categoryId"`
	SubCategoryId string   `json:"subCategoryId"`
	Stock         int32    `json:"stock"`
	GenderId      string   `json:"genderId"`
	StatusId      string   `json:"statusId"`
	Images        []string `json:"images"`
}

type GetProduct struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Category    GetCategory    `json:"category"`
	SubCategory GetSubCategory     `json:"subCategory"`
	Stock       int32      `json:"stock"`
	Gender      GetGender     `json:"gender"`
	Status      GetProductStatus     `json:"status"`
	Images      []GetProductImages `json:"images"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
}