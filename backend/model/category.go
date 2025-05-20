package model

type Category struct {
	IDCategory   int    `gorm:"column:id_category;primaryKey;autoIncrement"`
	CategoryName string `gorm:"column:category_name;type:varchar(100);unique;not null"`
}

// TableName especifica el nombre exacto de la tabla en la base de datos
func (Category) TableName() string {
	return "categories"
}

type Categories []Category
