package main

import "fmt"

func sayHello (text string) {
	fmt.Println(text)
}

func cal (x, y int) (res int) {
	res = x + y
	return
}

func calc2 (x, y int) (a int,b int) {
	a = x * y
	b = x + y
	return
}

func main() {
// x:=12
// y:=12

	// if age := x+y ; age >= 20 {
	// 	fmt.Println("adult",age)
	// } else if age >= 10 {
	// 	fmt.Println("girl")
	// } else {
	// 	fmt.Println("kids")
	// }

	// if x >= 5 && x< 10 {
	// 	fmt.Println("aaaaaaa")
	// }

	// i:=0
	// for i <= 10 {
	// 	fmt.Println("loop 1", i)
	// 	i++
	// }

	// x := 0
	// for  {
	// 	fmt.Println("loop 2", x)
	// 	if x == 10 {
	// 		break
	// 	}
	// 	x++
	// }

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("loop 3", i)
	// }

	func (greeting string) {
		fmt.Println(greeting)
	}("Evening")

	sayHello("Morning")
	res1, res2 := calc2(10, 2)
	fmt.Println(res1, res2)

	fmt.Println(cal(10, 5))
}



