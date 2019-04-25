package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	rand.Seed(10)
	fmt.Println("My favorite number is", rand.Intn(10))


	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)
}
