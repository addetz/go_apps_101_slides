package main 

import "fmt" 

type Fruit struct {
	Name string
}
// START OMIT
func renamePassByReference(f *Fruit, newName string) {
	f.Name = newName
}

func renamePassByValue(f Fruit, newName string) {
	f.Name = newName
}

func main () { 
	f := Fruit { Name: "Apple" } 
	fmt.Printf("Original: %v \n", f)
	renamePassByValue(f, "Banana")
	fmt.Printf("Rename attempt 1: %v \n", f)
	renamePassByReference(&f, "Kiwi")
	fmt.Printf("Rename attempt 2: %v, \n", f)
}
//END OMIT
