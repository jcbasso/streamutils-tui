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
			line = filterNonGSM(line)

			err = client.handlePing(line)
			if err != nil {
				messages <- entities.Response{Error: fmt.Errorf("read error: %w", err)}
				close(messages)
			}

			if client.joinedSuccessfullyMessage(line) {
				messages <- entities.Response{
					Payload: entities.Payload{
						Action: "SUCCESS",
					},
				}
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

func (client *Client) joinedSuccessfullyMessage(message string) bool {
	re := regexp.MustCompile(`^:tmi\.twitch\.tv\s+001\s+\w+\s+:Welcome,\s*GLHF!$`)
	return re.MatchString(message)
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

func filterNonGSM(s string) string {
	gsmChars := map[rune]bool{
		'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
		'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true,
		'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true,
		'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
		'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true,
		'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true,
		'o': true, 'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true,
		'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
		'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true,
		'7': true, '8': true, '9': true,
		'!': true, '#': true, ' ': true, '"': true, '%': true, '&': true, '\'': true,
		'(': true, ')': true, '*': true, ',': true, '.': true, '?': true, '+': true,
		'-': true, '/': true, ';': true, ':': true, '<': true, '=': true, '>': true,
		'¡': true, '¿': true, '_': true, '@': true, '$': true, '£': true, '¥': true,
		'¤': true,
		'è': true, 'é': true, 'ù': true, 'ì': true, 'ò': true, 'Ç': true, 'Ø': true,
		'ø': true, 'Æ': true, 'æ': true, 'ß': true, 'É': true, 'Å': true, 'å': true,
		'Ä': true, 'Ö': true, 'Ñ': true, 'Ü': true, '§': true, 'ä': true, 'ö': true,
		'ñ': true, 'ü': true, 'à': true,
		'Δ': true, 'Φ': true, 'Ξ': true, 'Γ': true, 'Ω': true, 'Π': true, 'Ψ': true,
		'Σ': true, 'Θ': true, 'Λ': true, '\n': true, '\r': true, '\t': true,
		0x0C: true, // Form Feed
		0x1B: true, // Escape
	}

	var res strings.Builder // Use strings.Builder for efficient string concatenation
	for _, r := range s {
		if _, ok := gsmChars[r]; ok {
			res.WriteRune(r)
		}
	}
	return strings.TrimSpace(res.String())
}
