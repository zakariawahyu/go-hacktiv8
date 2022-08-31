package service

type User struct {
	Name string
}

type UserServices struct {
}

type UserInterface interface {
	Register(user *User) string
}

func NewUserServices() UserInterface {
	return &UserServices{}
}

func (u *UserServices) Register(user *User) string {
	return user.Name + "Berhasil didaftarkan"
}
