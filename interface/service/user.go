package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	GetUserGin(ctx *gin.Context)
	PostUserGin(ctx *gin.Context)
	GetUserByIdGin(ctx *gin.Context)
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

func (u *UserServices) GetUserGin(ctx *gin.Context) {
	users := u.GetUser()

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (u *UserServices) PostUserGin(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := u.Register(&User{
		ID:   user.ID,
		Name: user.Name,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func (u *UserServices) GetUserByIdGin(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	result := u.GetUserByID(id)

	if result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"user": result,
		})
	}
}
