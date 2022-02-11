package main

import "fmt"

var name string

func Hello(food string) {
	fmt.Println(name + " like " + food)
}

func main() {
	name = "Tony.xu"
	Hello("apple")
}
