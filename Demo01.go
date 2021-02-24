package main

import "fmt"

func hello(str string) string {

	return "hello " + str
}

func main() {
	fmt.Print(hello("golang"))
}
