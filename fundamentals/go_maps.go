package main

import "fmt"

//START OMIT
func main() {
    fruits := map[string]float64 {
        "Apple" : 0.75,
        "Banana" : 1.5,
    }
    fmt.Println(fruits)
    fmt.Printf("fruits[%s]= %f \n", "Apple", fruits["Apple"])
    fmt.Printf("fruits[%s]= %f \n", "Kiwi", fruits["Kiwi"])
}
//END OMIT
