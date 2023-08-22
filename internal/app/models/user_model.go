package models

type User struct {
	BaseModel
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Validate() []string {
	var errs []string

	if len(u.Name) == 0 {
		errs = append(errs, "Name cannot be empty")
	}

	return errs
}
