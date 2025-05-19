package model

import (
	"time"
)

type Subscription struct {
	IDSubscription   int       `gorm:"column:id_subscription;primaryKey;autoIncrement"`
	IDCourse         int       `gorm:"column:id_course;not null"`
	IDUser           int       `gorm:"column:id_user;not null"`
	CreateDate       time.Time `gorm:"column:create_date;not null;default:CURRENT_TIMESTAMP"`
	IndividualRating float64   `gorm:"column:individual_rating;type:decimal(3,2);check:individual_rating>=0 AND individual_rating<=5"`
	Comment          string    `gorm:"column:comment;type:varchar(300)"`
	LastUpdateDate   time.Time `gorm:"column:LastUpdate_date;not null;default:CURRENT_TIMESTAMP"`
}

// TableName especifica el nombre exacto de la tabla en la base de datos
func (Subscription) TableName() string {
	return "subscription"
}

type Subscriptions []Subscription
