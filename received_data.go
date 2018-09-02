package gotg

import "encoding/json"

type ReceivedData struct {
	Event
	Data []byte
}

func NewReceivedData(data []byte) (*ReceivedData, error) {
	var rd = ReceivedData{
		Data: data,
	}
	err := json.Unmarshal(data, &rd)
	return &rd, err
}

func (rd *ReceivedData) Type() string {
	return rd.Event.Event
}

func (rd *ReceivedData) ToMessage() (Message, bool) {
	var m Message
	if rd.Type() != "message" {
		return m, false
	}
	err := json.Unmarshal(rd.Data, &m)
	return m, err == nil
}
