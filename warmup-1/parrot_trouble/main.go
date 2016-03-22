package main

import "fmt"

func parrot_trouble(talking bool, hour int) bool {
    if (talking && (hour <  7 || hour > 20)) {
        return true
    } else {
        return false
    }
}
func main() {
    fmt.Println(parrot_trouble(true,6))
    fmt.Println(parrot_trouble(true,7))
    fmt.Println(parrot_trouble(false,6))
    fmt.Println(parrot_trouble(true,21))
    fmt.Println(parrot_trouble(false,21))
    fmt.Println(parrot_trouble(false,20))
    fmt.Println(parrot_trouble(true,23))
    fmt.Println(parrot_trouble(false,23))
    fmt.Println(parrot_trouble(true,20))
    fmt.Println(parrot_trouble(false,12))


}
