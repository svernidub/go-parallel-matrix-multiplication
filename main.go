package main

import (
	"fmt"
	"github.com/svernidub/go-parallel-matrix-multiplication/matrix"
)

func main() {
	m := matrix.Matrix{}

	/*
	err := m.SetFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8 ,9},
	})
	 */

	err := m.InitWithRandom(10000, 10000)

	if err != nil {
		fmt.Println(err)
	} else {
		m2, _ := m.GoroutinesMultiply(m)
		fmt.Println(m2)
	}
}