package main

import (
	"fmt"

	authController "github.com/ecommerce-api-services/src/auth/controller"
	database "github.com/ecommerce-api-services/src/database/sql"
	productController "github.com/ecommerce-api-services/src/product/controller"
	"github.com/gin-gonic/gin"
)

 func main() {
	dbConnectionerr := database.Connect()
	if dbConnectionerr != nil {	
		fmt.Println(dbConnectionerr)
		panic(dbConnectionerr)

	}
	fmt.Println("Database connected successfully")
	// migrateError := database.Migrate()
	// if migrateError != nil {
	// 	panic(migrateError)
	// }
	// fmt.Println("Database migrated successfully")
	app := gin.Default()

	productRoutes := app.Group("/v1/product")
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/:productId", productController.GetProduct)
	}

	authRoutes := app.Group("/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.GET("/access-token", authController.GenerateAccessToken)
	}
	app.Run() // listen and serve on 0.0.0.0:8080
}