package main

import (
	"balance_example/weighted"
	"fmt"
)

func main() {

	// 平滑加权轮询 smooth weighted round-robin balancing algorithm
	w := &weighted.SW{}
	w.Add("A", 4)
	w.Add("B", 2)
	w.Add("C", 1)

	results := make(map[string]int)

	for i := 0; i < 8; i++ {
		//The weighted is NOT goroutine-safe so you MUST use the synchronization primitive to protect it (the Next method) in concurrent cases.
		s := w.Next().(string)
		fmt.Printf("%s,", s)
		results[s]++
	}
	fmt.Printf("\n %+v", results)
}
