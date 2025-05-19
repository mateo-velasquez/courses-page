package model

import (
	"time"
)

type Course struct {
	IDCourse    int        `gorm:"column:id_course;primaryKey;autoIncrement"`
	IDImage     int        `gorm:"column:id_image;not null"`
	CreateDate  time.Time  `gorm:"column:create_date;not null;default:CURRENT_TIMESTAMP"`
	LastUpdate  time.Time  `gorm:"column:LastUpdate_date;not null;default:CURRENT_TIMESTAMP"`
	CourseName  string     `gorm:"column:course_name;type:varchar(300);not null"`
	Price       float64    `gorm:"column:price;type:decimal(10,2);not null"`
	InitDate    string     `gorm:"column:init_date;type:varchar(16);not null"`
	Description string     `gorm:"column:description;type:varchar(1000)"`
	Duration    string     `gorm:"column:duration;type:varchar(100);not null"`
	Rating      float64    `gorm:"column:rating;type:decimal(3,2);not null"`
	Categories  Categories `gorm:"many2many:course_category;"`
}

// TableName especifica el nombre exacto de la tabla en la base de datos
func (Course) TableName() string {
	return "course"
}

type Courses []Course
