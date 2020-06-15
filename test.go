package main

import "fmt"

func main() {
	z := "中文abc"
	for b, v := range z {
		fmt.Println(b, v)
	}

	a, b := 65, 2
	fmt.Println(a / b)
	fmt.Println(a)

	switch 65 {
	case 1:
	case 'A':
		fmt.Println("牛逼哥")
	}

	a, b = 1, 2

}
