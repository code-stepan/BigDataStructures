package main

import (
	"fmt"
	"strconv"

	"github.com/code-stepan/BigDataStructures/bloomfilter"
	"github.com/code-stepan/BigDataStructures/bst"
	"github.com/code-stepan/BigDataStructures/btree"
	"github.com/code-stepan/BigDataStructures/trie"
)

func main() {
	var inputString string

	for {
		fmt.Print("\nВыбор задачи(от 1 до 8)[используйте 0 для выхода] \n >> ")
		fmt.Scan(&inputString)

		input, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("Некорректный ввод")
			continue
		}

		if input == 0 {
			break
		}

		switch input {
		case 1:
			bst.StartBST()
		case 2:
			btree.StartBTree()
		case 3:
			bloomfilter.StartBloomFilter()
		case 4:
			trie.StartTrie()
		case 5:
		case 6:
		case 7:
		case 8:
		default:
			fmt.Println("Некорректная задача")
		}
	}
}
