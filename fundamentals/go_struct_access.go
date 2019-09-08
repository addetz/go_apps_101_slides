package main

import "fmt"

//START OMIT
type Person struct {
	Name string
	Age int 
}

func main() {
	p := Person { Name: "Adelina", Age: 18 }
	fmt.Printf("Name: %v", p.Name)
}
//END OMIT