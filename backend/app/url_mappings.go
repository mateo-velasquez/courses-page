package app

import (
	//"project/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	//USER
	//router.POST("/user", controller.InsertUser)     //check
	//router.GET("/users", controller.GetUsers)       //check
	//router.GET("/user/:id", controller.GetUserById) //check
	//router.POST("/login", controller.UserLogin)     //check

	//IMAGE
	//router.POST("/course/image", controller.InsertImage)
	//router.GET("/course/image/id", controller.GetImageById)
	//router.GET("/images", controller.GetImages) //check

	//COURSE
	//router.GET("/course/:id", controller.GetCourseById)
	//router.GET("/courses", controller.GetCourses) //check
	//router.GET("/search/:name", controller.GetCoursesByName)
	//router.POST("/course", controller.InsertCourse)
	//router.PUT("/course/:id", controller.PutCourseById)
	//router.DELETE("/course/:id", controller.DeleteCourseById)

	//SUBSCRIPTION
	//router.POST("/subscription", controller.InsertSuscription)
	//router.GET("/subscription/:id", controller.GetSubscriptionById)
	//router.GET("/subscriptions", controller.GetSubscriptions) //check
	//router.GET("/user/subscription/:id", controller.GetSubscriptionsByUserId)
	//router.GET("/course/subscription/:id", controller.GetSubscriptionsByCourseId)
	//router.PUT("/subscription/rating/:id", controller.PutRating)
	//router.PUT("/subscription/comment/:id", controller.PutComment)

	//CATEGORY
	//router.POST("/category", controller.InsertCategory)     //check
	//router.GET("/category/:id", controller.GetCategoryById) //check PERO TOMARLO CON PINZAS
	//router.GET("/categories", controller.GetCategories)     //check

	log.Info("Finishing mappings configurations")
}
