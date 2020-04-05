package interfaces

type User struct {
	Username string
	Surname  string
}

func (u *User) GetUsername() string {
	return u.Username;
}

func (u *User) GetSurname() string {
	return u.Surname;
}
