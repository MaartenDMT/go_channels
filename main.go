package main

import "fmt"

func main() {
	userch := make(chan string)
	
	go func() {
		userch <- "Bob" //blocking
	}()

	user := <-userch

	fmt.Println(user)
}
