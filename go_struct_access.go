package main 

import "fmt"

type Person struct {
	Name string
	Age int 
}

func main() {
	p := Person { Name: "Adelina", Age: 18 }
	fmt.Printf("Name: %v", p.Name)
}
