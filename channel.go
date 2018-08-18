package gotg

import "time"

type Channel struct {
	Peer
	Title       string `json:"title"`
	AdminsCount uint64 `json:"admins_count"`
	KickedCount uint64 `json:"kicked_count"`
}

func (c *Channel) Search(pattern string, from time.Time, limit, offset uint64) ([]Message, error) {
	return c.Peer.Search(pattern, from, limit, offset)
}

func (c *Channel) Message(msg string) error {
	return c.Peer.Message(msg)
}
