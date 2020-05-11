package queue

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// Queue provides the ability to handle SQS messages.
type Queue struct {
	Client sqsiface.SQSAPI
	URL    string
}

func (q *Queue) ConsumeMessage(waitTimeout int64) ([]Message, error) {

	params := sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(q.URL),
		MaxNumberOfMessages: aws.Int64(10),
	}

	result, err := q.Client.ReceiveMessage(&params)

	if err != nil {
		fmt.Printf("Error retrieving message: %v\n", err)
	}

	msgs := make([]Message, len(result.Messages))
	for i, msg := range result.Messages {

		parsedMsg := Message{}
		bodyMsg := aws.StringValue(msg.Body)

		if err := json.Unmarshal([]byte(bodyMsg), &parsedMsg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal message, %v \n", err)
		}
		msgs[i] = parsedMsg

		// msgs[i] = Message{bodyMsg}

		_, err := q.Client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &q.URL,
			ReceiptHandle: msg.ReceiptHandle,
		})

		if err != nil {
			fmt.Printf("Delete Error: %v\n", err)
		}
	}

	return msgs, nil
}
