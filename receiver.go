package gotg

import (
	"errors"

	"bytes"

	"gopkg.in/tomb.v2"
)

var (
	lineBreak = []byte{10}
)

type Receiver struct {
	client  *Client
	Channel chan *ReceivedData
	tomb    tomb.Tomb
	config  Config
}

func NewReceiver(config Config) (*Receiver, error) {
	receiver := &Receiver{
		config: config,
	}
	receiver.client = NewClient(&receiver.config)

	err := receiver.connect()
	if err != nil {
		return nil, err
	}

	receiver.tomb.Go(receiver.loop)

	return receiver, nil
}

func (r *Receiver) Stop() error {
	r.tomb.Kill(nil)
	return r.tomb.Wait()
}

func (r *Receiver) connect() error {
	r.Channel = make(chan *ReceivedData)
	err := r.client.connect()
	if err != nil {
		return err
	}

	err = r.client.MainSession()
	if err != nil {
		return err
	}

	return nil
}

func (r *Receiver) disconnect() error {
	close(r.Channel)
	return r.client.disconnect()
}

func (r *Receiver) receive() (*ReceivedData, error) {
	bts, err := r.client.readAnswer()
	if err != nil {
		return nil, err
	}
	bts2, err := r.client.readBytes(1)
	if err != nil {
		return nil, err
	}
	// Once a message is received an empty message is expected.
	if bytes.Compare(bts2, lineBreak) != 0 {
		return nil, errors.New("Malformed received message: Didn't end with double line break.")
	}
	return NewReceivedData(bts)
}

func (r *Receiver) loop() error {
	defer r.disconnect()
	for {
		answer, err := r.receive()
		if err != nil {
			continue
		}
		select {
		case r.Channel <- answer:
		case <-r.tomb.Dying():
			return nil
		default:
		}
	}
}
