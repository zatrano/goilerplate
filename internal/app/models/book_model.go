package models

type Book struct {
	BaseModel
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
}

func (Book) TableName() string {
	return "books"
}
