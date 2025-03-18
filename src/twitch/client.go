package twitch

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"regexp"
	"streamutils-tui/src/entities"
	"strings"
)

type Client struct {
	accessToken string
	username    string
	ircServer   string
	conn        net.Conn
}

func New(accessToken string, username string, ircServer string) *Client {
	return &Client{
		accessToken: accessToken,
		username:    username,
		ircServer:   ircServer,
	}
}

func (client *Client) Join(channel string) error {
	conn, err := net.Dial("tcp", client.ircServer)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(conn, "PASS oauth:%s\r\n", client.accessToken)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(conn, "NICK %s\r\n", client.username)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(conn, "JOIN #%s\r\n", channel)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(conn, "CAP REQ :twitch.tv/membership twitch.tv/tags twitch.tv/commands twitch.tv/badges\r\n")
	if err != nil {
		return err
	}

	client.conn = conn
	return nil
}

func (client *Client) StreamChat() chan entities.Response {
	messages := make(chan entities.Response)

	go func() {
		reader := bufio.NewReader(client.conn)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF || strings.Contains(err.Error(), "use of closed network connection") {
					messages <- entities.Response{Error: fmt.Errorf("connection closed: %w", err)}
				} else {
					messages <- entities.Response{Error: fmt.Errorf("read error: %w", err)}
				}
				close(messages)
				return
			}

			// Removing \r\n chars
			line = strings.Replace(line, "\n", "", -1)
			line = strings.Replace(line, "\r", "", -1)

			err = client.handlePing(line)
			if err != nil {
				messages <- entities.Response{Error: fmt.Errorf("read error: %w", err)}
				close(messages)
			}

			if client.isSystemMessage(line) {
				continue
			}

			line, tagsString, err := SeparateTags(line)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tags := ParseTags(tagsString)

			packet, err := ParsePacket(line)
			if err != nil {
				// TODO: check what to do
				// fmt.Println(err)
				continue
			}

			messages <- entities.Response{Tags: tags, Payload: packet}
		}
	}()

	return messages
}

func (client *Client) handlePing(message string) error {
	pingRegex := regexp.MustCompile(`^PING\s+(:.+)$`)
	match := pingRegex.FindStringSubmatch(message)
	if len(match) > 0 {
		pingData := match[1]
		pong := fmt.Sprintf("PONG %s\r\n", pingData)
		_, err := fmt.Fprintf(client.conn, pong)
		if err != nil {
			return fmt.Errorf("error handling PING: %w", err)
		}
	}
	return nil
}

func (client *Client) isSystemMessage(message string) bool {
	re := regexp.MustCompile(`^:tmi\.twitch\.tv.*$`)
	match := re.FindStringSubmatch(message)
	if len(match) > 0 {
		return true
	}
	return false
}

func (client *Client) Read() (string, error) {
	reader := bufio.NewReader(client.conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		if err.Error() == "EOF" || strings.Contains(err.Error(), "use of closed network connection") {
			return "", fmt.Errorf("connection closed")
		}
		return "", err
	}
	return line, nil
}
