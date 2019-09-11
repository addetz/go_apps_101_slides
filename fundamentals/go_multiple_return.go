package main

import (
	"errors"
	"fmt"
)
//START OMIT
func main() {
	greeting, err := sayHello("")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(greeting)
}

func sayHello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("cannot say hello to empty name")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}
//END OMIT
