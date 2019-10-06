package main

import "fmt"
//START OMIT
func main() {
    fruits := []string{"Apple", "Banana", "Kiwi", "Orange"}
    var i int
    for i = 0; i < len(fruits); i++ {
        fmt.Printf("%d - %s ", i, fruits[i])
    }
    fmt.Println()
    for i := 0; i < len(fruits); i++ {
        fmt.Printf("%d - %s ", i, fruits[i])
    }
}
//END OMIT