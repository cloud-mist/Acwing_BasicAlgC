package main

import "container/heap"

type heap1 []int

func (h heap1) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heap1) Len() int {
	return len(h)
}

func (h heap1) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *heap1) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *heap1) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func main() {
	a := &heap1{6, 3, 5, 2, 10, 1}
	heap.Init(a)

}
