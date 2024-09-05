package main

type Producer interface {
	sendMessage(topic string, message string) error
	addTopic(topic string)
	addConsumer(consumer Consumer)
	notifyConsumers(topic string)
}

type Twitter struct {
}

func (t *Twitter) sendMessage(topic string, message string) error {
	return nil
}

type Consumer interface {
	consume(topic string) (string, error)
}

type db struct {
}

func (db *db) consume(topic string) (string, error) {
	return "", nil
}

type Topic struct {
	name     string
	messages []string
}

type Queue struct {
	// topics    []Topic
	// producers []Producer
	// consumers []Consumer
	topic    Topic
	producer Producer
	consumer Consumer
}

func main() {
	topic := Topic{name: "transactions", messages: []string{}}
	twitter := Twitter{}
	queue := Queue{topic: topic, producer: twitter, consumer: db{}}
}
