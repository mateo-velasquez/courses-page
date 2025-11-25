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

func GetSubscriptionById(id int) model.Subscription {
	var subscription model.Subscription

	result := Db.Where("subscription_id = ?", id).First(&subscription)

	if result.Error != nil {
		log.Error("Failed to find Subscription.")
		subscription.IDSubscription = -1
		return subscription
	}

	log.Debug("Subscription: ", subscription)

	return subscription
}

func GetSubscriptionsByUserId(userId int) model.Subscriptions {
	var subscriptions model.Subscriptions

	result := Db.Where("user_id = ?", userId).Find(&subscriptions)

	if result.Error != nil {
		log.Error("Failed to find Subscription.")
		return subscriptions
	}

	log.Debug("Subscriptions: ", subscriptions)

	return subscriptions
}

func GetSubscriptionsByCourseId(courseId int) model.Subscriptions {
	var subscriptions model.Subscriptions

	result := Db.Where("course_id = ?", courseId).Find(&subscriptions)

	if result.Error != nil {
		log.Error("Failed to find Subscription.")
		return subscriptions
	}

	log.Debug("Subscriptions: ", subscriptions)

	return subscriptions
}

func PutRating(subscription model.Subscription) model.Subscription {

	// Primero buscamos el subscription original
	result := Db.Table("subscriptions").
		Where("subscription_id = ?", subscription.IDSubscription).
		Update("individual_rating", subscription.IndividualRating)

	if result.Error != nil {
		log.Error("Failed to Save Subscription rating.")
		subscription.IDSubscription = -2 // Código de error interno
		return subscription
	}

	// Verificamos si realmente se actualizó algo (si el ID existía)
	if result.RowsAffected == 0 {
		log.Error("Subscription not found to update.")
		subscription.IDSubscription = -1 // Código para "No encontrado"
		return subscription
	}

	return subscription
}

func PutComment(subscription model.Subscription) model.Subscription {

	// Primero buscamos el subscription original
	result := Db.Table("subscriptions").
		Where("subscription_id = ?", subscription.IDSubscription).
		Update("comment", subscription.Comment)

	if result.Error != nil {
		log.Error("Failed to Save Subscription comment.")
		subscription.IDSubscription = -2 // Código de error interno
		return subscription
	}

	// Verificamos si realmente se actualizó algo (si el ID existía)
	if result.RowsAffected == 0 {
		log.Error("Subscription not found to update.")
		subscription.IDSubscription = -1 // Código para "No encontrado"
		return subscription
	}

	return subscription
}
