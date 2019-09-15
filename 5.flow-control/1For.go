package main

import (
    "fmt"
    "math"
)

func sum() {
    sum := 0

    for i := 0; i < 10000; i++ {
        sum += 1
    }
    fmt.Println("sum:", sum)
}

func sumContinue() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println("sumcontinue:", sum)
}

func sumWhileFor() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println("sumWhile:", sum)
}
func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(i int) {
    if v := math.Pow(float64(i), 2); v < 15 {
        fmt.Println(" nho hon cai dieu kien roi")
    } else {
        fmt.Println("OK man")
    }

}
func main() {
    sum()
    sumContinue()
    sumWhileFor()
    fmt.Println(sqrt(12), sqrt(4))
    pow(12)
}
