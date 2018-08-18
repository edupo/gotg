package gotg

import "time"

type Peer struct {
	client *Client
	Searcher
	Messager
	Id        string `json:"id"`
	PrintName string `json:"print_name"`
	Flags     uint64 `json:"flags"`
	PeerType  string `json:"peer_type"`
	PeerId    uint64 `json:"peer_id"`
}

func (p *Peer) Search(pattern string, from time.Time, limit, offset uint64) ([]Message, error) {
	return p.client.Search(p, pattern, limit, offset, from, time.Now())
}

func (p *Peer) Message(msg string) error {
	return p.client.Message(p, msg)
}
