package gotg

import (
	"fmt"
	"testing"
	"time"
)

func TestReceiver(t *testing.T) {
	receiver := NewReceiver(DefaultConfig)
	err := receiver.Start()
	if err != nil {
		t.Error(err)
	}
	go func() {
		time.Sleep(60 * time.Second)
		receiver.Stop()
	}()

	for {
		msg, ok := <-receiver.Channel
		if !ok {
			break
		}
		if msg.Type() == "message" {
			fmt.Println("Message received!")
		}
		fmt.Println(string(msg.Data))
	}

	err = receiver.tomb.Err()
	if err != nil {
		t.Error(err)
	}
}
