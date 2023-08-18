package models

type Book struct {
	BaseModel
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) Validate() []string {
	var errs []string

	if len(b.Title) == 0 {
		errs = append(errs, "Title cannot be empty")
	}

	if len(b.Author) == 0 {
		errs = append(errs, "Author cannot be empty")
	}

	return errs
}
