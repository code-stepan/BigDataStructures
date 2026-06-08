package heaps

// Узел биномиального дерева.
// Каждое биномиальное дерево B_k имеет 2^k узлов и degree = k.
// child — первый (самый левый) ребёнок, sibling — следующий брат (справа).
type BinomialNode[K any, V any] struct {
	key     K
	val     V
	degree  int
	child   *BinomialNode[K, V]
	sibling *BinomialNode[K, V]
}

// Биномиальная куча — набор биномиальных деревьев (лес).
// Деревья хранятся в map, где ключ — степень дерева.
// Свойство: в куче не более одного дерева каждой степени.
// Свойство max-кучи: ключ родителя >= ключей детей.
type BinomialHeap[K any, V any] struct {
	roots   map[int]*BinomialNode[K, V] // деревья, индексированные по степени
	compare func(a, b K) int
	size    int
}

func NewBinomialHeap[K any, V any](cmp func(a, b K) int) *BinomialHeap[K, V] {
	return &BinomialHeap[K, V]{
		roots:   make(map[int]*BinomialNode[K, V]),
		compare: cmp,
	}
}

func (h *BinomialHeap[K, V]) Len() int {
	return h.size
}

func (h *BinomialHeap[K, V]) IsEmpty() bool {
	return h.size == 0
}

// Insert добавляет элемент в кучу.
// Алгоритм: создаём дерево степени 0 (один узел) и «складываем» его с кучей
// через mergeSingle, как при сложении двоичных чисел.
// Сложность: O(log n) амортизированная.
func (h *BinomialHeap[K, V]) Insert(key K, val V) {
	node := &BinomialNode[K, V]{key: key, val: val}
	h.mergeSingle(node)
	h.size++
}

// Peek возвращает элемент с наибольшим приоритетом (максимум среди корней).
// В биномиальной куче максимум всегда в одном из корней деревьев.
// Сложность: O(log n) — нужно проверить все деревья.
func (h *BinomialHeap[K, V]) Peek() (K, V, bool) {
	var zeroK K
	var zeroV V
	if h.IsEmpty() {
		return zeroK, zeroV, false
	}

	var maxNode *BinomialNode[K, V]
	for _, r := range h.roots {
		if maxNode == nil || h.compare(r.key, maxNode.key) > 0 {
			maxNode = r
		}
	}
	return maxNode.key, maxNode.val, true
}

// ExtractMax удаляет и возвращает элемент с наибольшим приоритетом.
// Алгоритм:
//  1. Находим корень с максимальным ключом и удаляем его из леса.
//  2. Дочерние узлы удалённого корня — это биномиальные деревья B_0...B_{k-1}.
//     Разворачиваем их (children → sibling) и складываем с оставшейся кучей.
//  3. Уменьшаем размер.
//
// Сложность: O(log n).
func (h *BinomialHeap[K, V]) ExtractMax() (K, V, bool) {
	var zeroK K
	var zeroV V
	if h.IsEmpty() {
		return zeroK, zeroV, false
	}

	var maxDeg int
	var maxNode *BinomialNode[K, V]
	for deg, r := range h.roots {
		if maxNode == nil || h.compare(r.key, maxNode.key) > 0 {
			maxDeg = deg
			maxNode = r
		}
	}

	key, val := maxNode.key, maxNode.val
	delete(h.roots, maxDeg)

	// Разворачиваем детей и сливаем с кучей.
	// Дети удалённого корня — деревья B_0..B_{k-1} в цепочке sibling.
	child := maxNode.child
	for child != nil {
		next := child.sibling
		child.sibling = nil
		h.mergeSingle(child)
		child = next
	}

	h.size--
	return key, val, true
}

func (h *BinomialHeap[K, V]) Merge(other *BinomialHeap[K, V]) {
	if other.IsEmpty() {
		return
	}
	for _, node := range other.roots {
		h.mergeSingle(node)
	}
	h.size += other.size
	other.roots = make(map[int]*BinomialNode[K, V])
	other.size = 0
}

// mergeSingle — вставляет одно биномиальное дерево в кучу.
// Аналогия с двоичным сложением:
//   - Если дерева нужной степени нет → ставим его.
//   - Если уже есть → связываем два дерева одной степени в дерево степени+1 (carry).
//     Carry продолжает «двигаться» вверх по степеням, пока не найдёт свободный слот.
//
// Сложность: O(log n) — дерево может «прокатиться» через все степени.
func (h *BinomialHeap[K, V]) mergeSingle(node *BinomialNode[K, V]) {
	carry := node
	for carry != nil {
		deg := carry.degree
		if existing, ok := h.roots[deg]; ok {
			// Два дерева одной степени → связываем в дерево степени deg+1
			delete(h.roots, deg)
			carry = linkTrees(existing, carry, h.compare)
		} else {
			// Свободный слот → ставим дерево и завершаем
			h.roots[deg] = carry
			carry = nil
		}
	}
}

// linkTrees связывает два биномиальных дерева одной степени.
// Больший корень становится родителем, меньший — первым ребёнком.
// Увеличивает degree на 1.
// Сложность: O(1).
func linkTrees[K any, V any](a, b *BinomialNode[K, V], cmp func(a, b K) int) *BinomialNode[K, V] {
	if cmp(a.key, b.key) < 0 {
		a, b = b, a
	}
	b.sibling = a.child
	a.child = b
	a.degree++
	return a
}
