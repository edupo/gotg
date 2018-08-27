package gotg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// ContactList fetch from telegram
func (s *Sender) ContactList() ([]Contact, error) {
	buf, err := s.Send("contact_list")
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	err = json.Unmarshal(buf, &contacts)
	if err != nil {
		return nil, err
	}
	for _, contact := range contacts {
		contact.Peer.client = s
	}
	return contacts, nil
}

// ChannelList fetch from telegram
func (s *Sender) ChannelList(limit, offset int) ([]Channel, error) {
	buf, err := s.Send("channel_list", strconv.Itoa(limit), strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	var channels []Channel
	err = json.Unmarshal(buf, &channels)
	if err != nil {
		return nil, err
	}
	for _, channel := range channels {
		channel.Peer.client = s
	}
	return channels, nil
}

// SendMessage send a string to a peer
func (s *Sender) SendMessage(peer *Peer, msg string) error {
	buf, err := s.Send("msg", peer.PrintName, strconv.Quote(msg))
	print(buf)
	return err
}

// MainSession ask telegram-cli to send updates to this session
func (s *Sender) MainSession() error {
	return s.SendNoReceive("main_session")
}

func (s *Sender) Search(peer *Peer, pattern string, limit, offset uint64, from, to time.Time) ([]Message, error) {
	//search [peer] [limit] [from] [to] [offset] pattern
	buf, err := s.Send(fmt.Sprintf("search %v %v %v %v %v %v",
		peer.PrintName, limit, from.Unix(), to.Unix(), offset, pattern))
	if err != nil {
		return nil, err
	}

	var messages []Message
	err = json.Unmarshal(buf, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
