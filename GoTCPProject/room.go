package main

import (
	"net"
)

//room object, assigning name and member list
type room struct {
	name    string
	members map[net.Addr]*client
}

//broadcasts message to everyone in room thru server file
func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
