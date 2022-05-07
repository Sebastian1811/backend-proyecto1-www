package entity

type User struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Email    string `json:"email" binding:"required,gte=0,lte=50,email" gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(256)"`
}
