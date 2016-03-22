package main

import "fmt"

func diff21(n int) int {
    if n > 21 {
        return (n-21) * 2
    } else {
        return 21-n
    }
}

func main() {
    a := [12]int{19,10,21,22,25,30,0,1,2,-1,-2,50}
    fmt.Println(a)
    for _,e := range a {
        fmt.Println(diff21(e))
    }
}
