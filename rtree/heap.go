package rtree

import (
	"fmt"
	"strings"
)

type MaxHeap struct {
	array    []int64
	capacity int // 容量
	size     int // 当前堆熔炼
}

func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		array:    make([]int64, capacity, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (heap *MaxHeap) IsFull() bool {
	return heap.size >= heap.capacity
}

func (heap *MaxHeap) IsEmpty() bool {
	return heap.size <= 0
}

func (heap *MaxHeap) Put(value int64) bool {
	if heap.IsFull() {
		return false
	}
	// 放在堆最后
	heap.array[heap.size] = value
	// 与父亲节点比较，自下而上堆化
	valueIdx := heap.size
	fatherIdx := (valueIdx - 1) / 2
	for heap.array[valueIdx] > heap.array[fatherIdx] {
		heap.array[valueIdx], heap.array[fatherIdx] = heap.array[fatherIdx], heap.array[valueIdx]
		valueIdx = fatherIdx
		fatherIdx = (valueIdx - 1) / 2
	}
	heap.size++
	return true
}

func (heap *MaxHeap) Pop() int64 {
	if heap.IsEmpty() {
		return -1 // 暂时用 -1 代替无效值
	}
	popVal := heap.array[0]
	// 堆尾元素放在堆顶
	heap.array[0] = heap.array[heap.size-1]
	heap.size--
	// 与子节点比较，自上而下堆化
	lastIdx := 0
	maxIdx := lastIdx
	leftIdx := 2*maxIdx + 1
	rightIdx := 2*maxIdx + 2
	for {
		if leftIdx < heap.size && heap.array[leftIdx] > heap.array[maxIdx] {
			maxIdx = leftIdx
		}
		if rightIdx < heap.size && heap.array[rightIdx] > heap.array[maxIdx] {
			maxIdx = rightIdx
		}
		if maxIdx == lastIdx {
			break
		}
		heap.array[lastIdx], heap.array[maxIdx] = heap.array[maxIdx], heap.array[lastIdx]
		lastIdx = maxIdx
		leftIdx = 2*maxIdx + 1
		rightIdx = 2*maxIdx + 2
	}

	return popVal
}

func (heap *MaxHeap) String() string {
	var sb strings.Builder
	for i := 0; i < heap.size; i++ {
		sb.WriteString(fmt.Sprintf(" %d", heap.array[i]))
	}
	return sb.String()
}
