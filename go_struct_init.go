package main 

import "fmt"

type Person struct {
	Name string
	Age int 
}

func main() {
	var zero Person 
	fmt.Printf("%v \n", zero)
	
	p := Person { Name: "Adelina", Age: 18 }
	fmt.Printf("%v", p)
}
