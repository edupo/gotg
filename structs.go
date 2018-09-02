package gotg

type Event struct {
	Event string `json:"event"`
}

type Result struct {
	Result string `json:"result"`
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
