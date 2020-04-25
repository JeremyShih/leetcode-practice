// 2020/4/18
package main

import (
	"fmt"
	// "math"
	"container/heap"
	"time"
)

func main() {
	start := time.Now()
	m := MedianFinder{}
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian() == 1.5)
	m.AddNum(3)
	fmt.Println(m.FindMedian() == 2)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	var x interface{}
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

type MedianFinder struct {
	even bool
	minh *minHeap
	maxh *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	minh := &minHeap{}
	heap.Init(minh)
	maxh := &minHeap{}
	heap.Init(maxh)
	return MedianFinder{
		minh: minh,
		maxh: maxh,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	fmt.Println(num)
	heap.Push(mf.minh, num)
	heap.Push(mf.maxh, -heap.Remove(mf.minh, 0).(int))

	if mf.minh.Len() < mf.maxh.Len() {
		heap.Push(mf.minh, -heap.Remove(mf.maxh, 0).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	var median float64
	if mf.minh.Len() > mf.maxh.Len() {
		median = float64((*mf.minh)[0])
	} else {
		median = float64((*mf.minh)[0]-((*mf.maxh)[0])) / 2
	}
	return median
}
