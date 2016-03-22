package main

import "fmt"

func missing_char(s string, n int) string {
    return s[:n] + s[n+1:]
}

func main() {
    //there is no way to determine which order go processes a map,
    //thus we should reorganize this with a 2d slice/array or similar
    m := map[string]int{
        "kittenk": 1,
        "kittenz": 0,
        "kittens": 4,
        "Hix"    : 0,
        "His"    : 1,
        "codes"  : 0,
        "codex"  : 1,
        "codec"  : 2,
        "coded"  : 3,
        "chocolate": 8,
    }
      
    fmt.Println(m)
    for k,v := range m {
        fmt.Println(missing_char(k,v))
    }
}

