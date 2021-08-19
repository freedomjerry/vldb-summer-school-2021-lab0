package main

import (
	"container/heap"
	"strconv"
)

type Items struct {
	key string
	priority int
}
type kvheap []*Items
func (h kvheap) Len() int {
	return len(h)
}
func (h kvheap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].key > h[j].key
	}
	return h[i].priority < h[j].priority
}
func (h kvheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *kvheap) Push(x interface{})  {
	*h = append(*h, x.(*Items))
}
func  (h *kvheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}
func (h *kvheap) update(item *Items, key string, priority int) {
	item.key = key
	item.priority = priority
}
func Top10heap(kv []*urlCount) []KeyValue {
	result := make([]KeyValue, 0, 11)
	if kv == nil {
		return result
	}
	kvh := &kvheap{}
	heap.Init(kvh)
	for _, uc := range kv {
		if uc == nil {
			continue
		}
		keyv := &Items{
			key: uc.url,
			priority: uc.cnt,
		}
		heap.Push(kvh, keyv)
		if kvh.Len() > 10 {
			heap.Pop(kvh)
		}

	}
	times := kvh.Len()
	for i:= 0; i < times; i++ {
		item := heap.Pop(kvh).(*Items)
		val := strconv.Itoa(item.priority)
		result = append(result, KeyValue{" ", item.key + " " + val})
	}
	return  result
}