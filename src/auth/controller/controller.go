package authController

import (
	"net/http"

	authService "github.com/ecommerce-api-services/src/auth/service"
	Datatype "github.com/ecommerce-api-services/src/model/type"
	"github.com/gin-gonic/gin"
)


func Login(context *gin.Context){
	loginPostBody := Datatype.LoginPostBody{}
	err := context.BindJSON(&loginPostBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload",
			"data": nil,
		})
		context.Abort()
		return
	}	
	loginData,err := authService.Login(&loginPostBody)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Login failed",
			"data": nil,
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":loginData,
	})
}


func GenerateAccessToken(context *gin.Context){
	refreshToken := Datatype.RefreshAccessTokenQuery{}
	err := context.ShouldBindQuery(&refreshToken)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload",
			"data": nil,
		})
		context.Abort()
		return
	}	
	accessToken,err := authService.GenerateAccessToken(refreshToken.RefreshToken)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Token Generation Failed",
			"data": nil,
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Token generated successfully",
		"data": accessToken,
	})
}