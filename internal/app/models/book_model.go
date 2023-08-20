package models

type Book struct {
	BaseModel
	Title    string `gorm:"not null" json:"title"`
	AuthorID uint   `gorm:"not null" json:"author_id"`
	Author   Author `gorm:"foreignKey:AuthorID" json:"author"`
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) Validate() []string {
	var errs []string

	if len(b.Title) == 0 {
		errs = append(errs, "Title cannot be empty")
	}

	return errs
}
