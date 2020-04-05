package interfaces

type UserInterface interface {
	GetUsername() string
	GetSurname() string
}


func NewUser() UserInterface {
	return &User{
		"Jeffrey",
		"Lim",
	}
}
