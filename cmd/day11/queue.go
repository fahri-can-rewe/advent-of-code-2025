package main

// Generic growable ring-buffer queue (Go 1.18+).
type Queue[T any] struct {
	buf        []T
	head, size int
}

// New creates an empty queue.
func New[T any]() *Queue[T] { return &Queue[T]{} }

// Len returns the number of elements.
func (q *Queue[T]) Len() int { return q.size }

// IsEmpty reports whether the queue has no elements.
func (q *Queue[T]) IsEmpty() bool { return q.size == 0 }

// Enqueue pushes v to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	if q.buf == nil {
		q.buf = make([]T, 8)
	}
	if q.size == len(q.buf) {
		q.grow()
	}
	idx := (q.head + q.size) % len(q.buf)
	q.buf[idx] = v
	q.size++
}

// Dequeue pops from the front. The bool is false if the queue was empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	v := q.buf[q.head]
	q.buf[q.head] = zero // avoid memory leak for reference types
	q.head = (q.head + 1) % len(q.buf)
	q.size--
	return v, true
}

// Peek returns the front element without removing it. The bool is false if empty.
func (q *Queue[T]) Peek() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	return q.buf[q.head], true
}

// Clear removes all elements.
func (q *Queue[T]) Clear() {
	// nil out slice to allow GC of referenced objects
	for i := 0; i < q.size; i++ {
		q.buf[(q.head+i)%len(q.buf)] = *new(T)
	}
	q.buf = nil
	q.head = 0
	q.size = 0
}

func (q *Queue[T]) grow() {
	old := q.buf
	n := len(old) * 2
	if n == 0 {
		n = 8
	}
	q.buf = make([]T, n)
	if q.size == 0 {
		q.head = 0
		return
	}
	// copy elements in order starting at new head 0
	for i := 0; i < q.size; i++ {
		q.buf[i] = old[(q.head+i)%len(old)]
	}
	q.head = 0
}
