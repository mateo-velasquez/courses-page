package client

import (
	"project/model"

	log "github.com/sirupsen/logrus"
)

func GetCourseById(id int) model.Course {
	var course model.Course

	Db.Where("id_course = ?", id).Preload("Categories").First(&course)
	log.Debug("Course: ", course)

	return course
}

func GetCourses() model.Courses {
	var courses model.Courses
	Db.Preload("Categories").Find(&courses)

	log.Debug("Courses: ", courses)

	return courses
}

func GetCoursesByName(name string) model.Courses {
	var courses model.Courses

	Db.Where("name LIKE ?", "%"+name+"%").Find(&courses)
	log.Debug("Users: ", courses)

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
	err := Db.Delete(&course).Error

	if err != nil {
		log.Debug("Failed to delete course")
	} else {
		log.Debug("Course deleted: ", course.IDCourse)
	}
	return err
}
