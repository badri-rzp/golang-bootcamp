package main

import (
	"fmt"

	"github.com/badri-rzp/golang-bootcamp/day01/matrix"
)

func main() {
	matrix1 := matrix.CreateMatrix(3, 3)
	matrix2 := matrix.CreateMatrix(3, 3)

	matrix1.PopulateMatrixWithRandomValues(100)
	matrix2.PopulateMatrixWithRandomValues(200)

	matrix1.ShowMatrix()
	fmt.Println()
	matrix2.ShowMatrix()

	if matrix3, error := matrix1.AddMatrices(&matrix2); error != nil {
		fmt.Println(error)
	} else {
		fmt.Println()
		matrix3.ShowMatrix()
	}
}
