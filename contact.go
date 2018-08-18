package gotg

import "time"

type Contact struct {
	Peer
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	When      string `json:"when"`
	Phone     string `json:"phone"`
}

func (c *Contact) Search(pattern string, from time.Time, limit, offset uint64) ([]Message, error) {
	return c.Peer.Search(pattern, from, limit, offset)
}

func (c *Contact) Message(msg string) error {
	return c.Peer.Message(msg)
}
