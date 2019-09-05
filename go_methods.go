package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) sayHello() {
	fmt.Printf("Hi! I'm %v!", p.Name)
}

func main() {
	p := Person { Name : "Skip" }
	p.sayHello()
}
