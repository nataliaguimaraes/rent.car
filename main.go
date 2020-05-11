//sempre tem que existir um main.go na raiz
package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"rent.car/repository"

	"net/http"

	"rent.car/controller"
	"rent.car/queue"
)

func main() {

	fmt.Println("oi")
	uuid, _ := gocql.RandomUUID()
	car := repository.Car{Id: uuid, Plate: "xxxx"}
	repository.CreateCar(car)
	fmt.Println(repository.GetAllCars())

	go queue.Start()

	handleRequests()
}

func handleRequests() {

	uc := controller.NewUserController()
	http.Handle("/user", *uc)

	http.ListenAndServe(":8080", nil)
}
