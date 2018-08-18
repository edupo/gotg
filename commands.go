package gotg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// ContactList fetch from telegram
func (c *Client) ContactList() ([]Contact, error) {

	err := c.command("contact_list")
	if err != nil {
		return nil, err
	}

	buf, err := c.readAnswer()
	if err != nil {
		return nil, err
	}
	var contacts []Contact
	err = json.Unmarshal(buf, &contacts)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// ChannelList fetch from telegram
func (c *Client) ChannelList(limit, offset int) ([]Channel, error) {

	err := c.command("channel_list", strconv.Itoa(limit), strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	buf, err := c.readAnswer()
	if err != nil {
		return nil, err
	}
	var channels []Channel
	err = json.Unmarshal(buf, &channels)
	if err != nil {
		return nil, err
	}
	return channels, nil
}

// Message send a string to a peer
func (c *Client) Message(peer *Peer, msg string) error {
	return c.command("msg", peer.PrintName, msg)
}

// MainSession ask telegram-cli to send updates to this session
func (c *Client) MainSession() error {
	return c.command("main_session")
}

func (c *Client) Search(peer *Peer, pattern string, limit, offset uint64, from, to time.Time) ([]Message, error) {
	//search [peer] [limit] [from] [to] [offset] pattern
	err := c.command(fmt.Sprintf("search %v %v %v %v %v %v",
		peer.PrintName, limit, from.Unix(), to.Unix(), offset, pattern))
	if err != nil {
		return nil, err
	}

	buf, err := c.readAnswer()
	if err != nil {
		return nil, err
	}
	var messages []Message
	err = json.Unmarshal(buf, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
