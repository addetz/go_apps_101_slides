package main

import "fmt"

//START OMIT
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
//END OMIT
