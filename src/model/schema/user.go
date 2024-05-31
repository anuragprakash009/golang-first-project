package schema

type User struct {
	Id         string `gorm: "primaryKey;column:beast_id"`
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
	Password   string `gorm:"column:password"`
	UserName   string `gorm:"column:user_name"`
	EmailId    string `gorm:"column:email_id"`
	GenderId   string `gorm:"column:gender_id"`
	Gender     Gender `gorm: "reference:Id`
	IsActive   bool   `gorm:"column:is_active"`
	CreatedAt  string `gorm:"column:created_at"`
	UpdatedAt  string `gorm:"column:updated_at"`
	DeletedAt  string `gorm:"column:deleted_at"`
}

func (this *User) TableName() string {
	return "users"
}
