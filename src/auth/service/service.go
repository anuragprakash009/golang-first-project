package authService

import (
	"errors"
	"time"

	authRepository "github.com/ecommerce-api-services/src/auth/repository"
	Datatype "github.com/ecommerce-api-services/src/model/type"
	jwt "github.com/golang-jwt/jwt"
)


var secretKey = []byte("asdkjhsdjkhsjkdhsjkhjshdjkhjksahjkdhsjkdh")


func Login(login *Datatype.LoginPostBody) (*Datatype.LoginResponse, error) {
	user, err:=authRepository.FindUserByEmail(login.EmailId)
	if user == nil{
		return nil, errors.New("Couldn't find user")
	}

	if err != nil{
		return nil, err
	}

	accessTokenExpiration := time.Now().Add(time.Hour * 3).Unix()

	refreshTokenExpiration := time.Now().AddDate(0,3,0).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ 
        "userId": user.UserName,
		"firstName": user.FirstName,
		"middleName": user.MiddleName,
		"lastName": user.LastName,
		"emailId": user.EmailId,
		"gender":user.Gender.Name, 
        "exp": accessTokenExpiration, 
    })

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ 
        "userId": user.UserName,
		"firstName": user.FirstName,
		"middleName": user.MiddleName,
		"lastName": user.LastName,
		"emailId": user.EmailId,
		"gender": user.Gender.Name,
        "exp": refreshTokenExpiration,
    })

	accessTokenString, err := accessToken.SignedString(secretKey)
    if err != nil {
    	return nil, errors.New("Access token generation failed: " + err.Error())
    }

	refreshTokenString, err := refreshToken.SignedString(secretKey)
    if err != nil {
    	return nil, errors.New("Refresh token generation failed: " + err.Error())
    }
	loginBody := &Datatype.LoginResponse{
		AccessToken: accessTokenString,
		RefreshToken: refreshTokenString,
		AccessTokenValidity: accessTokenExpiration,
		RefreshTokenValidity: refreshTokenExpiration,
	}
	return loginBody, nil
}


func GenerateAccessToken(refreshToken string) (*Datatype.LoginResponse, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	 })
	 if err != nil {
		return nil, err
	 }
	
	 if !token.Valid {
		return nil, errors.New("Invalid refresh token")
	 }
	 claims, _ := token.Claims.(jwt.MapClaims)

	 emailId := claims["emailId"]

	 user, err:=authRepository.FindUserByEmail(emailId)
	 if user == nil{
		 return nil, errors.New("Couldn't find user")
	 }
 
	 if err != nil{
		 return nil, err
	 }
	 return nil, nil
}