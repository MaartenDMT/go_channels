package main

import (
	"fmt"
	"time"
)

func main() {

	main2()
	server := NewServer()
	server.Start()

	go func() {
		time.Sleep(2 * time.Second)
		close(server.quitch)

	}()

	// this blocks
	select {}
}

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
free:
	for {
		select {
		case msg := <-s.userch:
			fmt.Println(msg)
		case <-s.quitch:
			fmt.Println("server needs to quit")
			break free
		default:
		}
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func sendMessage(msgh chan<- string) {
	msgh <- "hello!"
}

func readMessage(msghr <-chan string) {
	msg := <-msghr
	fmt.Println(msg)
}
