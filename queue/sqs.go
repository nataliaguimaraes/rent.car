package queue

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"rent.car/repository"
)

func Start() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")})

	if err != nil {
		fmt.Printf("Error creating session: %s \n", err)
	}

	q := Queue{
		Client: sqs.New(sess),
		URL:    "https://sqs.us-west-2.amazonaws.com/544340882321/reservation.fifo",
	}

	fmt.Println(q.URL)
	for {

		msgs, err := q.ConsumeMessage(0)

		if err != nil {
			fmt.Printf("error to retrieve message: %s \n", err)
			continue
		}

		for _, msg := range msgs {

			rent := repository.Rent{
				Message: repository.Message{
					To: msg.To,
					From: msg.From,
					Body: msg.Body,
				},
			}
			repository.CreateRent(rent)
		}

		rents := repository.GetAllRents()
		fmt.Println(rents)
	}
}
