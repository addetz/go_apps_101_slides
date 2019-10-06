package main

import "fmt"

//START OMIT
type Fruit struct {
	Name string
	Price float64
}

func main() {
	var zero Fruit
	fmt.Printf("%v \n", zero)
	
	p := Fruit { Name: "Banana", Price: 0.75 }
	fmt.Printf("%v", p)
}
//END OMIT
