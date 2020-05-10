package queue

// Message is a concrete representation of the SQS message
type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}
