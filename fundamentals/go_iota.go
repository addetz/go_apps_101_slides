package main

import "fmt"

//START OMIT
type DayOfWeek int
const (
	MON DayOfWeek = iota
	TUE
	WED
	THUR
	FRI
)

func main() {
	fmt.Println(MON, TUE, WED, THUR, FRI)

}
//END OMIT

