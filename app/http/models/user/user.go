package user

import "goblog/app/http/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"colum:name;type:varchar(255);not null;unique"`
	Email    string `gorm:"colum:email;type:varchar(255);default:NULL;unique"`
	Password string `gorm:"colum:password;type:varchar(255)"`
}
