package gotg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// ContactList fetch from telegram
func (c *Client) ContactList() ([]Contact, error) {
	buf, err := c.Send("contact_list")
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	err = json.Unmarshal(buf, &contacts)
	if err != nil {
		return nil, err
	}
	// Each fetched peer knows it'c config
	for _, contact := range contacts {
		contact.Peer.config = c.config
	}
	return contacts, nil
}

// ChannelList fetch from telegram
func (c *Client) ChannelList(limit, offset int) ([]Channel, error) {
	buf, err := c.Send("channel_list", strconv.Itoa(limit), strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	var channels []Channel
	err = json.Unmarshal(buf, &channels)
	if err != nil {
		return nil, err
	}
	// Each fetched peer knows it'c config
	for _, channel := range channels {
		channel.Peer.config = c.config
	}
	return channels, nil
}

// SendMessage send a string to a peer
func (c *Client) SendMessage(peer string, msg string) error {
	buf, err := c.Send("msg", peer, strconv.Quote(msg))
	if err != nil {
		return err
	}
	return checkSuccess(buf)
}

// MainSession ask telegram-cli to send updates to this session.
// --- This function does not close the connection! ---
func (c *Client) MainSession() error {
	err := c.connect()
	if err != nil {
		return err
	}
	err = c.sendCommand("main_session")
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Search(peer string, pattern string, limit, offset uint64, from, to time.Time) ([]Message, error) {
	//search [peer] [limit] [from] [to] [offset] pattern
	buf, err := c.Send(fmt.Sprintf("search %v %v %v %v %v %v",
		peer, limit, from.Unix(), to.Unix(), offset, pattern))
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

func (c *Client) GetSelf() (Peer, error) {
	buf, err := c.Send("get_self")
	if err != nil {
		return Peer{}, err
	}
	var peer Peer
	err = json.Unmarshal(buf, &peer)
	return peer, err
}
