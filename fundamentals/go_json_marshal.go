package main

//START OMIT
import (
    "encoding/json"
    "fmt"
)
type Salad struct{
    Type string `json:"type"`
    Ingredients []string `json:"ingredients"`
}
func main() {
    s := Salad{
        Type : "Fruit",
        Ingredients:[]string{"banana", "apple", "orange"},
    }
    bytes, _ := json.Marshal(s)
    fmt.Printf("JSON enconded: \n%s \n", string(bytes))
}
//END OMIT