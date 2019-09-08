package main

import "fmt"

//START OMIT
type Fruit struct {
    Name string
}
func (f Fruit) renameByValue(n string) {
    f.Name = n
}
func (f *Fruit) renameByReference(n string) {
    f.Name = n
}
func main() {
    f := Fruit { Name : "Apple" }
    fmt.Printf("Initial fruit: %+v \n", f)

    f.renameByValue("Banana")
    fmt.Printf("After rename by value: %+v \n", f)

    f.renameByReference("Banana")
    fmt.Printf("After rename by reference: %+v \n", f)
}
//END OMIT