package collections_test

import (
	"fmt"

	"github.com/carlmjohnson/collections"
)

func ExampleDequeManager() {
	var ss []string
	deque := collections.NewDeque(0, func(pivot int) int {
		ns := make([]string, len(ss)*2+1)
		copied := copy(ns, (ss)[pivot:])
		copy(ns[copied:], (ss)[:pivot])
		ss = ns
		return len(ns)
	})

	fmt.Printf("%q %v\n", ss, deque)
	ss[deque.PushTail()] = "hello"
	fmt.Printf("%q %v\n", ss, deque)
	ss[deque.PushTail()] = "world"
	fmt.Printf("%q %v\n", ss, deque)
	fmt.Printf("%q %q %v\n", ss[deque.PopHead()], ss, deque)
	ss[deque.PushHead()] = ","
	ss[deque.PushHead()] = "Hello"
	fmt.Printf("%q %v\n", ss, deque)
	ss[deque.PushTail()] = "!"
	fmt.Printf("%q %v\n", ss, deque)
	// Output:
	// [] collections.DequeManager{head: -1, tail: -1, length: 0, pivot: 0}
	// ["hello"] collections.DequeManager{head: 0, tail: 0, length: 1, pivot: 0}
	// ["hello" "world" ""] collections.DequeManager{head: 0, tail: 1, length: 2, pivot: 0}
	// "hello" ["hello" "world" ""] collections.DequeManager{head: 1, tail: 1, length: 1, pivot: 1}
	// ["," "world" "Hello"] collections.DequeManager{head: 2, tail: 1, length: 3, pivot: 2}
	// ["Hello" "," "world" "!" "" "" ""] collections.DequeManager{head: 0, tail: 3, length: 4, pivot: 0}
}

func ExampleNewDequeForSlice() {
	var ss []string
	deque := collections.NewDequeForSlice(&ss)

	fmt.Printf("%q %v\n", ss, deque)
	ss[deque.PushTail()] = "hello"
	fmt.Printf("%q %v\n", ss, deque)
	ss[deque.PushTail()] = "world"
	fmt.Printf("%q %v\n", ss, deque)
	fmt.Printf("%q %q %v\n", ss[deque.PopHead()], ss, deque)
	ss[deque.PushHead()] = ","
	ss[deque.PushHead()] = "Hello"
	fmt.Printf("%q %v\n", ss, deque)
	ss[deque.PushTail()] = "!"
	fmt.Printf("%q %v\n", ss, deque)
	// Output:
	// [] collections.DequeManager{head: -1, tail: -1, length: 0, pivot: 0}
	// ["hello"] collections.DequeManager{head: 0, tail: 0, length: 1, pivot: 0}
	// ["hello" "world" ""] collections.DequeManager{head: 0, tail: 1, length: 2, pivot: 0}
	// "hello" ["hello" "world" ""] collections.DequeManager{head: 1, tail: 1, length: 1, pivot: 1}
	// ["," "world" "Hello"] collections.DequeManager{head: 2, tail: 1, length: 3, pivot: 2}
	// ["Hello" "," "world" "!" "" "" ""] collections.DequeManager{head: 0, tail: 3, length: 4, pivot: 0}
}
