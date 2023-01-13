package Models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   string `gorm:"primary_key;not_null;type:varchar(100)" json:"user_id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	IsTutor  bool   `gorm:"type:bool;default:false" json:"is_tutor"`
}

func (user User) Create(db *gorm.DB) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashedPassword)
	err = db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) FindOne(db *gorm.DB) (User, error) {
	err := db.First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
