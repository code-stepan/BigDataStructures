package segmenttree

type SegmentTree[T any] struct {
	n    int
	tree []T
	op   func(a, b T) T
	id   T
}

func New[T any](n int, op func(a, b T) T, id T) *SegmentTree[T] {
	tree := make([]T, n*4)
	for i := range tree {
		tree[i] = id
	}
	return &SegmentTree[T]{n: n, tree: tree, op: op, id: id}
}

func (st *SegmentTree[T]) Update(pos int, value T) {
	st.update(1, 0, st.n-1, pos, value)
}

func (st *SegmentTree[T]) update(v, tl, tr, pos int, value T) {
	if tl == tr {
		st.tree[v] = value
		return
	}
	tm := (tl + tr) / 2
	if pos <= tm {
		st.update(2*v, tl, tm, pos, value)
	} else {
		st.update(2*v+1, tm+1, tr, pos, value)
	}
	st.tree[v] = st.op(st.tree[2*v], st.tree[2*v+1])
}
