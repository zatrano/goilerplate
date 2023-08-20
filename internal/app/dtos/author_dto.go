package dtos

type AuthorDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateAuthorDTO struct {
	Name string `json:"name"`
}

type UpdateAuthorDTO struct {
	Name string `json:"name"`
}
