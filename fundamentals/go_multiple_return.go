package main

import "fmt"
//START OMIT
func main() {
	greeting, err := sayHello()
	if err != nil {
		fmy.Println(err)
	}
	fmt.Println(greeting)
}

func sayHello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Cannot say hello to empty name.")
	}
	return fmt.Sprintf("Hello, %s!", name), nil
}
//END OMIT
