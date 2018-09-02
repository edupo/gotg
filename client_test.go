package gotg

import (
	"fmt"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

func TestContactList(t *testing.T) {
	sender := NewClient(&DefaultConfig)

	contacts, err := sender.ContactList()
	if err != nil {
		t.Error(err)
	}

	fmt.Print(fmt.Sprintf("Fetched %v contacts\n", len(contacts)))
}

func TestChannelList(t *testing.T) {
	sender := NewClient(&DefaultConfig)

	channels, err := sender.ChannelList(50, 0)
	if err != nil {
		t.Error(err)
	}

	fmt.Print(fmt.Sprintf("Fetched %v channels\n", len(channels)))
}

func TestMessage(t *testing.T) {
	sender := NewClient(&DefaultConfig)

	err := sender.SendMessage("Prueba", "This is a test message.")
	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {
	sender := NewClient(&DefaultConfig)
	logrus.SetLevel(logrus.DebugLevel)

	messages, err := sender.Search("Prueba",
		"This",
		10,
		0,
		time.Now().AddDate(0, 0, -1),
		time.Now())
	if err != nil {
		t.Error(err)
	}

	fmt.Print(fmt.Sprintf("Fetched %v messages\n", len(messages)))
}
