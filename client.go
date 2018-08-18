package gotg

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

// NewClient create and connect to a telegram-cli daemon. It must be started like:
// telegram-cli -P 4458 -W --json
// Where the most important flag is --json
func NewClient(address string) (*Client, error) {
	client := new(Client)
	err := client.connect(address)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) connect(address string) error {
	conn, err := net.Dial("tcp", address)
	c.conn = conn
	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)
	return err
}

func (c *Client) command(command ...string) error {
	_, err := c.writer.WriteString(strings.Join(command, " ") + "\n")
	if err != nil {
		return err
	}
	return c.writer.Flush()
}

func (c *Client) readAnswer() ([]byte, error) {
	// First line is formatted like:
	// ANSWER <Message size>
	str, err := c.reader.ReadString('\n')
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
	buf, err := c.reader.ReadBytes(byte('\n'))
	if err != nil {
		return nil, err
	}
	if size != int64(len(buf)) {
		return nil, errors.New(fmt.Sprintf("Unexpected answer: Size expected %v is not %v", size, len(buf)))
	}
	return buf, nil
}
