package main

import "fmt"

type Matrix struct {
	row, col int
	matrix   [][]int
}

func CreateMatrixStruct(row int, col int) Matrix {
	m := Matrix{}
	m.row = row
	m.col = col
	m.matrix = make([][]int, row)
	for i, _ := range m.matrix {
		temp := make([]int, col)
		m.matrix[i] = temp
	}
	return m

}

func main() {
	fmt.Println(CreateMatrixStruct(3, 3))
}
