package repository

import (
	"fmt"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "rent_car_keyspace"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")

	err = Session.Query("CREATE KEYSPACE IF NOT EXISTS rent_car_keyspace " +
		"WITH REPLICATION={'class': 'SimpleStrategy', 'replication_factor': 1};").Exec()
	if err != nil {
		fmt.Println("error to retrieve cassandra keyspace, error:", err)
		panic(err)
	}

	err = Session.Query("CREATE TABLE IF NOT EXISTS car(car_id uuid PRIMARY KEY, plate text);").Exec()
	if err !=  nil {
		fmt.Println("error to create CAR table, error:", err)
		panic(err)
	}

	err = Session.Query("CREATE TABLE IF NOT EXISTS rent(rent_id uuid PRIMARY KEY, msg_to text, msg_from text, msg_body text);").Exec()
	if err !=  nil {
		fmt.Println("error to create RENT table, error:", err)
		panic(err)
	}
}