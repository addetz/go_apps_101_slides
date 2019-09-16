package main

import "fmt"

//START OMIT
type Greeter interface {
    greet() string
}
type Person struct { Name string }
type Friend struct { Name string }
func (p Person) greet() string {
    return fmt.Sprintf("Hi, I'm %s!", p.Name)
}
func (f Friend) greet() string {
    return fmt.Sprintf("Hi, I'm %s!", f.Name)
}
func hiBack(g Greeter) {
    fmt.Println(g.greet())
    fmt.Println("Hi Back!")
}
func main() {
    p := Person{ Name: "Adelina" }
    f := Friend{ Name: "Skip" }
    hiBack(p)
    hiBack(f)
}
//END OMIT
