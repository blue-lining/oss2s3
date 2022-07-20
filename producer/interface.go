package producer

type Producer interface {
	Put(string, timeout int) error
}
