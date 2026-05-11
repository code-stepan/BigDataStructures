package main

import (
	"fmt"
	"strconv"

	"github.com/code-stepan/BigDataStructures/bloomfilter"
	"github.com/code-stepan/BigDataStructures/bst"
	"github.com/code-stepan/BigDataStructures/hashtable"
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
			// bst
			intTree := bst.New[int, string](func(a, b int) int {
				return a - b
			})

			intTree.Insert(50, "fifty")
			intTree.Insert(30, "thirty")
			intTree.Insert(70, "seventy")
			intTree.Insert(20, "twenty")
			intTree.Insert(40, "forty")
			intTree.Insert(60, "sixty")
			intTree.Insert(80, "eighty")

			// Поиск
			if val, ok := intTree.Get(40); ok {
				fmt.Println("Найдено: ", val)
			}

			// Удаление
			deleted := intTree.Delete(50)
			fmt.Println("Удалние: ", deleted)

			// Обходы
			fmt.Println("PreOrder: ")
			intTree.PreOrder(func(k int, v string) {
				fmt.Printf("%d(%s) ", k, v)
			})
			fmt.Println()

			fmt.Println("InOrder: ")
			intTree.InOrder(func(k int, v string) {
				fmt.Printf("%d(%s) ", k, v)
			})
			fmt.Println()

			fmt.Println("PostOrder: ")
			intTree.PostOrder(func(k int, v string) {
				fmt.Printf("%d(%s) ", k, v)
			})
			fmt.Println()
		case 2:
			// btree
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
			// trie
		case 5:
			// hashtable
			h, err := hashtable.New[string, int](8, func(s string) uint64 {
				var hash uint64
				for i := 0; i < len(s); i++ {
					hash = hash*31 + uint64(s[i])
				}
				return hash
			})
			if err != nil {
				panic(err)
			}

			h.Set("go", 1)
			h.Set("rust", 2)
			h.Set("zig", 3)

			fmt.Printf("Размер: %d\n", h.Len())

			if val, ok := h.Get("rust"); ok {
				fmt.Printf("rust гайден: %d\n", val)
			}

			if h.Delete("go") {
				fmt.Println("go удален")
			}
			fmt.Printf("Размер: %d\n", h.Len())
		case 6:
			// Count min sketch
		case 7:
			// segment tree
		case 8:
			// Fenwick Tree
		case 9:
		case 10:
		case 11:
		default:
			fmt.Println("Некорректная задача")
		}
	}
}
