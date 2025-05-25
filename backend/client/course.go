package client

import (
	"project/model"

	log "github.com/sirupsen/logrus"
)

func GetCourseById(id int) model.Course {
	var course model.Course

	Db.Where("course_id = ?", id).First(&course)
	// Aparentemente GORM está generando un JOIN sin la condición ON. Eso rompe completamente la consulta.
	//Db.Where("course_id = ?", id).Preload("Categories").First(&course)
	var categories model.Categories
	Db.Table("categories").
		Joins("JOIN course_categories ON course_categories.category_id = categories.category_id").
		Where("course_categories.course_id = ?", id).
		Scan(&categories)

	course.Categories = categories

	log.Debug("Course: ", course)

	return course
}

func GetCourses() model.Courses {
	var courses model.Courses
	Db.Find(&courses)

	for i := range courses {
		var categories model.Categories
		Db.Table("categories").
			Joins("JOIN course_categories ON course_categories.category_id = categories.category_id").
			Where("course_categories.course_id = ?", courses[i].IDCourse).
			Scan(&categories)
		courses[i].Categories = categories
	}

	log.Debug("Courses: ", courses)

	return courses
}

func GetCoursesByName(name string) model.Courses {
	var courses model.Courses

	Db.Where("course_name LIKE ?", "%"+name+"%").Find(&courses)

	for i := range courses {
		var categories model.Categories
		Db.Table("categories").
			Joins("JOIN course_categories ON course_categories.category_id = categories.category_id").
			Where("course_categories.course_id = ?", courses[i].IDCourse).
			Scan(&categories)
		courses[i].Categories = categories
	}

	log.Debug("Courses: ", courses)

	return courses
}

func InsertCourse(course model.Course) model.Course {
	result := Db.Create(&course)

	if result.Error != nil {
		log.Error("Failed to insert course.")
		return course
	}

	log.Debug("Course created:", course.IDCourse)
	return course
}

func PutCourseById(course model.Course) model.Course {
	var newCategories model.Categories

	for _, category := range course.Categories {
		newCategories = append(newCategories, category)
	}

	result := Db.Save(&course)

	Db.Model(&course).Association("Categories").Replace(newCategories)

	if result.Error != nil {
		log.Debug("Failed to update course")
		return model.Course{}
	}

	log.Debug("Updated course: ", course.IDCourse)
	return course
}

func DeleteCourseById(course model.Course) error {
	err := Db.Debug().Delete(&course).Error

	if err != nil {
		log.Debug("Failed to delete course")
	} else {
		log.Debug("Course deleted: ", course.IDCourse)
	}
	return err
}
