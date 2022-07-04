// Copyright (c) 2022, Kinson Chow. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package priorityqueue implements a priority queue backed by binary queue.
//
// An unbounded priority queue based on a priority queue.
// The elements of the priority queue are ordered by a comparator provided at queue construction time.
//
// The heap of this queue is the least/smallest element with respect to the specified ordering.
// If multiple elements are tied for least value, the heap is one of those elements arbitrarily.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/Priority_queue
package priorityqueue

import (
	"fmt"
	"github.com/kcswag/kcgods/trees/binaryheap"
	"github.com/kcswag/kcgods/utils"
	"strings"
)

// Assert Queue implementation
//var _ queues.Queue = (*Queue)(nil)

// Queue holds elements in an array-list
type Queue[E any] struct {
	heap       *binaryheap.Heap[E]
	Comparator utils.Comparator
}

// NewWith instantiates a new empty queue with the custom comparator.
func NewWith[E any](comparator utils.Comparator) *Queue[E] {
	return &Queue[E]{heap: binaryheap.NewWith[E](comparator), Comparator: comparator}
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[E]) Enqueue(value E) {
	queue.heap.Push(value)
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[E]) Dequeue() (value E, ok bool) {
	return queue.heap.Pop()
}

// Peek returns top element on the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[E]) Peek() (value E, ok bool) {
	return queue.heap.Peek()
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[E]) Empty() bool {
	return queue.heap.Empty()
}

// Size returns number of elements within the queue.
func (queue *Queue[E]) Size() int {
	return queue.heap.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[E]) Clear() {
	queue.heap.Clear()
}

// Values returns all elements in the queue.
func (queue *Queue[E]) Values() []E {
	return queue.heap.Values()
}

// String returns a string representation of container
func (queue *Queue[E]) String() string {
	str := "PriorityQueue\n"
	values := make([]string, queue.heap.Size(), queue.heap.Size())
	for index, value := range queue.heap.Values() {
		values[index] = fmt.Sprintf("%v", value)
	}
	str += strings.Join(values, ", ")
	return str
}
