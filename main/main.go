package main

import (
	"example.com/hello_prev"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(hello_prev.Hello())
	fmt.Println(hello_prev.Proverb())

	m := make(map[string][]int)
	m["hoge"] = []int{1, 2, 4}
	m["hoge2"] = []int{10, 12, 14}
	m["fuga"] = []int{1, 2, 3}

	for key, value := range m {
		fmt.Printf("%v : %v;\n", key, value)
	}
}
