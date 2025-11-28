package client

import (
	"project/model"

	log "github.com/sirupsen/logrus"
)

func InsertImage(image model.Image) model.Image {

	result := Db.Debug().Create(&image)

	if result.Error != nil {
		log.Error("Failed to insert image")
		image.ImageId = -1
		return image
	}

	log.Debug("Image created:", image.ImageId)
	return image
}

func GetImageById(id int) model.Image {
	var image model.Image

	Db.Where("image_id = ?", id).First(&image)
	log.Debug("Image: ", image)

	return image
}

func GetImages() model.Images {
	var images model.Images
	Db.Find(&images)

	log.Debug("Images: ", images)

	return images
}

func GetIdMayor() int {
	var image model.Image

	// Buscamos un solo registro, ordenado del más grande al más chico
	result := Db.Order("image_id desc").First(&image)

	// Si hay error (ej. tabla vacía), devolvemos 0.
	// El controlador sumará 1 y la primera imagen será Image-1.jpg (correcto)
	if result.Error != nil {
		log.Error("No se encontró ID mayor (puede que la tabla esté vacía): ", result.Error)
		return 0
	}

	log.Debug("Maximum ID found: ", image.ImageId)
	return image.ImageId
}
