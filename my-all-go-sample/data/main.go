package main

import "fmt"

type data struct {
	foo string
	baz string
}

func main() {
	fmt.Println("start!")

	d := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	fmt.Printf("%v \n", d)
	fmt.Println(d["foo"])

	d2 := data{
		foo: "hoge",
		baz: "fuga",
	}
	fmt.Printf("%#v \n", d2)
	fmt.Printf("%+v \n", d2)
	fmt.Println(d2.foo)
}
