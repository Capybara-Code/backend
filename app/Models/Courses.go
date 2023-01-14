package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CourseName  string    `gorm:"type:varchar(100)" json:"course_name"`
	Author      string    `gorm:"type:varchar(100)" json:"author"`
	Tags        string    `gorm:"type:varchar(100)" json:"tags"`
	Price       float64   `gorm:"type:float" json:"price"`
	Description string    `gorm:"type:varchar(200)" json:"description"`
	Authorpk    string    `gorm:"type:varchar(100)" json:"author_pk"`
}

func (course Course) Create(db *gorm.DB) (Course, error) {
	err := db.Create(&course).Error
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func (course Course) FindFuzzy(db *gorm.DB, search string) ([]Course, error) {
	courses := []Course{}
	err := db.Where("Author LIKE ?", "%"+search+"%").Find(&courses).Error
	if err != nil {
		return []Course{}, err
	}
	return courses, nil
}

func (course Course) FindOne(db *gorm.DB, id string) (Course, error) {

	err := db.Where("ID=?", id).First(&course).Error
	if err != nil {
		return Course{}, err
	}
	return course, nil
}

func (course Course) FindOneByAuthor(db *gorm.DB, author string) ([]Course, error) {
	courses := []Course{}
	err := db.Where("Author=?", author).Find(&courses).Error
	if err != nil {
		return []Course{}, err
	}
	return courses, nil
}

func (course Course) FindAll(db *gorm.DB) ([]Course, error) {
	var courses []Course
	err := db.Find(&courses).Error
	if err != nil {
		return []Course{}, err
	}
	return courses, nil
}

func (course Course) Delete(db *gorm.DB) (Course, error) {
	err := db.Delete(&course).Error
	if err != nil {
		return Course{}, err
	}
	return course, nil
}
