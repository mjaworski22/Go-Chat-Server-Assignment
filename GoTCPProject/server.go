package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

//server object
type server struct {
	rooms    map[string]*room
	commands chan command
}

//function for new server
func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

//define the commands
func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_USER:
			s.user(cmd.client, cmd.args[1])
		case CMD_JOIN:
			s.join(cmd.client, cmd.args[1])
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

//take on a new client
func (s *server) newClient(conn net.Conn) {
	log.Printf("A new client has joined: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		user:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

//user function
func (s *server) user(c *client, user string) {
	c.user = user
	c.msg(fmt.Sprintf("Your username is:  %s", user))
}

//join function
func (s *server) join(c *client, roomName string) {
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)
	c.room = r

	r.broadcast(c, fmt.Sprintf("%s joined the room.", c.user))

	c.msg(fmt.Sprintf("You have joined the room: %s", roomName))
}

//rooms function
func (s *server) listRooms(c *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("These are the available rooms: %s", strings.Join(rooms, ", ")))
}

//message function
func (s *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:len(args)], " ")
	c.room.broadcast(c, c.user+": "+msg)
}

//quit function
func (s *server) quit(c *client) {
	log.Printf("A client has left the chat: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("Bye")
	c.conn.Close()
}

//exit current room function
func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		oldRoom := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.user))
	}
}
