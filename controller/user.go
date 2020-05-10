package controller

import (
	"fmt"
	"net/http"
)

type UserController struct {
}

type userTo struct {
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Println("entrou GET")
		break
	case "POST":
		fmt.Println("entrou POST")
		break
	}

}

func NewUserController() *UserController {
	return &UserController{}
}
