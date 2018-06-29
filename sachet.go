package sachet

type Response struct {
	Status int
	Body   interface{}
}

type Provider interface {
	Send(message Message) (Response, error)
}

type Message struct {
	To   []string
	From string
	Text string
}
