package main

import (
	"fmt"
	"strconv"

	"github.com/code-stepan/BigDataStructures/bloomfilter"
	"github.com/code-stepan/BigDataStructures/bst"
	"github.com/code-stepan/BigDataStructures/countminsketch"
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

			if val, ok := intTree.Get(40); ok {
				fmt.Println("Найдено: ", val)
			}

			deleted := intTree.Delete(50)
			fmt.Println("Удалние: ", deleted)

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
			t := trie.New[string]("abcdefghijklmnopqrstuvwxyz")
			t.Insert("hello", "привет")
			t.Insert("world", "мир")
			t.Insert("help", "помощь")
			t.Insert("heap", "куча")

			if val, ok := t.Get("hello"); ok {
				fmt.Println("Найдено hello:", val)
			}
			if val, ok := t.Get("world"); ok {
				fmt.Println("Найдено world:", val)
			}

			fmt.Println("StartsWith he:", t.StartsWith("he"))
			fmt.Println("StartsWith xyz:", t.StartsWith("xyz"))

			if t.Delete("help") {
				fmt.Println("help удален")
			}
			if val, ok := t.Get("help"); ok {
				fmt.Println("Get help после удаления:", val)
			} else {
				fmt.Println("help не найден")
			}
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
			// count min sketch
			cms := countminsketch.New(0.01, 0.01)
			cms.Add([]byte("apple"))
			cms.Add([]byte("banana"))
			cms.Add([]byte("apple"))
			cms.Add([]byte("apple"))
			cms.Add([]byte("banana"))
			cms.Add([]byte("cherry"))

			fmt.Printf("apple count: %d\n", cms.Count([]byte("apple")))
			fmt.Printf("banana count: %d\n", cms.Count([]byte("banana")))
			fmt.Printf("cherry count: %d\n", cms.Count([]byte("cherry")))
			fmt.Printf("unknown count: %d\n", cms.Count([]byte("unknown")))
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
