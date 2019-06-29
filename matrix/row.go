package matrix

import (
	"errors"
	"fmt"
	"math/rand"
)

type Row struct {
	data []float64
}

func (row Row) Length() uint {
	return uint(len(row.data))
}

func (row *Row) Init(colNumber uint) error {
	if colNumber == 0 {
		return errors.New("colNumber == 0: Row can not be empty")
	}

	row.data = make([]float64, colNumber)
	return nil
}

func (row *Row) InitWithRandom(colNumber uint) error {
	if err := row.Init(colNumber); err != nil {
		return err
	}

	for i := range row.data {
		row.data[i] = rand.Float64()
	}

	return nil
}

func (row *Row) SetFromArray(array []float64) error {
	if len(array) == 0 {
		return errors.New("len(array) == 0: Row can not be empty")
	}

	row.data = array
	return nil
}

func (row Row) String() (str string) {
	for _, cell := range row.data {
		str += fmt.Sprintf("%f", cell) + "\t"
	}

	str += "\n"
	return
}

func (row Row) GetCell(cellN uint) (*float64, error) {
	if cellN >= row.Length() {
		return nil, errors.New(fmt.Sprintf("Row contains no cell with index %d", cellN))
	}

	return &row.data[cellN], nil
}
