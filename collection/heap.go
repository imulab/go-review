package collection

import (
	"github.com/imulab/go-review/object"
)

func NewMaxPriorityQueue() PriorityQueue {
	return &heap{
		list:      newArrayList(),
		direction: Descending,
	}
}

func NewMinPriorityQueue() PriorityQueue {
	return &heap{
		list:      newArrayList(),
		direction: Ascending,
	}
}

type heap struct {
	list      *arrayList
	direction Direction
}

func (h *heap) EnqueueWithPriority(item interface{}, priority object.Comparable) {
	h.enqueue(&heapItem{
		val:      item,
		priority: priority,
	})
}

func (h *heap) Enqueue(item object.Comparable) {
	h.enqueue(&heapItem{
		val:      item,
		priority: item,
	})
}

func (h *heap) enqueue(item *heapItem) {
	h.list.Add(item)
	h.promote(h.Size() - 1)
}

func (h *heap) Dequeue() interface{} {
	if h.IsEmpty() {
		return nil
	}

	var item interface{}
	{
		first, _ := h.list.Get(0)
		item = first.(*heapItem).val
	}

	h.swap(0, h.Size()-1)
	h.list.RemoveIndex(h.Size() - 1)
	h.demote(0)

	return item
}

func (h *heap) Size() int64 {
	return h.list.Size()
}

func (h *heap) IsEmpty() bool {
	return h.list.IsEmpty()
}

// Perform the promotion operation to eventually put the item at the right place
func (h *heap) promote(i int64) {
	for i != h.parentIndex(i) {
		switch h.direction {
		case Ascending:
			if h.less(i, h.parentIndex(i)) {
				h.swap(i, h.parentIndex(i))
			}
		case Descending:
			if h.greater(i, h.parentIndex(i)) {
				h.swap(i, h.parentIndex(i))
			}
		default:
			panic("impossible direction")
		}
		i = h.parentIndex(i)
	}
}

// Perform the demotion operation to eventually put the item at the right place
func (h *heap) demote(i int64) {
	for {
		candidate := h.leftChildIndex(i)
		{
			if candidate >= h.Size() {
				return
			}

			if h.rightSiblingIndex(candidate) < h.Size() {
				switch h.direction {
				case Ascending:
					if h.greater(candidate, h.rightSiblingIndex(candidate)) {
						candidate = h.rightSiblingIndex(candidate)
					}
				case Descending:
					if h.less(candidate, h.rightSiblingIndex(candidate)) {
						candidate = h.rightSiblingIndex(candidate)
					}
				default:
					panic("impossible direction")
				}
			}
		}

		switch h.direction {
		case Ascending:
			if h.greater(i, candidate) {
				h.swap(i, candidate)
				i = candidate
			} else {
				return
			}
		case Descending:
			if h.less(i, candidate) {
				h.swap(i, candidate)
				i = candidate
			} else {
				return
			}
		default:
			panic("impossible direction")
		}
	}
}

func (h *heap) less(i, j int64) bool {
	if i == j {
		return false
	}

	itemAtI, _ := h.list.Get(i)
	itemAtJ, _ := h.list.Get(j)
	return itemAtI.(*heapItem).compareTo(itemAtJ.(*heapItem)) == object.ComparedLess
}

func (h *heap) greater(i, j int64) bool {
	if i == j {
		return false
	}

	itemAtI, _ := h.list.Get(i)
	itemAtJ, _ := h.list.Get(j)
	return itemAtI.(*heapItem).compareTo(itemAtJ.(*heapItem)) == object.ComparedGreater
}

// Swap the two elements at two different indexes.
// Note that boundary check is not performed
func (h *heap) swap(i, j int64) {
	if i == j {
		return
	}

	itemAtI, _ := h.list.Get(i)
	itemAtJ, _ := h.list.Get(j)
	h.list.Set(i, itemAtJ)
	h.list.Set(j, itemAtI)
}

// Get the index for the left child
// Note that boundary check is not performed
func (h *heap) leftChildIndex(i int64) int64 {
	return i*2 + 1
}

// Get the index for the right child
// Note that boundary check if not performed
func (h *heap) rightChildIndex(i int64) int64 {
	return i*2 + 2
}

// Get the index for parent.
// Note that boundary check if not performed
func (h *heap) parentIndex(i int64) int64 {
	return (i - 1) / 2
}

// Get the index for the right sibling.
// Note that boundary check if not performed
func (h *heap) rightSiblingIndex(i int64) int64 {
	return i + 1
}

type heapItem struct {
	val      interface{}
	priority object.Comparable
}

func (n *heapItem) compareTo(m *heapItem) object.Comparison {
	r, _ := n.priority.CompareTo(m.priority)
	return r
}

type Direction bool

const (
	Ascending  Direction = true
	Descending Direction = false
)
