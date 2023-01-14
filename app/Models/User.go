package Models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Userid       string         `gorm:"unique_index;not_null;type:varchar(100)" json:"user_id"`
	Password     string         `gorm:"type:varchar(100)" json:"password"`
	Email        string         `gorm:"type:varchar(100);unique_index" json:"email"`
	IsTutor      bool           `gorm:"type:bool;default:false" json:"is_tutor"`
	Interests    string         `gorm:"type:varchar(100)" json:"interests"`
	Ownedcourses pq.StringArray `gorm:"type:string[];not_null" json:"owned_courses"`
}

func (user User) UpdateCourses(db *gorm.DB, course string) (User, error) {
	user.Ownedcourses = append(user.Ownedcourses, course)
	err := db.Save(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) Create(db *gorm.DB) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashedPassword)
	user.Ownedcourses = []string{}
	err = db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) FindOne(db *gorm.DB, userid string) (User, error) {
	err := db.Where("Userid=?", userid).First(&user).Error
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

func (user User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

type UserLogin struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
