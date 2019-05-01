package collections

import "fmt"

type DequeManager struct {
	grow                    func(pivot int) (capacity int)
	pivot, length, capacity int
}

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

func (dm *DequeManager) Head() int {
	if dm.length < 1 {
		return -1
	}
	return dm.pivot
}

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

func (dm *DequeManager) PushTail() int {
	if !dm.maybeGrow() {
		return -1
	}
	dm.length++
	return dm.Tail()
}

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

func (dm *DequeManager) PopTail() int {
	if dm.length < 1 {
		return -1
	}
	tail := dm.Tail()
	dm.length--
	return tail
}
