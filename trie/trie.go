package trie

import "fmt"

type node[V any] struct {
	children map[rune]*node[V]
	value    V
	hasValue bool
}

type Trie[V any] struct {
	root    *node[V]
	allowed map[rune]struct{}
}

func NewTrie[V any](alphabet []rune) *Trie[V] {
	t := &Trie[V]{
		root: &node[V]{
			children: make(map[rune]*node[V]),
		},
	}
	if len(alphabet) > 0 {
		t.allowed = make(map[rune]struct{}, len(alphabet))
		for _, r := range alphabet {
			t.allowed[r] = struct{}{}
		}
	}
	return t
}

func (t *Trie[V]) Insert(key string, value V) error {
	curr := t.root
	for _, ch := range key {
		if t.allowed != nil {
			if _, ok := t.allowed[ch]; !ok {
				return fmt.Errorf("Символ %q отсутствует в настроенном алфавите", ch)
			}
		}
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = &node[V]{
				children: make(map[rune]*node[V]),
			}
		}
		curr = curr.children[ch]
	}
	curr.value = value
	curr.hasValue = true
	return nil
}

func (t *Trie[V]) Search(key string) (V, bool) {
	curr := t.root
	for _, ch := range key {
		child, ok := curr.children[ch]
		if !ok {
			var zero V
			return zero, false
		}
		curr = child
	}
	return curr.value, curr.hasValue
}

func (t *Trie[V]) Delete(key string) bool {
	runes := []rune(key)
	return t.deleteHelper(t.root, runes, 0)
}

func (t *Trie[V]) deleteHelper(curr *node[V], runes []rune, depth int) bool {
	if depth == len(runes) {
		if !curr.hasValue {
			return false
		}
		curr.hasValue = false
		return len(curr.children) == 0
	}

	ch := runes[depth]
	child, ok := curr.children[ch]
	if !ok {
		return false
	}

	if t.deleteHelper(child, runes, depth+1) {
		delete(curr.children, ch)
		return !curr.hasValue && len(curr.children) == 0
	}
	return false
}

func StartTrie() {
	alphabet := make([]rune, 0, 26)
	for c := 'a'; c <= 'z'; c++ {
		alphabet = append(alphabet, c)
	}
	trie := NewTrie[string](alphabet)

	_ = trie.Insert("apple", "🍎")
	_ = trie.Insert("app", "📱")
	_ = trie.Insert("bat", "🦇")

	if val, ok := trie.Search("apple"); ok {
		fmt.Printf("apple: %s\n", val)
	}
	if _, ok := trie.Search("apricot"); !ok {
		fmt.Println("apricot: не найдено")
	}

	fmt.Println("\nУдаляем ключ 'app'...")
	removed := trie.Delete("app")
	fmt.Printf("Удалено успешно: %v\n", removed)

	if _, ok := trie.Search("app"); !ok {
		fmt.Println("app: больше не найден (удалён)")
	}
	if val, ok := trie.Search("apple"); ok {
		fmt.Printf("apple: %s (остался в дереве)\n", val)
	}

	err := trie.Insert("123", "число")
	if err != nil {
		fmt.Printf("\nОшибка вставки: %v\n", err)
	}
}
