package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct{
	rows []Row
}

func (matrix *Matrix) Init(rows uint, cols uint) error {
	matrix.rows = make([]Row, rows)

	for i := range matrix.rows {
		row := Row{}

		if err := row.Init(cols); err != nil {
			return err
		}

		matrix.rows[i] = row
	}

	return nil
}

func (matrix *Matrix) InitWithRandom(rows uint, cols uint) error {
	if err := matrix.Init(rows, cols); err != nil {
		return err
	}

	for i := range matrix.rows {
		r := Row{}
		_ = r.InitWithRandom(cols)
		matrix.rows[i] = r
	}

	return nil
}

func (matrix *Matrix) SetFromArrays(array [][]float64) (err error) {
	if err = validateSourceArray(array); err != nil { return }

	matrix.rows = make([]Row, len(array))

	for i := range matrix.rows {
		r := Row{}
		if err = r.SetFromArray(array[i]); err != nil { return }

		matrix.rows[i] = r
	}
	return
}

func (matrix Matrix) RowsNumber() uint {
	return uint(len(matrix.rows))
}

func (matrix Matrix) ColsNumber() uint {
	return matrix.rows[0].Length()
}

func (matrix Matrix) String() (str string) {
	str = "Matrix\n"

	for _, row := range matrix.rows {
		str += row.String()
	}

	return
}

func (matrix Matrix) GetRow(rowN uint) (*Row, error) {
	if rowN >= matrix.RowsNumber() {
		return nil, errors.New(fmt.Sprintf("Matrix contains no Row with index %d", rowN))
	}
	return &matrix.rows[rowN], nil
}

func (matrix Matrix) GetColumn(colN uint) (*Column, error) {
	if colN >= matrix.ColsNumber() {
		return nil, errors.New(fmt.Sprintf("Matrix contains no Column with index %d", colN))
	}

	column := Column{}

	if err := column.Init(matrix.RowsNumber()); err != nil {
		return nil, err
	}

	for i, row := range matrix.rows {
		if cell, err := row.GetCell(colN); err == nil {
			if err := column.SetCell(uint(i), *cell); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &column, nil
}

func validateSourceArray(array [][]float64) error {
	if len(array) == 0 {
		return errors.New("zero rows array could not be transformed to Matrix")
	}

	cols := len(array[0])

	for _, row := range array {
		if cols != len(row) {
			return errors.New("each Matrix row length should be equal")
		}
	}

	return nil
}
