package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//client type
type client struct {
	conn     net.Conn
	user     string
	room     *room
	commands chan<- command
}

//take client input from terminal
func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		//sense for command input
		switch cmd {
		case "/user":
			c.commands <- command{
				id:     CMD_USER,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

//client error
func (c *client) err(err error) {
	c.conn.Write([]byte("err: " + err.Error() + "\n"))
}

//message function and formatting
func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + "[" + msg + "]" + "\n"))
}
