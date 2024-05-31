package productService

import (
	"time"

	schemas "github.com/ecommerce-api-services/src/model/schema"
	Datatype "github.com/ecommerce-api-services/src/model/type"
	mapper "github.com/ecommerce-api-services/src/product/mapper"
	productRepository "github.com/ecommerce-api-services/src/product/repository"
	"github.com/google/uuid"
)


func CreateProduct(product *Datatype.CreateProductBody) (*Datatype.GetProduct,error){
	productId := uuid.New().String()
	newProduct := schemas.Product{
		Id: productId,
		Name: product.Name,
		Description: product.Description,
		CategoryId: product.CategoryId,
		SubCategoryId: product.SubCategoryId,
		GenderId: product.GenderId,
		Stock: product.Stock,
		IsActive: true,
		StatusId: product.StatusId,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		DeletedDate: nil,
	}
	var newImages []schemas.ProductImage
	for i := 0; i < len(product.Images); i++ {
		newImage := schemas.ProductImage{
			Id: uuid.New().String(),
			ImgData: product.Images[i],
			ProductId: &productId,
			IsActive: true,
		}
		newImages = append(newImages, newImage)
	}
	err:=productRepository.CreateProductAndImages(&newProduct, &newImages)
	if err != nil {
		return nil,err
	}
	responseProduct, err:= GetProduct(productId)
	if err != nil {
		return nil, err
	}
	return responseProduct,nil
}


func GetProduct(productId string) (*Datatype.GetProduct,error) {
	data, err:= productRepository.GetProductById(productId)
	if err != nil {
		return nil,err
	}
	response:=mapper.DBProductTOResponseMapper(data)
	return response,nil
}