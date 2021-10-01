// package matrix takes a string and creates a matrix with methods
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// type Matrix stores a matix of integers
type Matrix [][]int

// New creates a matrix from an input string and returns an error depending on input
func New(input string) (*Matrix, error) {
	new := Matrix{}
	rows := strings.Split(input, "\n")
	for i, row := range rows {
		if strings.TrimSpace(row) == "" {
			return nil, errors.New("blank row in input")
		}
		row_sl := strings.Split(strings.TrimSpace(row), " ")
		if i > 0 && len(strings.Split(rows[i-1], " ")) != len(row_sl) {
			return nil, errors.New("uneven rows")
		}
		col_sl := []int{}
		for _, col := range row_sl {
			val, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			col_sl = append(col_sl, val)
		}
		new = append(new, col_sl)

	}
	return &new, nil
}

// Set method sets one cell of a Matrix
func (m Matrix) Set(row, column, value int) bool {
	if row > len(m)-1 || row < 0 || column > len(m[0])-1 || column < 0 || len(m) == 0 {
		return false
	}
	m[row][column] = value
	return true
}

// Rows method returns the rows of a Matrix
func (m Matrix) Rows() [][]int {
	matrix := make([][]int, len(m))
	for i, row := range m {
		matrix[i] = make([]int, len(row))
		for j, col := range row {
			matrix[i][j] = col
		}
	}
	return matrix
}

// Cols method returns the rows of a Matrix
func (m Matrix) Cols() [][]int {
	transpose := make([][]int, len(m[0]))
	for i, _ := range m[0] {
		transpose[i] = make([]int, len(m))
		for j, _ := range m {
			transpose[i][j] = m[j][i]
		}
	}
	return transpose
}
