package collections_test

import (
	"fmt"

	"github.com/carlmjohnson/collections"
)

func output(ss []string, dm *collections.DequeManager) {
	hi, ti, hv, tv := dm.Head(), dm.Tail(), "", ""
	if hi != -1 {
		hv = ss[hi]
		tv = ss[ti]
	}
	fmt.Printf("Head: %d %q Tail: %d %q Strings: %v\n",
		hi, hv, ti, tv, ss)
}

func ExampleDequeManager() {
	var ss []string
	deque := collections.NewDeque(0, func(pivot int) int {
		ns := make([]string, len(ss)*2+1)
		copied := copy(ns, (ss)[pivot:])
		copy(ns[copied:], (ss)[:pivot])
		ss = ns
		return len(ns)
	})

	output(ss, deque)
	ss[deque.PushTail()] = "hello"
	output(ss, deque)
	ss[deque.PushTail()] = "world"
	output(ss, deque)
	ss[deque.PushHead()] = "¡¡"
	ss[deque.PushTail()] = "!!"
	output(ss, deque)
	fmt.Println(ss[deque.PopHead()])
	fmt.Println(ss[deque.PopTail()])
	output(ss, deque)
	// Output:
	// Head: -1 "" Tail: -1 "" Strings: []
	// Head: 0 "hello" Tail: 0 "hello" Strings: [hello]
	// Head: 0 "hello" Tail: 1 "world" Strings: [hello world ]
	// Head: 0 "¡¡" Tail: 3 "!!" Strings: [¡¡ hello world !!   ]
	// ¡¡
	// !!
	// Head: 1 "hello" Tail: 2 "world" Strings: [¡¡ hello world !!   ]
}

func ExampleNewDequeForSlice() {
	var ss []string
	deque := collections.NewDequeForSlice(&ss)
	output(ss, deque)
	ss[deque.PushTail()] = "hello"
	output(ss, deque)
	ss[deque.PushTail()] = "world"
	output(ss, deque)
	ss[deque.PushHead()] = "¡¡"
	ss[deque.PushTail()] = "!!"
	output(ss, deque)
	fmt.Println(ss[deque.PopHead()])
	fmt.Println(ss[deque.PopTail()])
	output(ss, deque)
	// Output:
	// Head: -1 "" Tail: -1 "" Strings: []
	// Head: 0 "hello" Tail: 0 "hello" Strings: [hello]
	// Head: 0 "hello" Tail: 1 "world" Strings: [hello world ]
	// Head: 0 "¡¡" Tail: 3 "!!" Strings: [¡¡ hello world !!   ]
	// ¡¡
	// !!
	// Head: 1 "hello" Tail: 2 "world" Strings: [¡¡ hello world !!   ]
}
