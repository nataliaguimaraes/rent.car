package repository

import (
	"fmt"
	"github.com/gocql/gocql"
)

type Car struct {
	Id    gocql.UUID
	Plate string
}

func CreateCar(car Car) {
	fmt.Println("Creating a new Car", car)
	if err := Session.Query("INSERT INTO car (car_id, plate) VALUES (?, ?)", car.Id, car.Plate).Exec();
		err != nil {
		fmt.Println("Error while inserting Car", err)
	}
}

func GetAllCars() []Car {
	fmt.Println("Getting all Cars")
	var cars []Car
	row := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM car").Iter();
	for iter.MapScan(row) {
		cars = append(cars, Car{
			Id:    row["car_id"].(gocql.UUID),
			Plate: row["plate"].(string),
		})
		row = map[string]interface{}{}
	}

	return cars
}