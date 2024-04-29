package main

import "fmt"

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

	i:=0
	for i <= 10 {
		fmt.Println("loop 1", i)
		i++
	}

	x := 0
	for  {
		fmt.Println("loop 2", x)
		if x == 10 {
			break
		}
		x++
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("loop 3", i)
	}

}



