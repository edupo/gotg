package gotg

import (
	"fmt"
	"testing"
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
