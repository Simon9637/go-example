package amqp

import (
	"testing"
	"github.com/streadway/amqp"
	"fmt"
)

func TestConsumer(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5670/")
	if err != nil {
		t.Errorf("Failed to connnect the Rabbitmq, %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Errorf("Failed to open a channel, %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("queue_test", true, false, false, false, nil)
	if err != nil {
		t.Errorf("Failed to declare a queue, %s", err)
	}

	err = ch.QueueBind(q.Name, "routing_key_test", "exchange_test", false, nil)
	if err != nil {
		t.Errorf("Failed to bind a queue, %s", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}

}
