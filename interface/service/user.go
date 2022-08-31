package service

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserServices struct {
	db []*User
}

type UserInterface interface {
	Register(user *User) string
	GetUser() []*User
	GetUserByID(i int) *User
	GetUserServer(writer http.ResponseWriter, request *http.Request)
	PostUserServer(writer http.ResponseWriter, request *http.Request)
}

func NewUserServices(user []*User) UserServices {
	return UserServices{
		db: user,
	}
}

func (u *UserServices) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Name + " Berhasil dibuat"
}

func (u *UserServices) GetUser() []*User {
	return u.db
}

func (u *UserServices) GetUserByID(i int) *User {
	for _, val := range u.db {
		if val.ID == i {
			return val
		}
	}
	return nil
}

func (u *UserServices) PostUserServer(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		name := request.FormValue("name")
		id, _ := strconv.Atoi(request.FormValue("id"))

		result := u.Register(&User{
			ID:   id,
			Name: name,
		})

		json.NewEncoder(writer).Encode(result)
		return
	}
}
func (u *UserServices) GetUserServer(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		result := u.GetUser()
		json.NewEncoder(writer).Encode(result)
		return
	}
}
