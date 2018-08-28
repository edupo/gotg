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

type Client struct {
	config *Config
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewClient(config *Config) *Client {
	sender := &Client{
		config: config,
	}
	return sender
}

func (c *Client) connect() error {
	if c.conn != nil {
		return nil
	}
	// TODO use DialTimeout
	conn, err := net.Dial("tcp", c.config.Address)
	if err != nil {
		return err
	}
	c.conn = conn
	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)
	return nil
}

func (c *Client) disconnect() error {
	err := c.writer.Flush()
	c.writer = nil
	c.reader = nil

	if err2 := c.conn.Close(); err2 != nil {
		err = err2
	}
	c.conn = nil

	return err
}

// Send a command and wait for an answer. Return the answer string.
func (c *Client) Send(command ...string) ([]byte, error) {
	err := c.connect()
	if err != nil {
		return nil, err
	}
	defer c.disconnect()

	err = c.sendCommand(command...)
	if err != nil {
		return nil, err
	}

	bytes, err := c.readAnswer()
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *Client) sendCommand(command ...string) error {
	cmd := strings.Join(command, " ")
	logrus.WithField("command", cmd).Debug("Telegram command sent")
	_, err := c.writer.WriteString(cmd + "\n")
	if err != nil {
		return err
	}
	return c.writer.Flush()
}

func (c *Client) readAnswer() ([]byte, error) {
	size, err := c.readHeader()
	if err != nil {
		return nil, err
	}
	buf, err := c.readBytes(size)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Client) readHeader() (int64, error) {
	// First received line of any message looks like:
	// ANSWER <Message size>
	str, err := c.reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	fields := strings.Fields(str)
	if fields[0] != "ANSWER" {
		return -1, errors.New("Unexpected answer: " + str)
	}
	size, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return -1, err
	}
	return size, nil
}

func (c *Client) readBytes(size int64) ([]byte, error) {
	bytes, err := c.reader.ReadBytes(byte('\n'))
	if err != nil {
		return nil, err
	}
	if size != int64(len(bytes)) {
		return nil, errors.New(fmt.Sprintf("Unexpected answer: Size expected %v is not %v", size, len(bytes)))
	}
	return bytes, nil
}
