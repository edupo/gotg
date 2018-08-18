package gotg

import (
	"encoding/json"
	"strconv"
)

// ContactList fetch from telegram
func (c *Client) ContactList() ([]Contact, error) {
	c.command("contact_list")
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
	c.command("channel_list", strconv.Itoa(limit), strconv.Itoa(offset))
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
