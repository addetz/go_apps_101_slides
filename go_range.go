package main 

import "fmt" 

func main() {
	fruits := []string{"Apple", "Banana", "Kiwi", "Orange"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("%d - %s ", i, fruits[i])
	}
	fmt.Println()
	for i, fruit := range fruits {
		fmt.Printf("%d - %s ", i, fruit)
	}
}
