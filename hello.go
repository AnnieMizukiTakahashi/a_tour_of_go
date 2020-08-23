package main

import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
