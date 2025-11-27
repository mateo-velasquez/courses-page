package controller

import (
	"net/http"
	"project/dto"
	"project/service"

	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertSuscription(c *gin.Context) {
	var subscriptionDto dto.SubscriptionDTO
	err := c.BindJSON(&subscriptionDto)

	// If deserialization fails (fields or something is missing in the body) then throw a badrequest
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} //devolvemos un objeto JSON con un campo error que contiene el mensaje de error.

	subscriptionDto, er := service.SubscriptionService.InsertSubscription(subscriptionDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, subscriptionDto) // Si está todo bien devuelve un 201 y el SubscriptionDTO
}

func GetSubscriptions(c *gin.Context) {
	var subscriptionsDto dto.SubscriptionsDTO

	subscriptionsDto, err := service.SubscriptionService.GetSubscriptions()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptionsDto)
}

func GetSubscriptionById(c *gin.Context) {
	var subscriptionDto dto.SubscriptionDTO

	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido en la URL"})
		return
	}

	subscriptionDto, err := service.SubscriptionService.GetSubscriptionById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptionDto)
}

func GetSubscriptionsByUserId(c *gin.Context) {
	var subscriptionsDto dto.SubscriptionsDTO

	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido en la URL"})
		return
	}

	subscriptionsDto, err := service.SubscriptionService.GetSubscriptionsByUserId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptionsDto)
}

func GetSubscriptionsByCourseId(c *gin.Context) {
	var subscriptionsDto dto.SubscriptionsDTO

	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido en la URL"})
		return
	}

	subscriptionsDto, err := service.SubscriptionService.GetSubscriptionsByCourseId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptionsDto)
}

func PutRating(c *gin.Context) {
	var ratingDto dto.RatingDTO

	// Obtener el ID de la suscripción desde el parámetro de la URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Obtener el ID de la suscripción desde el parámetro de la URL
	id, _ := strconv.Atoi(c.Param("id"))

	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido en la URL"})
		return
	}

	if err := c.ShouldBindJSON(&ratingDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar el ID de la URL al DTO
	ratingDto.IDSubscription = id

	// Asignar el ID de la URL al DTO
	ratingDto.IDSubscription = id

	// Asigno el Id de la URL al DTO para que el servicio sepa qué actualizar
	ratingDto.IDSubscription = id
	ratingDto, err := service.SubscriptionService.PutRating(ratingDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ratingDto)
}

func PutComment(c *gin.Context) {
	var commentDto dto.CommentDTO

	// Obtener el ID de la suscripción desde el parámetro de la URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Obtener el ID de la suscripción desde el parámetro de la URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Parseo el body y el id
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido en la URL"})
		return
	}

	if err := c.ShouldBindJSON(&commentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar el ID de la URL al DTO
	commentDto.IDSubscription = id

	// Asignar el ID de la URL al DTO
	commentDto.IDSubscription = id

	commentDto.IDSubscription = id
	commentDto, err := service.SubscriptionService.PutComment(commentDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, commentDto)
}
