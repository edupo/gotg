package gotg

import "encoding/json"

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
