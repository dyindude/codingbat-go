package main

func diff21(n int) int {
	if n > 21 {
		return (n - 21) * 2
	} else {
		return 21 - n
	}
}
