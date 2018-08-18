package gotg

type Contact struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PrintName string `json:"print_name"`
	PeerType  string `json:"peer_type"`
	PeerId    uint64 `json:"peer_id"`
	Flags     uint64 `json:"flags"`
	When      string `json:"when"`
	Phone     string `json:"phone"`
}
