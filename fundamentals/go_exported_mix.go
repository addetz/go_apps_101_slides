package main

//START OMIT
type unexported struct { // only visible inside own package
    Exported string // only visible inside own package
}
type Exported struct { // visible outside own package
    Exported string // visible outside own package
    unexported string // only visible inside own package
}
//END OMIT
