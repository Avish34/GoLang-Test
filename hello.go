package main

import "fmt"

func recur(x int) int {

	if x == 0 {
		return 0
	}

	return x + recur(x-1)
}

func main() {

	var name string = "avish"
	fmt.Println(name)

	for x := 0; x < 5; x++ {

		fmt.Println(x)
	}
	fmt.Println(recur(5))
}
