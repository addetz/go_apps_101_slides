package main

import "fmt"
//START OMIT
type Person struct {
	Name string
	Title string
}

func main() {
	var zero Person 
	fmt.Printf("%v \n", zero)
	
	p := Person { Name: "Adelina", Title: "Miss" }
	fmt.Printf("%v", p)
}
//END OMIT
