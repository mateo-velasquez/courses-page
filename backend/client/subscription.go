package client

import (
	"project/model"

	log "github.com/sirupsen/logrus"
)

func InsertSubscription(subscription model.Subscription) model.Subscription {
	var subscriptionaux model.Subscription

	Db.Where("course_id = ? AND user_id = ?", subscription.IDCourse, subscription.IDUser).First(&subscriptionaux)

	if subscriptionaux.IDSubscription != 0 {
		log.Error("Subscription allready exists!")
		subscriptionaux.IDSubscription = -1
		return subscriptionaux
	}

	result := Db.Create(&subscription)

	if result.Error != nil {
		log.Error("Failed to insert Subscription.")
		return subscription
	}

	log.Debug("Subscription created:", subscription.IDSubscription)
	return subscription
}

func GetSubscriptions() model.Subscriptions {
	var subscriptions model.Subscriptions
	Db.Find(&subscriptions)

	log.Debug("Subscriptions: ", subscriptions)

	return subscriptions
}
