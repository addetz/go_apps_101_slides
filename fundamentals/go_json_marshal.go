package main

import "fmt"
//START OMIT
import "encoding/json"
type Salad struct{
    Type string `json:"type"`
    Ingredients []string `json:"ingredients"`
}
func main() {
    s := Salad{Type : "Fruit", Ingredients:[]string{"apple", "banana", "kiwi"}}
    bytes, _ := json.Marshal(s)
    fmt.Printf("JSON enconded: \n%s \n", string(bytes))
}
//END OMIT