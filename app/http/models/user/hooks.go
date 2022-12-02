package user

import (
	"goblog/pkg/password"
	"gorm.io/gorm"
)

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}
