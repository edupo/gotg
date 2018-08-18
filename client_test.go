package gotg

import (
	"fmt"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

var (
	test_address = "127.0.0.1:4458"
)

func TestContactList(t *testing.T) {
	client, err := NewClient(test_address)
	if err != nil {
		t.Error(err)
	}

	contacts, err := client.ContactList()
	if err != nil {
		t.Error(err)
	}

	fmt.Print(fmt.Sprintf("Fetched %v contacts\n", len(contacts)))
}

func TestChannelList(t *testing.T) {
	client, err := NewClient(test_address)
	if err != nil {
		t.Error(err)
	}

	channels, err := client.ChannelList(50, 0)
	if err != nil {
		t.Error(err)
	}

	fmt.Print(fmt.Sprintf("Fetched %v channels\n", len(channels)))
}

func TestMessage(t *testing.T) {
	client, err := NewClient(test_address)
	if err != nil {
		t.Error(err)
	}

	channel := Channel{
		Peer: Peer{
			PrintName: "Prueba",
		},
	}

	err = client.Message(&channel.Peer, "This is a test message.")
	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	client, err := NewClient(test_address)
	if err != nil {
		t.Error(err)
	}

	channel := Channel{
		Peer: Peer{
			PrintName: "Prueba",
		},
	}

	messages, err := client.Search(&channel.Peer,
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
