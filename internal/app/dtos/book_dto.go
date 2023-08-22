package dtos

type BookDTO struct {
	ID     uint      `json:"id"`
	Title  string    `json:"title"`
	Author AuthorDTO `json:"author"`
}

type CreateBookDTO struct {
	Title  string    `json:"title"`
	Author AuthorDTO `json:"author"`
}

type UpdateBookDTO struct {
	Title  string    `json:"title"`
	Author AuthorDTO `json:"author"`
}
