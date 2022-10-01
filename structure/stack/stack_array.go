// Stack implementation using Array
// Stack is a linear data structure in which the push and pop operations occur only at one end of the structure, referred to as the top of the stack.
// The order in which an element added to or removed from a stack is described as last in, first out (LIFO).
// Details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)

package stack

import "errors"

var ErrStackEmpty = errors.New("stack is empty")

/*
The methods can also be implemented directly on the slice.
```
type Array[T any] []T
```
However, this exposes the underlaying storage (slice) outside the package.
A struct is used instead, so that the underlying storage is not accessible outside the package.
*/

// Array is an implementation of stack with slice as underlying storage.
// ```
// stack := stack.NewArray[int]()
// ```
type Array[T any] struct {
	store []T
}

func NewArray[T any]() *Array[T] {
	return new(Array[T])
}

// Push inserts a new element to the stack
// Push on a nil stack will panic
func (s *Array[T]) Push(val T) {
	s.store = append(s.store, val)
}

// Peek the last inserted element without removing it from the stack
// If the stack is empty, ErrStackEmpty error is returned
func (s *Array[T]) Peek() (T, error) {
	var element T
	if s.Empty() {
		return element, ErrStackEmpty
	}
	return s.store[s.Len()-1], nil
}

func (s *Array[T]) Len() int {
	if s == nil {
		return 0
	}
	return len(s.store)
}

func (s *Array[T]) Empty() bool {
	return s.Len() == 0
}

// Pop returns last inserted element and removes it from the underlaying storage
// If the stack is empty, ErrStackEmpty error is returned
func (s *Array[T]) Pop() (T, error) {
	var element T
	if s.Empty() {
		return element, ErrStackEmpty
	}
	element = s.store[s.Len()-1]
	s.store = s.store[:s.Len()-1]
	return element, nil
}

// Clear removes all elements.
// The allocated capacity remains the same and will be reused for subsequent push operations
func (s *Array[T]) Clear() {
	if s == nil {
		return
	}
	s.store = s.store[:0]
}

// Truncate removes all elements and underlaying storage
func (s *Array[T]) Truncate() {
	if s == nil {
		return
	}
	s.store = nil
}
