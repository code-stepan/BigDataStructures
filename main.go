package main

import (
	"fmt"
	"strconv"

	"github.com/code-stepan/BigDataStructures/bloomfilter"
	"github.com/code-stepan/BigDataStructures/bst"
	"github.com/code-stepan/BigDataStructures/btree"
	"github.com/code-stepan/BigDataStructures/hashtable"
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
			bst.Start() // +
		case 2:
			btree.StartBTree() // -
		case 3:
			// bloomfilter
			hasher := bloomfilter.NewMurMur3Hasher()
			bf, err := bloomfilter.NewBloomFilter(1000, 0.01, hasher)
			if err != nil {
				fmt.Printf("Ошибка BloomFilter: %v\n", err)
				break
			}
			bf.Add([]byte("meme"))
			bf.Add([]byte("mem"))
			bf.Add([]byte("memes"))

			fmt.Println(bf.Test([]byte("ddjdghf")))
		case 4:
			trie.Start()
		case 5:
			hashtable.StartHashTable()
		case 6:
		case 7:
		case 8:
		case 9:
		case 10:
		case 11:
		default:
			fmt.Println("Некорректная задача")
		}
	}
}
