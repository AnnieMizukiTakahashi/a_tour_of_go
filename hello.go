package main

import (
	"fmt"
)

// func add(x, y int) int {
// 	return x + y
// }
// func main() {
// 	fmt.Println(add(42, 13))
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }
// func splint(sum int) (x, y int) {
// 	x = sum * 4 / 8
// 	y = sum - x
// 	return
// }

// func main() {
// 	fmt.Println(splint(6))
// }

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

// 初期化子があれば、関数内で型宣言がいらない
// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

// :=これを使うと関数の中だけで暗黙的に型宣言ができる
// func main() {
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "you"

// 	fmt.Println(i, j, k, c, python, java)
// }

// 基本の型　整数形がたくさんある
// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
// 	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
// 	fmt.Printf("Type: %T Value: %v\n", z, z)
// }

// 初期値がないと勝手に0が入る
// func main() {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }

// 型は変えられる 型(変数)
// func main() {
// 	var i int = 42
// 	var f float64 = float64(i)
// 	var u uint = uint(f)
// 	fmt.Println(i, f, u)
// 	fmt.Println(reflect.TypeOf(i)) //int
// 	fmt.Println(reflect.TypeOf(f)) //float64
// 	fmt.Println(reflect.TypeOf(u)) //uint
// }

// For
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
