package queue
import (
	"testing"
)

const enqueue_num = 1024

type queue interface {
	Enqueue(*Element)
	Dequeue() *Element
}

func benchmarkQueue(b *testing.B, q queue) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < enqueue_num; j++ {
			q.Enqueue(&Element{i})
		}
		for j := 0; j < enqueue_num; j++ {
			q.Dequeue()
		}
	}
}

func BenchmarkNonResizeSliceQueue(b *testing.B) {
	benchmarkQueue(b, new(NonResizeSliceQueue))
}
