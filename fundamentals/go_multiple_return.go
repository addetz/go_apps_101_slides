package main

import "fmt"

// START OMIT
func main() {
    sum, prod, diff := calculate(2,3)
    fmt.Printf("%d, %d, %d \n", sum, prod, diff)

    _, prod2, _ := calculate(4,5)
    fmt.Println(prod2)
}
func calculate(a,b int) (int, int, int){
    return a+b, a*b, a-b
}
//END OMIT