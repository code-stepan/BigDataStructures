package segmenttree

type SegmentTree[T any] struct {
	n    int
	tree []T
	op   func(a, b T) T
	id   T
}

func New[T any](n int, op func(a, b T) T, id T) *SegmentTree[T] {
	return &SegmentTree[T]{
		n:    n,
		tree: make([]T, 4*n),
		op:   op,
		id:   id,
	}
}

func (st *SegmentTree[T]) Build(arr []T) {
	st.build(1, 0, st.n-1, arr)
}

func (st *SegmentTree[T]) build(v, tl, tr int, arr []T) {
	if tl == tr {
		st.tree[v] = arr[tl]
		return
	}
}
