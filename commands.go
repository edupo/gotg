package gotg

import "encoding/json"

// ContactList from telegram
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

// Message send a string to a peer
func (c *Client) Message(peer *Peer, msg string) error {
	return c.command("msg", peer.PrintName, msg)
}
