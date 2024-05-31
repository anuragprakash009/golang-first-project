package mapper

import (
	schema "github.com/ecommerce-api-services/src/model/schema"
	datatypes "github.com/ecommerce-api-services/src/model/type"
)

func DBProductTOResponseMapper(record *schema.Product) *datatypes.GetProduct{
	category := DBCategoryToResponseMapper(&record.Category)
	subCategory := DBSubCategoryToResponseMapper(&record.SubCategory)
	gender := DBGenderToResponseMapper(&record.Gender)
	status := DBStatusToResponseMapper(&record.Status)
	images := DBImagesToResponseMapper(record.Images)
	product := &datatypes.GetProduct{
		Id: record.Id,
		Name: record.Name,
		Description: record.Description,
		Category: category,
		SubCategory: subCategory,
		Stock: record.Stock,
		Gender: gender,
		Status: status,
		Images: images,
		CreatedDate: record.CreatedDate,
		UpdatedDate: record.UpdatedDate,

	}
	return product
}


func DBCategoryToResponseMapper(record *schema.ProductCategory) datatypes.GetCategory{
	category := datatypes.GetCategory{
		Id: record.Id,
		Name: record.Name,
	}
	return category
}

func DBSubCategoryToResponseMapper(record *schema.ProductSubCategory) datatypes.GetSubCategory{
	subCategory := datatypes.GetSubCategory{
		Id: record.Id,
		Name: record.Name,
	}
	return subCategory
}

func DBGenderToResponseMapper(record *schema.Gender) datatypes.GetGender{
	gender := datatypes.GetGender{
		Id: record.Id,
		Name: record.Name,
	}
	return gender
}

func DBStatusToResponseMapper(record *schema.ProductStatus) datatypes.GetProductStatus{
	status := datatypes.GetProductStatus{
		Id: record.Id,
		Name: record.Name,
	}
	return status
}

func DBImagesToResponseMapper(records []schema.ProductImage) []datatypes.GetProductImages{
	productImages:= []datatypes.GetProductImages{}

	for i := 0; i < len(records); i++ {
		image := datatypes.GetProductImages{
			Id: records[i].Id,
			ImgData: records[i].ImgData,
		}
		productImages = append(productImages, image)
	}

	return productImages
}