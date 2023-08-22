package dtos

type BookDTO struct {
	ID     uint      `json:"id"`
	Title  string    `json:"title" validate:"required"`
	Author AuthorDTO `json:"author" validate:"required"`
}

type CreateBookDTO struct {
	Title  string    `json:"title" validate:"required"`
	Author AuthorDTO `json:"author" validate:"required"`
}

type UpdateBookDTO struct {
	Title  string    `json:"title" validate:"required"`
	Author AuthorDTO `json:"author" validate:"required"`
}
