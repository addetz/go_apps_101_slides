package main

import "fmt"

//START OMIT
type Person struct {
    Name string
}

func main() {
    p := Person{Name : "Adelina"}

    fmt.Printf("%v \n", p)
    fmt.Printf("%+v \n", p)
}
//END OMIT