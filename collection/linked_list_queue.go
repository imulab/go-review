package collection

func NewLinkedListQueue() Queue {
	return &linkedListQueue{
		list: &linkedList{},
	}
}

type linkedListQueue struct {
	list 	*linkedList
}

func (q *linkedListQueue) Enqueue(item interface{}) {
	q.list.Add(item)
}

func (q *linkedListQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item, _ := q.list.Get(0)
	defer q.list.RemoveIndex(0)
	return item
}

func (q *linkedListQueue) Size() int64 {
	return q.list.Size()
}

func (q *linkedListQueue) IsEmpty() bool {
	return q.list.IsEmpty()
}

