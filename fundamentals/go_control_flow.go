package main

import "fmt"
//START OMIT
func main() {
	if s := sum(3,5); s != 0 {
		fmt.Sprintf("Sum is %d", s)
	}
	fmt.Println(s)
}

func sum(a int, b int) int {
	return a + b
}
//END OMIT
