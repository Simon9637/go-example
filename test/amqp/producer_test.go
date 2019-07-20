package amqp

import (
	"testing"
	"github.com/streadway/amqp"
	"fmt"
)

func TestProduce(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")
	if err != nil {
		t.Errorf("Failed to connnect the Rabbitmq, %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		t.Errorf("Failed to open a channel, %s", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("exchange_test", "topic", true, false, false, false, nil)
	if err != nil {
		t.Errorf("Failed to declare exchaneg, %s", err)
	}

	for i := 0; i < 10; i++ {
		err = ch.Publish("exchange_test", "routing_key_test", false, false, amqp.Publishing{ContentType: "application/json", Body: []byte(fmt.Sprintf("This is aaa test message : %d", i))})
		if err != nil {
			t.Errorf("Failed to publish a msg, %s", err)
		}
	}

}
