package consumer

type Consumer interface {
	Get(timeout int) (string, error)
	Ack(string) error
}
