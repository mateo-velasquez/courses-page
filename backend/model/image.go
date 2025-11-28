package model

type Image struct {
	ImageId   int    `gorm:"column:image_id;primaryKey"`
	ImagePath string `gorm:"column:image_path;type:varchar(300); not null"`
}

type Images []Image
