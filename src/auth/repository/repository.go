package authRepository

import (
	database "github.com/ecommerce-api-services/src/database/sql"
	"github.com/ecommerce-api-services/src/model/schema"
)

func FindUserByEmail(email string) (*schema.User,error) {
	conn, err:= database.GetConnection()
	if err != nil {
		return nil,err
	}
	user := &schema.User{
		EmailId:email,
		IsActive: true,
	}
	conn.Preload("Gender").Find(&user)
	if user.Id == ""{
		return nil,nil
	}
	return user, nil
}
