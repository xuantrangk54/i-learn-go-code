package main

import (
	"fmt"

	"github.com/evenskylearngo/greetings/greetings"
)

func main() {
	fmt.Println(greetings.Hello("Go Developer"))
	fmt.Println(greetings.Goodbye("trang"))
	fmt.Println(greetings.Plus(1, 2))
	_, err := greetings.Hello("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
