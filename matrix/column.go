package matrix

import (
	"errors"
	"fmt"
)

type Column struct {
	data []float64
}

func (col Column) Length() uint {
	return uint(len(col.data))
}

func (col Column) String() (str string) {
	for _, cell := range col.data {
		str += fmt.Sprintf("%f", cell) + "\n"
	}

	return
}

func (col *Column) Init(cellNumber uint) error {
	if cellNumber == 0 {
		return errors.New("cellNumber == 0: Column can not be empty")
	}

	col.data = make([]float64, cellNumber)
	return nil
}

func (col *Column) SetFromArray(array []float64) error {
	if len(array) == 0 {
		return errors.New("len(array) == 0: Column can not be empty")
	}

	col.data = array
	return nil
}

func (col *Column) SetCell(cellN uint, value float64) error {
	if cellN >= col.Length() {
		return errors.New(fmt.Sprintf("Column contains no cell with index %d", cellN))
	}

	col.data[cellN] = value
	return nil
}

func (col Column) GetCell(cellN uint) (*float64, error) {
	if cellN >= col.Length() {
		return nil, errors.New(fmt.Sprintf("Column contains no cell with index %d", cellN))
	}

	return &col.data[cellN], nil
}
