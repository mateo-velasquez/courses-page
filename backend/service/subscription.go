package service

import (
	"errors"
	"project/client"
	"project/dto"
	"project/model"
)

type subscriptionService struct{}

type subscriptionServiceInterface interface {
	InsertSubscription(subscriptionDTO dto.SubscriptionDTO) (dto.SubscriptionDTO, error)
	GetSubscriptions() (dto.SubscriptionsDTO, error)
	GetSubscriptionById(id int) (dto.SubscriptionDTO, error)
	GetSubscriptionsByUserId(userId int) (dto.SubscriptionsDTO, error)
	GetSubscriptionsByCourseId(courseId int) (dto.SubscriptionsDTO, error)
}

var SubscriptionService subscriptionServiceInterface

func init() {
	SubscriptionService = &subscriptionService{}
}

func (s *subscriptionService) InsertSubscription(subscriptionDTO dto.SubscriptionDTO) (dto.SubscriptionDTO, error) {
	var subscription model.Subscription

	// Registramos los datos del dto en el subscription (model)
	subscription.IDCourse = subscriptionDTO.IDCourse
	subscription.IDUser = subscriptionDTO.IDUser

	if subscription.IDCourse <= 0 {
		return subscriptionDTO, errors.New("error creating subscription. The course doesn't exist")
	}

	if subscription.IDUser <= 0 {
		return subscriptionDTO, errors.New("error creating subscription. The user doesn't exist")
	}

	if subscriptionDTO.CourseRole == "" {
		subscription.CourseRole = "Student"
	} else {
		subscription.CourseRole = "Teacher"
	}

	//Llamamos al cliente para que interactua con la base de datos:
	subscription = client.InsertSubscription(subscription)

	if subscription.IDSubscription == -1 {
		return subscriptionDTO, errors.New("error creating subscription: Subscription allready exists")
	}

	subscriptionDTO.IDSubscription = subscription.IDSubscription
	subscriptionDTO.IndividualRating = subscription.IndividualRating
	subscriptionDTO.Comment = subscription.Comment
	subscriptionDTO.LastUpdateDate = subscription.LastUpdateDate
	subscriptionDTO.CreateDate = subscription.CreateDate

	return subscriptionDTO, nil
}

func (s *subscriptionService) GetSubscriptions() (dto.SubscriptionsDTO, error) {
	var subscriptionsDTO dto.SubscriptionsDTO
	var subscriptions model.Subscriptions = client.GetSubscriptions()
	var subscriptionDTO dto.SubscriptionDTO

	for _, subscription := range subscriptions {
		subscriptionDTO.IDSubscription = subscription.IDSubscription
		subscriptionDTO.IDCourse = subscription.IDCourse
		subscriptionDTO.IDUser = subscription.IDUser
		subscriptionDTO.IndividualRating = subscription.IndividualRating
		subscriptionDTO.Comment = subscription.Comment
		subscriptionDTO.CourseRole = subscription.CourseRole
		subscriptionDTO.LastUpdateDate = subscription.LastUpdateDate
		subscriptionDTO.CreateDate = subscription.CreateDate

		subscriptionsDTO = append(subscriptionsDTO, subscriptionDTO)
	}

	return subscriptionsDTO, nil
}

func (s *subscriptionService) GetSubscriptionById(id int) (dto.SubscriptionDTO, error) {
	var subscription model.Subscription
	var subscriptionDTO dto.SubscriptionDTO

	if id <= 0 {
		return subscriptionDTO, errors.New("ID not found")
	}

	subscription = client.GetSubscriptionById(id)

	if subscription.IDSubscription == 0 {
		return subscriptionDTO, errors.New("user not found")
	}

	subscriptionDTO.IDSubscription = subscription.IDSubscription
	subscriptionDTO.IDCourse = subscription.IDCourse
	subscriptionDTO.IDUser = subscription.IDUser
	subscriptionDTO.IndividualRating = subscription.IndividualRating
	subscriptionDTO.Comment = subscription.Comment
	subscriptionDTO.CourseRole = subscription.CourseRole
	subscriptionDTO.LastUpdateDate = subscription.LastUpdateDate
	subscriptionDTO.CreateDate = subscription.CreateDate

	return subscriptionDTO, nil
}

func (s *subscriptionService) GetSubscriptionsByUserId(userId int) (dto.SubscriptionsDTO, error) {
	var subscriptionsDTO dto.SubscriptionsDTO
	var subscriptions model.Subscriptions = client.GetSubscriptionsByUserId(userId)
	var subscriptionDTO dto.SubscriptionDTO

	for _, subscription := range subscriptions {
		subscriptionDTO.IDSubscription = subscription.IDSubscription
		subscriptionDTO.IDCourse = subscription.IDCourse
		subscriptionDTO.IDUser = subscription.IDUser
		subscriptionDTO.IndividualRating = subscription.IndividualRating
		subscriptionDTO.Comment = subscription.Comment
		subscriptionDTO.CourseRole = subscription.CourseRole
		subscriptionDTO.LastUpdateDate = subscription.LastUpdateDate
		subscriptionDTO.CreateDate = subscription.CreateDate

		subscriptionsDTO = append(subscriptionsDTO, subscriptionDTO)
	}

	return subscriptionsDTO, nil
}

func (s *subscriptionService) GetSubscriptionsByCourseId(courseId int) (dto.SubscriptionsDTO, error) {
	var subscriptionsDTO dto.SubscriptionsDTO
	var subscriptions model.Subscriptions = client.GetSubscriptionsByCourseId(courseId)
	var subscriptionDTO dto.SubscriptionDTO

	for _, subscription := range subscriptions {
		subscriptionDTO.IDSubscription = subscription.IDSubscription
		subscriptionDTO.IDCourse = subscription.IDCourse
		subscriptionDTO.IDUser = subscription.IDUser
		subscriptionDTO.IndividualRating = subscription.IndividualRating
		subscriptionDTO.Comment = subscription.Comment
		subscriptionDTO.CourseRole = subscription.CourseRole
		subscriptionDTO.LastUpdateDate = subscription.LastUpdateDate
		subscriptionDTO.CreateDate = subscription.CreateDate

		subscriptionsDTO = append(subscriptionsDTO, subscriptionDTO)
	}

	return subscriptionsDTO, nil
}
