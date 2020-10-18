package main

import (
	"fmt"
	// "math"
	// "time"
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

// For文の書き方
// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }

// Forのイテレーションの初期値と後処理文は書かなくてもいい
// func main() {
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// ; セミコロンは省略できる。goではwhileをforで書く
// func main() {
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 	}
// 	fmt.Println(sum)
// }

// if文
// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

// func main() {
// 	fmt.Println(sqrt(2), sqrt(-5))
// }

// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}
// 	return lim
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }

//使ってるosを取得できる(!)
// func main() {
// 	fmt.Print("Go run on ")
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X")
// 	case "linux":
// 		fmt.Println("Linux")
// 	default:
// 		fmt.Printf("%s.\n", os)
// 	}
// }

// case に式を書く時
// func main() {
// 	fmt.Println("土曜日はいつ？")
// 	today := time.Now().Weekday()
// 	switch time.Saturday {
// 	case today + 0:
// 		fmt.Println("今日")
// 	case today + 1:
// 		fmt.Println("昨日")
// 	case today + 2:
// 		fmt.Println("in two days")
// 	default:
// 		fmt.Println("too far away")
// 	}
// }

// swichに何もconditionを入れないと、switch trueと同じ意味になる。
// この構文だと、if then else文を短く書けるメリットがある。
// func main() {
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("Good morning")
// 	case t.Hour() < 17:
// 		fmt.Println("Good afternoon")
// 	default:
// 		fmt.Println("Good evening.")
// 	}
// }

// func main() {
// 	defer fmt.Println("world")
// 	fmt.Println("hello")
// }

// deferへ渡した関数が複数の処理(forみたいに)するとき、lastから実行される
// func main() {
// 	fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}
// 	fmt.Println("done")
// }

// methods

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }

// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(Abs(v))
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// func main() {
// 	i, j := 42, 2701

// 	p := &i         // point to i
// 	fmt.Println(*p) // read i through the pointer
// 	*p = 21         // set i through the pointer
// 	fmt.Println(i)  // see the new value of i

// 	p = &j         // point to j
// 	*p = *p / 37   // divide j through the pointer
// 	fmt.Println(j) // see the new value of j
// }

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
