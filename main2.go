package main

import (
	"fmt"
	"time"
)

func main2() {
	userch := make(chan string, 1)

	go func() {
		time.Sleep(5 * time.Second)
		// close(userch)
		userch <- "bob"

	}()

	username, ok := <-userch

	if !ok {
		fmt.Println("userchan closed returning")
		return
	}

	if len(username) == 0 {
		panic(" ai ai ai user not good")
	}

	fmt.Println(username)
}
