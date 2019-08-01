package collection

import (
	"github.com/imulab/go-review/object"
)

func NewAVLTreeTable() Table {
	return &avlTree{}
}

type avlTree struct {
	root *avlNode
	size int64
}

func (t *avlTree) Put(key object.Comparable, value interface{}) {
	t.root = t.put(t.root, key, value)
}

func (t *avlTree) put(x *avlNode, key object.Comparable, value interface{}) *avlNode {
	if x == nil {
		t.size++
		return newAVLNode(key, value)
	}

	switch t.compareTo(key, x.key) {
	case object.ComparedLess:
		x.left = t.put(x.left, key, value)
	case object.ComparedGreater:
		x.right = t.put(x.right, key, value)
	default:
		x.value = value
		return x
	}

	balance := x.balance()

	// left - left
	if balance > 1 && t.compareTo(key, x.left.key) == object.ComparedLess {
		return t.rightRotate(x)
	}

	// left - right
	if balance > 1 && t.compareTo(key, x.left.key) == object.ComparedGreater {
		x.left = t.leftRotate(x.left)
		return t.rightRotate(x)
	}

	// right - right
	if balance < -1 && t.compareTo(key, x.right.key) == object.ComparedGreater {
		return t.leftRotate(x)
	}

	// right - left
	if balance < -1 && t.compareTo(key, x.right.key) == object.ComparedLess {
		x.right = t.rightRotate(x.right)
		return t.leftRotate(x)
	}

	return x
}

func (t *avlTree) Get(key object.Comparable) interface{} {
	cursor := t.root

	for cursor != nil {
		cmp, _ := key.CompareTo(cursor.key)
		switch cmp {
		case object.ComparedLess:
			cursor = cursor.left
		case object.ComparedGreater:
			cursor = cursor.right
		default:
			return cursor.value
		}
	}

	return nil
}

func (t *avlTree) Del(key object.Comparable) {
	t.root = t.del(t.root, key)
}

func (t *avlTree) del(x *avlNode, key object.Comparable) *avlNode {
	if x == nil {
		return nil
	}

	switch t.compareTo(key, x.key) {
	case object.ComparedLess:
		x.left = t.del(x.left, key)
	case object.ComparedGreater:
		x.right = t.del(x.right, key)
	default:
		x0 := x
		x = t.min(x.right)
		if x != nil {
			x.right = t.delMin(x0.right)
			x.left = x0.left
			x0.left = nil
			x0.right = nil
		}
		t.size--
	}

	balance := x.balance()

	// left - left
	if balance > 1 && x.left.balance() >= 0 {
		return t.rightRotate(x)
	}

	// left - right
	if balance > 1 && x.left.balance() < 0 {
		x.left = t.leftRotate(x.left)
		return t.rightRotate(x)
	}

	// right - right
	if balance < -1 && x.right.balance() < 0 {
		return t.leftRotate(x)
	}

	// right - left
	if balance < -1 && x.right.balance() >= 0 {
		x.right = t.rightRotate(x.right)
		return t.leftRotate(x)
	}

	return x
}

func (t *avlTree) min(x *avlNode) *avlNode {
	if x == nil {
		return nil
	}
	cursor := x
	for cursor.left != nil {
		cursor = cursor.left
	}
	return cursor
}

func (t *avlTree) delMin(x *avlNode) *avlNode {
	if x == nil {
		return nil
	}
	if x.left == nil {
		return x.right
	}
	x.left = t.delMin(x.left)
	return x
}

func (t *avlTree) Size() int64 {
	return t.size
}

func (t *avlTree) IsEmpty() bool {
	return t.size == 0
}

func (t *avlTree) leftRotate(h *avlNode) *avlNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.updateHeight()
	return x
}

func (t *avlTree) rightRotate(h *avlNode) *avlNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.updateHeight()
	return x
}

func (t *avlTree) compareTo(a, b object.Comparable) object.Comparison {
	cmp, _ := a.CompareTo(b)
	return cmp
}

func newAVLNode(key object.Comparable, value interface{}) *avlNode {
	return &avlNode{
		key:   key,
		value: value,
		h:     1,
	}
}

type avlNode struct {
	key   object.Comparable
	value interface{}
	left  *avlNode
	right *avlNode
	h     int64
}

func (n *avlNode) updateHeight() {
	if n == nil {
		return
	}

	h := n.left.height()
	if n.right.height() > h {
		h = n.right.height()
	}
	n.h = 1 + h
}

func (n *avlNode) height() int64 {
	if n == nil {
		return 0
	}
	return n.h
}

func (n *avlNode) balance() int64 {
	if n == nil {
		return 0
	}
	return n.left.height() - n.right.height()
}
