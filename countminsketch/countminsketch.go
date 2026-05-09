package countminsketch

import (

)

type HashFunc func(data []byte) uint64

type CountMinSketch struct {
	depth int
	width int
	table [][]uint64
	hashFunc []func(data []byte) int
}