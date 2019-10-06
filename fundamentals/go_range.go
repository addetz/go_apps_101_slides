package main

import "fmt"
//START OMIT
func main() {
	fruits := []string{"Apple", "Banana", "Kiwi", "Orange"}
	for i, fruit := range fruits {
		fmt.Printf("%d - %s ", i, fruit)
	}
	fmt.Println()
	for _, fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
}
//END OMIT
