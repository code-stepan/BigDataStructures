package fenwicktree

type Monoid[T any] struct {
	Identity T
	Combine  func(a, b T) T
	Inverse  func(a T) T
}

type FenwickTree[T any] struct {
	tree   []T
	data   []T
	size   int
	monoid Monoid[T]
}
