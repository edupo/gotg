package gotg

type Peer struct {
	client    *Client
	Id        string `json:"id"`
	PrintName string `json:"print_name"`
	Flags     uint64 `json:"flags"`
	PeerType  string `json:"peer_type"`
	PeerId    uint64 `json:"peer_id"`
}

type Contact struct {
	Peer
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	When      string `json:"when"`
	Phone     string `json:"phone"`
}

type Channel struct {
	Peer
	Title       string `json:"title"`
	AdminsCount uint64 `json:"admins_count"`
	KickedCount uint64 `json:"kicked_count"`
}

type Event struct {
	Event string `json:"event"`
}

type Message struct {
	Event
	Date    uint64 `json:"date"`
	Flags   uint64 `json:"flags"`
	Id      string `json:"id"`
	Out     bool   `json:"out"`
	Service bool   `json:"service"`
	Unread  bool   `json:"unread"`
	Text    string `json:"text"`
	From    Peer   `json:"from"`
	To      Peer   `json:"to"`
}
