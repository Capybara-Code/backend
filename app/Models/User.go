package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"primary_key;not_null;type:varchar(100)" json:"user_id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	IsTutor  bool   `gorm:"type:bool;default:false" json:"is_tutor"`
}

func (user User) Create(db *gorm.DB) (User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) FindAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (user User) FindOne(db *gorm.DB) (User, error) {
	err := db.First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) Update(db *gorm.DB) (User, error) {
	err := db.Save(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) Delete(db *gorm.DB) (User, error) {
	err := db.Delete(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
