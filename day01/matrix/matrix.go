package matrix

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
)

type Matrix struct {
	row    int     `json:"row"`
	col    int     `json:"col"`
	matrix [][]int `json:"matrix"`
}

func CreateMatrix(row, col int) Matrix {
	m := Matrix{}
	m.row = row
	m.col = col
	m.matrix = make([][]int, row)
	for i := range row {
		m.matrix[i] = make([]int, col)
	}
	return m
}

// pointer receivers
func (m *Matrix) GetRows() int {
	return m.row
}

func (m *Matrix) GetCols() int {
	return m.col
}

func (m *Matrix) Set(row, col, val int) {
	m.matrix[row][col] = val
}

func (m *Matrix) Get(row, col int) int {
	return m.matrix[row][col]
}

func (m1 *Matrix) AddMatrices(m2 *Matrix) (Matrix, error) {
	if m1.GetRows() != m2.GetRows() || m1.GetCols() != m2.GetCols() {
		return Matrix{}, errors.New("Cannot Add. Different Dimensions")
	}
	m3 := CreateMatrix(m1.GetRows(), m1.GetCols())
	for row := range m1.GetRows() {
		for col := range m1.GetCols() {
			m3.Set(row, col, m1.Get(row, col)+m2.Get(row, col))
		}
	}
	return m3, nil
}

func (m *Matrix) ShowMatrix() {
	for row := range m.GetRows() {
		for col := range m.GetCols() {
			fmt.Print(m.matrix[row][col], " ")
		}
		fmt.Println()
	}
}

func (m *Matrix) PopulateMatrixWithRandomValues(i int) {
	for row := range m.GetRows() {
		for col := range m.GetCols() {
			m.Set(row, col, rand.Intn(i))
		}
	}
}

func (m *Matrix) ToJson() (string, error) {
	json_matrix_bytes, error := json.Marshal(m.matrix)
	if error != nil {
		return "", error
	}
	return string(json_matrix_bytes), nil
}

func main() {
	fmt.Println(CreateMatrix(3, 3))
}
