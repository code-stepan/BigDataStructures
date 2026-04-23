package btree

import "fmt"

// type Node[K any, V any] struct {
// 	keys     []K
// 	values   []V
// 	children []*Node[K, V]
// 	isleaf   bool
// }

// type BTree[K any, V any] struct {
// 	root    *Node[K, V]
// 	t       int
// 	compare func(K, K) int
// }

// func (bt *BTree[K, V]) inorderRecursive(root *Node[K, V], v *[]V) {
// 	panic("unimplemented")
// }

// func New[K any, V any](t int, cmp func(K, K) int) (*BTree[K, V], error) {
// 	if t < 2 {
// 		return nil, errors.New("Минимальное t > 2")
// 	}
// 	return &BTree[K, V]{
// 		root: &Node[K, V]{
// 			isleaf:   true,
// 			keys:     make([]K, 0, 2*t-1),
// 			values:   make([]V, 0, 2*t-1),
// 			children: make([]*Node[K, V], 0, 2*t),
// 		},
// 		t:       t,
// 		compare: cmp,
// 	}, nil
// }

// func (bt *BTree[K, V]) Destroy() {
// 	bt.root = nil
// }

// func (bt *BTree[K, V]) Search(key K) (V, bool) {
// 	var zero V
// 	node := bt.root
// 	for node != nil {
// 		i := 0
// 		for i < len(node.keys) && bt.compare(key, node.keys[i]) > 0 {
// 			i++
// 		}
// 		if i < len(node.keys) && bt.compare(key, node.keys[i]) == 0 {
// 			return node.values[i], true
// 		}
// 		if node.isleaf {
// 			return zero, false
// 		}
// 		node = node.children[i]
// 	}
// 	return zero, false
// }

// func (bt *BTree[K, V]) Insert(key K, value V) {
// 	root := bt.root
// 	if len(root.keys) == 2*bt.t-1 {
// 		newRoot := &Node[K, V]{
// 			isleaf:   false,
// 			keys:     make([]K, 0, 2*bt.t-1),
// 			values:   make([]V, 0, 2*bt.t-1),
// 			children: make([]*Node[K, V], 0, 2*bt.t),
// 		}
// 		newRoot.children = append(newRoot.children, root)
// 		bt.splitChild(newRoot, Node)
// 		bt.root = newRoot
// 		bt.insertNonFull(newRoot, key, value)
// 	} else {
// 		bt.insertNonFull(root, key, value)
// 	}
// }

// func (bt *BTree[K, V]) Delete(key K) (V, bool) {
// 	if bt.root == nil || len(bt.root.keys) == 0 {
// 		var zero V
// 		return zero, false
// 	}

// 	val, found := bt.deleteRecursive(bt.root, key)

// 	if len(bt.root.keys) == 0 && !bt.root.isleaf {
// 		bt.root = bt.root.children[0]
// 	}
// 	return val, found
// }

// func (bt *BTree[K, V]) InorderTraversal() []V {
// 	result := make([]V, 0)
// 	bt.inorderRecursive(bt.root, &result)
// 	return result
// }

// Доп методы

// func (bt *BTree[K, V]) splitChild(parent *Node[K, V], i int) {
// 	t := bt.t
// 	y := parent.children[i]
// 	z := &Node[K, V]{
// 		isleaf:   y.isleaf,
// 		keys:     make([]K, 0, 2*t-1),
// 		values:   make([]V, 0, 2*t-1),
// 		children: make([]*Node[K, V], 0, 2*bt.t),
// 	}

// 	midKey := y.keys[t-1]
// 	midVal := y.values[t-1]

// }

func StartBTree() {
	fmt.Println("Btree")
}
