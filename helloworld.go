package main

import (
	"fmt"
	greetingsx2 "ilearngocode/hello/stringlib"

	greetings "github.com/evenskylearngo/greetings/greetings"
)

func main() {
	x1 := greetingsx2.RandomFormat()
	fmt.Println(x1)
	fmt.Println(greetings.Hello("Go Developer"))
	fmt.Println(greetings.Goodbye("trang"))
	fmt.Println(greetings.Plus(1, 2))
	_, err := greetings.Hello("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
