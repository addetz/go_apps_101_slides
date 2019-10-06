package main

import "fmt"

//START OMIT
type Greeter interface {
    greet() string
}
type Person struct { Name string }
func (p Person) greet() string { return fmt.Sprintf("Hi, I'm %s!", p.Name) }
func (p Person) sayBye() string { return fmt.Sprintf("Byeeeeee!") }
func hiBack(g Greeter) {
    fmt.Println(g.greet())
    fmt.Println(g.sayBye())
}
func main() {
    p := Person{ Name: "Adelina" }
    hiBack(p)
}
//END OMIT

