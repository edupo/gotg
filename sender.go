package gotg

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
)

type Sender struct {
	config Config
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewSender(config Config) *Sender {
	sender := &Sender{
		config: config,
	}
	return sender
}

func (s *Sender) connect() error {
	conn, err := net.Dial("tcp", s.config.Address)
	if err != nil {
		return err
	}
	s.conn = conn
	s.reader = bufio.NewReader(s.conn)
	s.writer = bufio.NewWriter(s.conn)
	return nil
}

func (s *Sender) disconnect() error {
	err := s.writer.Flush()
	s.writer = nil
	s.reader = nil

	if err2 := s.conn.Close(); err2 != nil {
		err = err2
	}
	s.conn = nil

	return err
}

// SendNoReceive a command without waiting for answer. We assume the command will have a possitive answer.
func (s *Sender) SendNoReceive(command ...string) error {
	err := s.connect()
	if err != nil {
		return err
	}
	defer s.disconnect()

	return s.sendCommand(command...)
}

// Send a command and wait for an answer. Return the answer string.
func (s *Sender) Send(command ...string) ([]byte, error) {
	err := s.connect()
	if err != nil {
		return nil, err
	}
	defer s.disconnect()

	err = s.sendCommand(command...)
	if err != nil {
		return nil, err
	}

	buf, err := s.readAnswer()
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Sender) sendCommand(command ...string) error {
	cmd := strings.Join(command, " ")
	logrus.WithField("command", cmd).Debug("Telegram command sent")
	_, err := s.writer.WriteString(cmd + "\n")
	if err != nil {
		return err
	}
	return s.writer.Flush()
}

func (s *Sender) readAnswer() ([]byte, error) {
	// First line is formatted like:
	// ANSWER <SendMessage size>
	str, err := s.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	fields := strings.Fields(str)
	if fields[0] != "ANSWER" {
		return nil, errors.New("Unexpected answer: " + str)
	}
	size, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return nil, err
	}

	// Second line is the answer message itself. We honor the size in specified in the header.
	buf, err := s.reader.ReadBytes(byte('\n'))
	if err != nil {
		return nil, err
	}
	if size != int64(len(buf)) {
		return nil, errors.New(fmt.Sprintf("Unexpected answer: Size expected %v is not %v", size, len(buf)))
	}
	return buf, nil
}
