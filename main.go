//sempre tem que existir um main.go na raiz
package main

import (
	"fmt"

	"net/http"

	"rent.car/controller"
	"rent.car/queue"
)

type X struct {
	Value int    `json:valor`
	Type  string `json:tipo`
}

func main() {

	fmt.Println("oi")

	go queue.Start()

	handleRequests()
}

func handleRequests() {

	uc := controller.NewUserController()
	http.Handle("/user", *uc)

	http.ListenAndServe(":8080", nil)
}
