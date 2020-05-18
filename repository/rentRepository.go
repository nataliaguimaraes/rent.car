package repository

import (
	"fmt"
	"github.com/gocql/gocql"
)

type Message struct {
	From string
	To   string
	Body string
}

type Rent struct {
	Id gocql.UUID
	Message Message
}

func CreateRent(rent Rent){

	uuid, _ := gocql.RandomUUID()

	if err := Session.Query("INSERT INTO rent (rent_id,msg_to, msg_from, msg_body) values (?,?,?,?)", uuid, rent.Message.To, rent.Message.From, rent.Message.From).Exec()
	err != nil {
		fmt.Println("Error saving rent, error:", err)
	} else {
		fmt.Println("Rent message saved!")
	}

}

func GetAllRents() []Rent {

	var rents []Rent
	row := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM rent").Iter()
	for iter.MapScan(row) {
		rents = append(rents, Rent{
			Message: Message{
				To: row["msg_to"].(string),
				From: row["msg_from"].(string),
				Body: row["msg_body"].(string)},
			Id: row["rent_id"].(gocql.UUID),
		})
		row = map[string]interface{}{}
	}

	return rents

}

