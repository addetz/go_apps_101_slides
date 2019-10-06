package main

import "fmt"
//START OMIT
import "encoding/json"
type Salad struct{
    Type string `json:"type"`
    Ingredients []string `json:"ingredients"`
}
func main() {
    s := `{"type": "Egg", "ingredients": ["egg", "mayo", "mustard"]}`
    resp := Salad{}
    json.Unmarshal([]byte(s), &resp)
    fmt.Printf("JSON decoded: \n%v \n", resp)
}
//END OMIT