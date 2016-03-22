package main

import "fmt"
import "math"

func near_hundred(n int) bool {
    if (math.Abs(100 - float64(n)) <= 10) || (math.Abs(200-float64(n)) <= 10){
        return true
    } else {
        return false
    }
}

func main() {
    a := [...]int{93,90,89,110,111,121,0,5,191,189,190,200,210,211,290}
    fmt.Println(a)
    for _,e := range a {
        fmt.Println(near_hundred(e))
    }
}
