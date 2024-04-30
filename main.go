package main

import "fmt"

func Hello(greetString string) string {
	return "Hello, " + greetString
}

func main() {
	fmt.Println(Hello("harshit"))
}
