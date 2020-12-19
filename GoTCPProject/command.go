package main

type commandID int

//command IDs
const (
	CMD_USER commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

//command object to pass in user input to command functions
type command struct {
	id     commandID
	client *client
	args   []string
}
