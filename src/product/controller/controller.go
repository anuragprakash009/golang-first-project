package productController

import (
	"net/http"

	Datatype "github.com/ecommerce-api-services/src/model/type"
	productService "github.com/ecommerce-api-services/src/product/service"
	"github.com/gin-gonic/gin"
)




func CreateProduct(context *gin.Context) {
	var createProductBody Datatype.CreateProductBody
	err := context.BindJSON(&createProductBody)
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Invalid payload",
		})
		context.Abort()
		return
	}
	newProduct,err := productService.CreateProduct(&createProductBody);
	if err != nil {
		context.JSON(500, gin.H{
			"message": "Unable to create product",
		})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{
		"message": "Product Added successfully",
		"data":newProduct,
	})
}


func GetProduct(context *gin.Context) {
	productId := context.Param("productId")
	if productId == " " {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		context.Abort()
		return
	}
	product, err:=productService.GetProduct(productId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid product id",
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Product fetched successfully",
		"data": product,
	})
}
