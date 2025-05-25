package app

import (
	"project/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	//USER
	router.POST("/user", controller.InsertUser)     //check
	router.GET("/users", controller.GetUsers)       //check
	router.GET("/user/:id", controller.GetUserById) //check
	router.POST("/login", controller.UserLogin)     //check

	//IMAGE
	router.POST("/course/image", controller.InsertImage) // Este que sea para guardar una relación de imagen con curso // no check
	router.GET("/image/:id", controller.GetImageById)    //check
	router.GET("/images", controller.GetImages)          //check

	//COURSE
	router.GET("/course/:id", controller.GetCourseById) //check
	router.GET("/courses", controller.GetCourses)       //check
	//router.GET("/search", controller.GetCoursesByNameAndCategory)
	router.GET("/search/:name", controller.GetCoursesByName)  //fail (No quiero implementar la búsqueda hasta que no corrija el CRUD)
	router.POST("/course", controller.InsertCourse)           //semi fail (no guarda las relaciones de las categorias)
	router.PUT("/course/:id", controller.PutCourseById)       //fail (me da un 500)
	router.DELETE("/course/:id", controller.DeleteCourseById) //fail (borra todos los cursos en lugar de uno sólo. raro)

	//SUBSCRIPTION
	//router.POST("/subscription", controller.InsertSuscription)
	//router.GET("/subscription/:id", controller.GetSubscriptionById)
	//router.GET("/subscriptions", controller.GetSubscriptions) //check
	//router.GET("/user/subscription/:id", controller.GetSubscriptionsByUserId)
	//router.GET("/course/subscription/:id", controller.GetSubscriptionsByCourseId)
	//router.PUT("/subscription/rating/:id", controller.PutRating)
	//router.PUT("/subscription/comment/:id", controller.PutComment)

	//CATEGORY
	router.POST("/category", controller.InsertCategory)     //check
	router.GET("/category/:id", controller.GetCategoryById) //check PERO TOMARLO CON PINZAS
	router.GET("/categories", controller.GetCategories)     //check

	//FILES

	log.Info("Finishing mappings configurations")
}
