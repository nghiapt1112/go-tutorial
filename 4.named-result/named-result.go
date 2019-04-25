package main

import "fmt"

/**
A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
 */

func sum(sum int) (x, y int) {
	x = sum * 2
	y = sum -1
	return
}
func main() {
	fmt.Print(sum(12))
}
