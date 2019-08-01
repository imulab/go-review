package collection

import (
	"github.com/imulab/go-review/object"
	"github.com/pkg/errors"
)

func NewLinkedList() List {
	return &linkedList{}
}

// A doubly linked list
type linkedList struct {
	head 	*listNode
	tail	*listNode
	size	int64
}

func (l *linkedList) Add(item interface{}) error {
	if l.size + 1 < 0 {
		return errors.New("maximum elements reached")
	}

	newNode := &listNode{val: item}

	switch l.size {
	case 0:
		l.head = newNode
		l.tail = newNode
		break

	default:
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		break
	}

	l.size++
	return nil
}

func (l *linkedList) RemoveIndex(index int64) error {
	node, err := l.getNodeAtIndex(index)
	if err != nil {
		return err
	}
	l.remove(node)
	return nil
}

func (l *linkedList) Remove(item interface{}) error {
	cursor, _ := l.getNode(item)
	if cursor == nil {
		return errors.New("item not found")
	}
	l.remove(cursor)
	return nil
}

func (l *linkedList) remove(node *listNode) {
	if node == nil {
		return
	}

	if node.prev == nil {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil {
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}

	node.next = nil
	node.prev = nil

	l.size--
}

func (l *linkedList) Size() int64 {
	return l.size
}

func (l *linkedList) IsEmpty() bool {
	return l.Size() == 0
}

func (l *linkedList) Contains(item interface{}) bool {
	return l.IndexOf(item) >= 0
}

func (l *linkedList) getNode(item interface{}) (*listNode, int64) {
	i := int64(0)
	cursor := l.head
	for cursor != nil {
		if cursor.Is(item) {
			return cursor, i
		}
		cursor = cursor.next
		i++
	}
	return nil, -1
}

func (l *linkedList) getNodeAtIndex(index int64) (*listNode, error) {
	if err := l.checkIndex(index); err != nil {
		return nil, err
	}

	if index < l.size / 2 {
		cursor := l.head
		for i := int64(0); i < index; i++ {
			cursor = cursor.next
		}
		return cursor, nil
	} else {
		cursor := l.tail
		for i := l.size - 1; i > index; i-- {
			cursor = cursor.prev
		}
		return cursor, nil
	}
}

func (l *linkedList) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l *linkedList) ToArray() []interface{} {
	array := make([]interface{}, 0, l.size)
	cursor := l.head
	for cursor != nil {
		array = append(array, cursor.val)
		cursor = cursor.next
	}
	return array
}

func (l *linkedList) Get(index int64) (interface{}, error) {
	node, err := l.getNodeAtIndex(index)
	if err != nil {
		return nil, err
	}
	return node.val, nil
}

func (l *linkedList) Set(index int64, item interface{}) error {
	node, err := l.getNodeAtIndex(index)
	if err != nil {
		return err
	}
	node.val = item
	return nil
}

func (l *linkedList) checkIndex(index int64) error {
	if index < 0 || index >= l.size {
		return errors.New("index out of bounds")
	}
	return nil
}

func (l *linkedList) IndexOf(item interface{}) int64 {
	_, i := l.getNode(item)
	return i
}

// A doubly linked list node
type listNode struct {
	prev 	*listNode
	next	*listNode
	val 	interface{}
}

func (n *listNode) Is(item interface{}) bool {
	if obj, ok := n.val.(object.Equality); ok {
		return obj.Equals(item)
	} else {
		return n.val == item
	}
}