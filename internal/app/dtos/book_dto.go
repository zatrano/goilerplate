package dtos

type BookDTO struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
