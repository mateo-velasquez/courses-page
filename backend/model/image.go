package model

type Image struct {
	IDImage   int    `gorm:"column:id_image;primaryKey;autoIncrement"`
	ImagePath string `gorm:"column:image_path;type:varchar(300);unique;not null"`
}

// TableName especifica el nombre exacto de la tabla en la base de datos
func (Image) TableName() string {
	return "images"
}

type Images []Image
