package collection

func NewLinkedListStack() Stack {
	return &linkedListStack{
		list: &linkedList{},
	}
}

type linkedListStack struct {
	list 	*linkedList
}

func (s *linkedListStack) Push(item interface{}) {
	s.list.Add(item)
}

func (s *linkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	item, _ := s.list.Get(s.Size() - 1)
	defer s.list.RemoveIndex(s.Size() - 1)
	return item
}

func (s *linkedListStack) Size() int64 {
	return s.list.Size()
}

func (s *linkedListStack) IsEmpty() bool {
	return s.list.IsEmpty()
}



