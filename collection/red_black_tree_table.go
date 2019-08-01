package collection

import (
	"github.com/imulab/go-review/object"
)

func NewRedBlackTreeTable() Table {
	return &redBlackTree{}
}

type redBlackTree struct {
	root	*rbtNode
	size 	int64
}

func (t *redBlackTree) Put(key object.Comparable, value interface{}) {
	t.root = t.put(t.root, key, value)
}

func (t *redBlackTree) put(x *rbtNode, key object.Comparable, value interface{}) *rbtNode {
	if x == nil {
		t.size++
		return newRBTNode(key, value)
	}

	cmp, _ := key.CompareTo(x.key)
	switch cmp {
	case object.ComparedLess:
		x.left = t.put(x.left, key, value)
	case object.ComparedGreater:
		x.right = t.put(x.right, key, value)
	default:
		x.value = value
	}

	x = t.adjust(x)
	return x
}

func (t *redBlackTree) Get(key object.Comparable) interface{} {
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

func (t *redBlackTree) Del(key object.Comparable) {
	t.root = t.del(t.root, key)
}

func (t *redBlackTree) del(x *rbtNode, key object.Comparable) *rbtNode {
	if x == nil {
		return x
	}

	cmp, _ := key.CompareTo(x.key)
	switch cmp {
	case object.ComparedLess:
		x.left = t.del(x.left, key)
		return x
	case object.ComparedGreater:
		x.right = t.del(x.right, key)
		return x
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
		return t.adjust(x)
	}
}

func (t *redBlackTree) min(x *rbtNode) *rbtNode {
	if x == nil {
		return nil
	}

	cursor := x
	for cursor.left != nil {
		cursor = cursor.left
	}
	return cursor
}

func (t *redBlackTree) delMin(x *rbtNode) *rbtNode {
	if x == nil {
		return nil
	}

	if x.left == nil {
		return x.right
	}
	x.left = t.delMin(x.left)
	return x
}

func (t *redBlackTree) Size() int64 {
	return t.size
}

func (t *redBlackTree) IsEmpty() bool {
	return t.size == 0
}

func (t *redBlackTree) adjust(x *rbtNode) *rbtNode {
	if x == nil {
		return x
	}

	if t.isRed(x.right) && !t.isRed(x.left) {
		x = t.leftRotate(x)
	}

	if t.isRed(x.left) && t.isRed(x.left.left) {
		x = t.rightRotate(x)
	}

	if t.isRed(x.left) && t.isRed(x.right) {
		t.flipColor(x)
	}

	return x
}

func (t *redBlackTree) leftRotate(h *rbtNode) *rbtNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = red
	return x
}

func (t *redBlackTree) rightRotate(h *rbtNode) *rbtNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = red
	return x
}

func (t *redBlackTree) flipColor(h *rbtNode) {
	h.left.color = black
	h.right.color = black
	h.color = red
}

func (_ *redBlackTree) isRed(node *rbtNode) bool {
	return node != nil && node.color == red
}

func newRBTNode(key object.Comparable, value interface{}) *rbtNode {
	return &rbtNode{
		key:   key,
		value: value,
		color: red,
	}
}

type rbtNode struct {
	key   object.Comparable
	value interface{}
	left  *rbtNode
	right *rbtNode
	color color
}

type color bool

const (
	red   color = true
	black color = false
)
