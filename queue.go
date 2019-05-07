package collections

import (
	"fmt"
	"reflect"
)

// DequeManager manages unboxed, type-safe double-ended deques backed by a slice.
type DequeManager struct {
	grow                    func(pivot int) (capacity int)
	pivot, length, capacity int
}

// NewDeque returns a DequeManager ready to manage a slice.
// Initial length of the slice must be given by capacity.
// Grow function should copy the existing slice around the pivot
// to a new slice and return the length of the new slice.
// To signal that the slice cannot be grown (e.g. to use a fixed deque size)
// return -1.
func NewDeque(capacity int, grow func(pivot int) (capacity int)) *DequeManager {
	if capacity < 0 {
		panic(fmt.Errorf("impossible starting capacity: %d", capacity))
	}
	return &DequeManager{
		grow:     grow,
		pivot:    0,
		length:   0,
		capacity: capacity,
	}
}

// Head returns the index of the current head of the slice or -1 if the deque is empty.
func (dm *DequeManager) Head() int {
	if dm.length < 1 {
		return -1
	}
	return dm.pivot
}

// Tail returns the index of the current tail of the slice or -1 if the deque is empty.
func (dm *DequeManager) Tail() int {
	if dm.length < 1 {
		return -1
	}
	return (dm.pivot + dm.length - 1) % dm.capacity
}

func (dm *DequeManager) maybeGrow() bool {
	if dm.length == dm.capacity {
		newcap := dm.grow(dm.pivot)
		if newcap < dm.capacity {
			panic(fmt.Errorf("bad growth capacity for DequeManager: %d < %d",
				newcap, dm.capacity))
		}
		if newcap == -1 {
			return false
		}
		dm.capacity = newcap
		dm.pivot = 0
	}
	return true
}

// PushHead returns the index of the next head of the slice.
// It may call the grow function if necessary.
// Returns -1 if the deque needs growth but cannot be grown.
func (dm *DequeManager) PushHead() int {
	if !dm.maybeGrow() {
		return -1
	}
	dm.length++
	dm.pivot -= 1
	if dm.pivot < 0 {
		dm.pivot = dm.capacity - 1
	}
	return dm.pivot
}

// PushTail returns the index of the next tail of the slice.
// It may call the grow function if necessary.
// Returns -1 if the deque needs growth but cannot be grown.
func (dm *DequeManager) PushTail() int {
	if !dm.maybeGrow() {
		return -1
	}
	dm.length++
	return dm.Tail()
}

// PopHead returns the index of the head of the slice
// and removes it from the deque.
// Returns -1 if the deque is empty.
// Note that actual values in the slice are not affected by deque operations.
func (dm *DequeManager) PopHead() int {
	if dm.length < 1 {
		return -1
	}
	head := dm.Head()
	dm.length--
	dm.pivot++
	if dm.pivot < 0 {
		dm.pivot = dm.capacity - 1
	}
	return head
}

// PopTail returns the index of the tail of the slice
// and removes it from the deque.
// Returns -1 if the deque is empty.
// Note that actual values in the slice are not affected by deque operations.
func (dm *DequeManager) PopTail() int {
	if dm.length < 1 {
		return -1
	}
	tail := dm.Tail()
	dm.length--
	return tail
}

// String implements fmt.Stringer
func (dm *DequeManager) String() string {
	return fmt.Sprintf("collections.DequeManager{head: %d, tail: %d, length: %d, pivot: %d}",
		dm.Head(), dm.Tail(), dm.length, dm.pivot,
	)
}

// NewDequeForSlice is a convenience initializer that uses reflection
// to call NewDeque with an appropriate capacity and growth function.
func NewDequeForSlice(slicepointer interface{}) *DequeManager {
	value := reflect.ValueOf(slicepointer)
	if value.Kind() != reflect.Ptr {
		panic("slicepointer must be a pointer to a slice")
	}
	slice := value.Elem()
	if !slice.IsValid() || slice.Kind() != reflect.Slice {
		panic("slicepointer must be a pointer to a slice")
	}
	return NewDeque(slice.Len(), func(pivot int) int {
		newcap := slice.Len()*2 + 1
		ns := reflect.MakeSlice(slice.Type(), newcap, newcap)
		copied := reflect.Copy(ns, slice.Slice(pivot, slice.Len()))
		reflect.Copy(ns.Slice(copied, newcap), slice.Slice(0, pivot))
		slice = ns
		value.Elem().Set(ns)
		return newcap
	})
}
