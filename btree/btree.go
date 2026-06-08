package btree

// type Node[K any, V any] struct {
// 	keys     []K
// 	values   []V
// 	children []*Node[K, V]
// 	leaf     bool
// }

// type BTree[K any, V any] struct {
// 	root    *Node[K, V]
// 	t       int
// 	compare func(a, b K) int
// }

// func New[K any, V any](t int, cmp func(a, b K) int) *BTree[K, V] {
// 	if t < 2 {
// 		t = 2
// 	}
// 	return &BTree[K, V]{
// 		t:       t,
// 		compare: cmp,
// 	}
// }

// func (tree *BTree[K, V]) newLeaf() *Node[K, V] {
// 	return &Node[K, V]{leaf: true}
// }

// func (tree *BTree[K, V]) newInternal() *Node[K, V] {
// 	return &Node[K, V]{leaf: false}
// }

// func (tree *BTree[K, V]) Get(key K) (V, bool) {
// 	var zero V
// 	node, idx := tree.search(tree.root, key)
// 	if node == nil {
// 		return zero, false
// 	}
// 	return node.values[idx], true
// }

// func (tree *BTree[K, V]) search(node *Node[K, V], key K) (*Node[K, V], int) {
// 	if node == nil {
// 		return nil, -1
// 	}

// 	i := 0
// 	for i < len(node.keys) {
// 		cmp := tree.compare(key, node.keys[i])
// 		if cmp == 0 {
// 			return node, i
// 		}
// 		if cmp < 0 {
// 			break
// 		}
// 		i++
// 	}

// 	if node.leaf {
// 		return nil, -1
// 	}
// 	return tree.search(node.children[i], key)
// }

// func (tree *BTree[K, V]) Insert(key K, value V) {
// 	if tree.root == nil {
// 		tree.root = tree.newLeaf()
// 		tree.root.keys = []K{key}
// 		tree.root.values = []V{value}
// 		return
// 	}

// 	if len(tree.root.keys) == 2*tree.t-1 {
// 		newRoot := tree.newInternal()
// 		newRoot.children = append(newRoot.children, tree.root)
// 		tree.splitChild(newRoot, 0)
// 		tree.root = newRoot
// 	}

// 	tree.insertNonFull(tree.root, key, value)
// }

// func (tree *BTree[K, V]) insertNonFull(node *Node[K, V], key K, value V) {
// 	i := len(node.keys) - 1

// 	if node.leaf {
// 		node.keys = append(node.keys, key)
// 		node.values = append(node.values, value)
// 		for i >= 0 && tree.compare(key, node.keys[i]) < 0 {
// 			node.keys[i+1] = node.keys[i]
// 			node.values[i+1] = node.values[i]
// 			i--
// 		}
// 		node.keys[i+1] = key
// 		node.values[i+1] = value
// 		return
// 	}

// 	for i >= 0 && tree.compare(key, node.keys[i]) < 0 {
// 		i--
// 	}
// 	i++

// 	if len(node.children[i].keys) == 2*tree.t-1 {
// 		tree.splitChild(node, i)
// 		if tree.compare(key, node.keys[i]) > 0 {
// 			i++
// 		}
// 	}
// 	tree.insertNonFull(node.children[i], key, value)
// }
