package rtree

import "testing"

func TestMapHeap(t *testing.T) {
	values := []int64{9, 10, 7, 3, 5, 6, 4, 2, 1, 8}

	maxHeap := NewMaxHeap(10)
	for _, val := range values {
		suc := maxHeap.Put(val)
		t.Logf("put(%d)=%t", val, suc)
	}
	t.Log(maxHeap)
	for !maxHeap.IsEmpty() {
		val := maxHeap.Pop()
		t.Logf("pop: %d, heap: %s", val, maxHeap)
	}
}
