// Package pq implements a priority queue data structure on top of container/heap.
// As an addition to regular operations, it allows an update of an items priority,
// allowing the queue to be used in graph search algorithms like Dijkstra's algorithm.
// Computational complexities of operations are mainly determined by container/heap.
// In addition, a map of items is maintained, allowing O(1) lookup needed for priority updates,
// which themselves are O(log n).
package gopq

import (
	"container/heap"
)

type PriorityQueue interface {
	Len() int
	Insert(interface{}, int)
	Pop() interface{}
	Contains(interface{}) (int, bool)
	UpdatePriority(interface{}, int)
}

// priorityQueue represents the queue
type priorityQueue struct {
	itemHeap *itemHeap
	lookup   map[interface{}]*item
}

// New initializes an empty priority queue.
func New() PriorityQueue {
	return &priorityQueue{
		itemHeap: &itemHeap{},
		lookup:   make(map[interface{}]*item),
	}
}

// Len returns the number of elements in the queue.
func (p *priorityQueue) Len() int {
	return p.itemHeap.Len()
}

// Insert inserts a new element into the queue. No action is performed on duplicate elements.
func (p *priorityQueue) Insert(v interface{}, priority int) {
	_, ok := p.lookup[v]
	if ok {
		return
	}

	newItem := &item{
		value:    v,
		priority: priority,
	}
	heap.Push(p.itemHeap, newItem)
	p.lookup[v] = newItem
}

// Pop removes the element with the highest priority from the queue and returns it.
// In case of an empty queue, an error is returned.
func (p *priorityQueue) Pop() interface{} {
	if len(*p.itemHeap) == 0 {
		panic("empty queue")
	}

	item := heap.Pop(p.itemHeap).(*item)
	delete(p.lookup, item.value)
	return item.value
}

// UpdatePriority changes the priority of a given item.
// If the specified item is not present in the queue, no action is performed.
func (p *priorityQueue) UpdatePriority(x interface{}, newPriority int) {
	item, ok := p.lookup[x]
	if !ok {
		return
	}

	item.priority = newPriority
	heap.Fix(p.itemHeap, item.index)
}

//A simple method to know if the queue contains a value, and it's priority
func (p *priorityQueue) Contains(x interface{}) (int, bool) {
	i, ok := p.lookup[x]
	if !ok {
		return 0, false
	}
	return i.priority, true
}

type itemHeap []*item

type item struct {
	value    interface{}
	priority int
	index    int
}

func (ih *itemHeap) Len() int {
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].priority < (*ih)[j].priority
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap) Push(x interface{}) {
	it := x.(*item)
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap) Pop() interface{} {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
