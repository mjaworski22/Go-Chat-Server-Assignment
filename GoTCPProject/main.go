package main

import (
	"log"
	"net"
)

func main() {
	//create a new server and run it
	s := newServer()
	go s.run()

	//add listener to sense connection close, open, and failure
	//Make the server usable only with a TCP connection on Port 8888
	listener, err := net.Listen("tcp", ":8888")

	//Server startup fail message if an error is sensed
	if err != nil {
		log.Fatalf("Server start failed: %s", err.Error())
	}

	//if there are no errors then close the listener and tell the user that a server was started
	defer listener.Close()
	log.Printf("Started Server")
	log.Printf("Welcome to the chat server. Go open up any amount of terminals.")
	log.Printf("Type 'telnet localhost 8888' in each of them and press enter to activate them")
	log.Printf("Then you can start executing commands with them")
	log.Printf("To use the 'USER' command, which names your client, type '/user' ")
	log.Printf("To use the 'JOIN' command, which puts you in a room, type '/join #your_desired_room_name' ")
	log.Printf("To use the 'ROOMS' command, which tells you the available rooms to join, type '/rooms' ")
	log.Printf("To use the 'MSG' command, which sends a message, type '/msg' ")
	log.Printf("To use the 'QUIT' command, which disconnects you from the server, type '/quit' ")
	//Listen for errors while establishing a connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection Acceptance Failed: %s", err.Error())
			continue
		}
		//establish connection with client when connection acceptance no longer fails
		go s.newClient(conn)
	}
}
