package bstree

type Position uint8

const (
	begin, between, end Position = 0, 1, 2
)

type Iterator[K comparable, V any] struct {
	tree     *BSTree[K, V]
	node     *Node[K, V]
	position Position
}

func (tree *BSTree[K, V]) Iter() *Iterator[K, V] {
	return &Iterator[K, V]{
		tree: tree,
	}
}

func (tree *BSTree[K, V]) next(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		// return first node
		return tree.min(tree.root)
	}
	if node.parent == nil {
		// node is root
		return tree.min(node.right)
	}
	switch node {
	case node.parent.left:
		// node is left child
		if node.right != nil {
			return tree.min(node.right)
		} else {
			return node.parent
		}
	case node.parent.right:
		// node is right child
		if node.right != nil {
			return tree.min(node.right)
		} else {
			// end
			return nil
		}
	default:
		panic("tree structure error")
	}
}

func (tree *BSTree[K, V]) prev(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		// return last node
		return tree.max(tree.root)
	}
	if node.parent == nil {
		// node is root
		return tree.max(node.left)
	}
	switch node {
	case node.parent.right:
		// node is right child
		if node.left != nil {
			return tree.max(node.left)
		} else {
			return node.parent
		}
	case node.parent.left:
		// node is left child
		if node.right != nil {
			return tree.max(node.left)
		} else {
			// end
			return nil
		}
	default:
		panic("tree structure error")
	}
}

func (iter *Iterator[K, V]) Begin() {
	iter.node = nil
	iter.position = begin
}

func (iter *Iterator[K, V]) End() {
	iter.node = nil
	iter.position = end
}

func (iter *Iterator[K, V]) Next() bool {
	if iter.position == end {
		return false
	}

	if iter.position == begin {
		iter.position = between
	}

	iter.node = iter.tree.next(iter.node)
	if iter.node == nil {
		iter.position = end
	}

	return iter.position != end
}

func (iter *Iterator[K, V]) Prev() bool {
	if iter.position == begin {
		return false
	}

	if iter.position == end {
		iter.position = between
	}

	iter.node = iter.tree.prev(iter.node)
	if iter.node == nil {
		iter.position = begin
	}

	return iter.position != begin
}

func (iter *Iterator[K, V]) Key() (key K, ok bool) {
	if iter.node == nil {
		return key, false
	}
	return iter.node.key, true
}

func (iter *Iterator[K, V]) Value() (value V, ok bool) {
	if iter.node == nil {
		return value, false
	}
	return iter.node.value, true
}

func (iter *Iterator[K, V]) ValuePtr() (*V, bool) {
	if iter.node == nil {
		return nil, false
	}
	return &iter.node.value, true
}
