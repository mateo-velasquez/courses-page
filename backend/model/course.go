package model

import (
	"time"
)

type Course struct {
	IDCourse    int       `gorm:"column:course_id;primaryKey;autoIncrement"`
	IDImage     int       `gorm:"column:image_id;not null"`
	CreateDate  time.Time `gorm:"column:create_date;not null;default:CURRENT_TIMESTAMP"`
	LastUpdate  time.Time `gorm:"column:lastupdate_date;not null;default:CURRENT_TIMESTAMP"`
	CourseName  string    `gorm:"column:course_name;type:varchar(300);not null"`
	Price       float64   `gorm:"column:price;type:decimal(10,2);not null"`
	InitDate    time.Time `gorm:"column:init_date;not null"`
	Description string    `gorm:"column:course_description;type:varchar(1000)"`
	Duration    string    `gorm:"column:duration;type:varchar(100);not null"`
	Rating      float64   `gorm:"column:rating;type:decimal(3,2);not null"`
	Categories  Categories
	//Categories  Categories `gorm:"many2many:course_categories;foreignKey:CourseID;joinForeignKey:course_id;References:CategoryID;joinReferences:category_id"`
	//Categories  Categories `gorm:"many2many:course_categories"`
	//Categories  Categories `gorm:"many2many:course_categories;joinForeignKey:IDCourse"`
	//Categories Categories `gorm:"many2many:course_categories;foreignKey:IDCourse;joinForeignKey:IDCourse;joinReferences:IDCategory"`
	// Error: Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'WHERE (`course_categories`.`course_id_course` IN (?))' at line 1
	//Categories []Category `gorm:"many2many:course_categories;joinForeignKey:IDCourse;joinReferences:IDCategory;foreignKey:IDCourse;references:IDCategory"`
}

type Courses []Course
