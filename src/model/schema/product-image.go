package schema

type ProductImage struct {
	Id        string  `json: "id" gorm: "primaryKey"`
	ImgData   string  `json: "image" gorm: "index"`
	ProductId *string `json:"productId"`
	IsActive  bool    `json:"isActive"`
}
