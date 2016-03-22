package main

import "fmt"

func sleep_in(weekday bool, vacation bool) bool {
    if weekday != true || vacation {
        return true
    } else {
        return false
    }
}

func main() {
    fmt.Println(sleep_in(false,false))
    fmt.Println(sleep_in(true,false))
    fmt.Println(sleep_in(false,true))
    fmt.Println(sleep_in(true,true))
}
