package main

import "fmt"

//START OMIT
type SayHi interface {
    hello() string
}
type Person struct { Name string }
func (p Person) hello() string {
    return fmt.Sprintf("Hi, I'm %s!", p.Name)
}
func hiBack(h SayHi) {
    fmt.Println(h.hello())
    fmt.Println("Hi Back!")
}
func main() {
    p := Person{ Name: "Adelina" }
    hiBack(p)
}
//END OMIT
