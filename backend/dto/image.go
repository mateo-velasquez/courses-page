package dto

type ImageDTO struct {
	IDImage   int    `json:"id"`
	ImagePath string `json:"path"`
}

type ImagesDTO []ImageDTO
