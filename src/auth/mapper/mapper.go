package authMapper

import (
	"github.com/ecommerce-api-services/src/model/schema"
	Datatype "github.com/ecommerce-api-services/src/model/type"
)
func MapDBRecordToJWTPayload(record *schema.User) (* Datatype.JwtPayload){
	payload := Datatype.JwtPayload{
		FirstName: record.FirstName,
		MiddleName: record.MiddleName,
		LastName: record.LastName,
		UserId: record.UserName,
		EmailId: record.EmailId,
		Gender: record.Gender.Name,
	} 
	return &payload
}
