package schema

type ProductStatus struct {
	Id       string `json: "id" gorm: "primaryKey"`
	Name     string `json: "name"`
	IsActive bool   `json: "isActive" gorm:"default:false"`
}
