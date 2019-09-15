package main

import "fmt"

//START OMIT
type Person struct {
	Name string
	Title string
}

func main() {
	p := Person { Name: "Adelina", Title: "Miss" }
	fmt.Printf("Name: %v", p.Name)
}
//END OMIT