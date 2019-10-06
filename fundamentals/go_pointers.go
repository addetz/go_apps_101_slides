package main

import "fmt"

// START OMIT
func main () {
    i := 42
    ipointer := &i
    fmt.Println("i: ", i)
    fmt.Println("ipointer: ", ipointer)
    fmt.Println("Dereference ipointer: ", *ipointer)
}
//END OMIT

