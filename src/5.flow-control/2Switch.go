package main

import (
    "fmt"
    "runtime"
    "time"
)

func mySwitch() {
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.\n", os)
    }
}

func findDay() {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()

    switch today {
    case today + 0:
        fmt.Println("Today.")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days.")
    default:
        fmt.Println("Too far away.")
    }
}

func switchNoCondition() {
    t := time.Now()

    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
func main() {
    mySwitch()

    findDay()

    switchNoCondition()
}
