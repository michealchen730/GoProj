package models

import (
	"GoProj1.0/databases"
	"log"
)

type User struct {
	Id uint64 `gorm:"primary_key" json:"id"`

	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) UserSave() error {
	if err := databases.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) UserDelete() error {
	if err := databases.Db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func UserGetByUsername(username string) (User, error) {
	var user User
	if err := databases.Db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func UserTest() error {
	var user User
	if err := databases.Db.First(user).Error; err != nil {
		return err
	}
	log.Println(user.Username)
	log.Println(user.Password)
	return nil
}
