package queue

type NonResizeSliceQueue []*Element

func (q *NonResizeSliceQueue) Enqueue(e *Element) {
	*q = append(*q, e)
}

func (q *NonResizeSliceQueue) Dequeue() *Element {
	if len(*q) == 0 {
		return nil
	}
	e := (*q)[0]
	*q = (*q)[1:]
	return e
}
