package main

import "fmt"
//START OMIT
func main() {
	sum1 := sum (4, 10)
	if sum1 != 0 {
		fmt.Println(sum1)
	}
}
func sum(a int, b int) int {
	return a + b
}
//END OMIT
