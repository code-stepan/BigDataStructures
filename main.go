package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Print("Выбор задачи(от 1 до 8) >> ")
	fmt.Scan(&input)

	switch input {
	case 1:
		StartTask1()
	case 2:
		StartTask2()
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	default:
	}
}
