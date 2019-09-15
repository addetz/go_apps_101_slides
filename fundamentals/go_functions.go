package main

import "fmt"

//START OMIT
func main() {
	fmt.Println(helloWorld())
	h := func() string {
		return "Hello, again world!"
	}
	fmt.Println(h())
}
func helloWorld() string {
	return "Hello, world!"
}
//END OMIT
