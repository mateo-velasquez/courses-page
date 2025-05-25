package client

import (
	"project/model"

	log "github.com/sirupsen/logrus"
)

func InsertCourseCategories(relation model.CourseCategories) model.CourseCategories {
	result := Db.Create(&relation)

	if result.Error != nil {
		log.Error("Failed to insert course.")
		relation.IDCourse = -1
		return relation
	}

	log.Debug("Relation created:", relation.IDCourse, relation.IDCategory)
	return relation
}

//func DeleteRelationsByCourseID(id int) model.CourseCategories {
//
//}
