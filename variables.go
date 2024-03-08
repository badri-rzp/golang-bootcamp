package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

type Vertex struct {
	x, y int
}

func main() {
	var a = 1
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(add(3, 4))
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum", sum)
	m := map[string]Vertex{
		"hsr":      {1, 2},
		"sarjapur": {3, 4},
	}
	fmt.Println(m)
}
