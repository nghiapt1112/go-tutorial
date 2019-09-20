package model

import "fmt"

type Person struct {
    Name string
    Age  int
}


/**
*like toString() method in java <br>
*A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
 */
func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}