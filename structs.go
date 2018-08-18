package gotg

type Peer struct {
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
