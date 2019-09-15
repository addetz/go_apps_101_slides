package main

import "fmt"

//START OMIT
func main() {
    books := map[string]string {
        "The bell jar" : "Sylvia Plath",
        "1984" : "George Orwell",
    }
    fmt.Println(books)
    v, ok := books["1984"]
    fmt.Printf("books[%s]=%s with ok: %t \n", "1984",v, ok)
    u, ok := books["Mrs. Dalloway"]
    fmt.Printf("books[%s]= %s with ok: %t \n", "Mrs. Dalloway", u, ok)
}
//END OMIT
