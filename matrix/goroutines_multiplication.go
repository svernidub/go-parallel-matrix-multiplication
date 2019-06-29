package matrix

import (
	"errors"
	"fmt"
)

type goRoutineResult struct {
	row uint
	cell uint
	value float64
}

func (matrix Matrix) GoroutinesMultiply(m2 Matrix) (rMatrix *Matrix, err error) {
	if matrix.ColsNumber() != m2.RowsNumber() {
		return nil, errors.New(fmt.Sprintf("matrix with %d columns was expected", matrix.RowsNumber()))
	}

	rMatrix = &Matrix{}


	arr := make([][]float64, matrix.RowsNumber())
	for i := range arr {
		arr[i] = make([]float64, m2.ColsNumber())
	}

	input := make(chan goRoutineResult)

	go planCalculations(matrix, m2, input)

	tasks := matrix.RowsNumber() * m2.ColsNumber()

	for i := uint(0); i < tasks; i += 1 {
		select {
		case result := <-input:
			arr[result.row][result.cell] = result.value
		}
	}

	_ = rMatrix.SetFromArrays(arr)

	return
}

func planCalculations(m1, m2 Matrix, output chan<- goRoutineResult) {
	for i := uint(0); i < m1.RowsNumber(); i += 1 {
		for j := uint(0); j < m2.ColsNumber(); j += 1 {
			row, _ := m1.GetRow(i)
			col, _ := m2.GetColumn(j)

			go calculate(i, j, *row, *col, output)
		}
	}
}

func calculate(rowN uint, cellN uint, row Row, col Column, output chan<- goRoutineResult) {
	result, _ := row.Multiply(col)

	output <- goRoutineResult{
		row: rowN,
		cell: cellN,
		value: *result,
	}
}