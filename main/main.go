package main

import (
	"example.com/hello_prev"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(hello_prev.Hello())
	fmt.Println(hello_prev.Proverb())
}
