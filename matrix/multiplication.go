package matrix

import (
	"errors"
	"fmt"
)

func (row Row) Multiply(col Column) (*float64, error) {
	if row.Length() != col.Length() {
		return nil, errors.New(fmt.Sprintf("Row and Column has unequal size: %d vs %d",
			                               row.Length(), col.Length()))
	}

	var result float64

	for i := uint(0); i < row.Length(); i+= 1 {
		rValue, _ := row.GetCell(i)
		cValue, _ := col.GetCell(i)

		result += *rValue * *cValue
	}

	return &result, nil
}


func (matrix Matrix) Multiply(m2 Matrix) (rMatrix *Matrix, err error) {
	if matrix.ColsNumber() != m2.RowsNumber() {
		return nil, errors.New(fmt.Sprintf("matrix with %d columns was expected", matrix.RowsNumber()))
	}

	rMatrix = &Matrix{}

	arr := make([][]float64, 0)

	for i := uint(0); i < matrix.RowsNumber(); i += 1 {
		arrRow := make([]float64,0)

		for j := uint(0); j < m2.ColsNumber(); j += 1 {
			row, _ := matrix.GetRow(i)
			col, _ := m2.GetColumn(j)
			res, _ := row.Multiply(*col)

			arrRow = append(arrRow, *res)
		}

		arr = append(arr, arrRow)
	}

	_ = rMatrix.SetFromArrays(arr)
	
	return
}