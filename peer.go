package gotg

import "time"

type Peer struct {
	config *Config
	Searcher
	Messager
	Id        string `json:"id"`
	PrintName string `json:"print_name"`
	Flags     uint64 `json:"flags"`
	PeerType  string `json:"peer_type"`
	PeerId    uint64 `json:"peer_id"`
}

func (p *Peer) Search(pattern string, from time.Time, limit, offset uint64) ([]Message, error) {
	return NewClient(p.config).Search(p, pattern, limit, offset, from, time.Now())
}

func (p *Peer) SendMessage(msg string) error {
	return NewClient(p.config).SendMessage(p, msg)
}

func NewPeer(name string, config *Config) *Peer {
	return &Peer{
		config:    config,
		PrintName: name,
	}
}
