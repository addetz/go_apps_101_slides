package main

import "fmt"

//START OMIT
func main() {
    fruits := map[string]float64 {
        "Apple" : 0.75,
        "Banana" : 1.5,
    }
    fmt.Println(fruits)
    v, ok := fruits["Banana"]
    fmt.Printf("fruits[%s]=%f with ok: %t \n", "Banana",v, ok)
    u, ok := fruits["Kiwi"]
    fmt.Printf("fruits[%s]=%f with ok: %t \n", "Kiwi", u, ok)
}
//END OMIT
