package main

import "testing"

func Testdiff21(t *testing.T) {
    t.Log("Testing output of diff21()")
    a := [12]int{19,10,21,22,25,30,0,1,2,-1,-2,50}
    b := [12]int{2,11,0,2,8,18,21,20,19,22,23,58}
    for i,_ := range a {
        if diff21(a[i]) != b[i] {
            t.Errorf("Expected output of %d, but it was %d instead.", b[i], a[i])
        }
    }

}
