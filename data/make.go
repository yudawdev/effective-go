package data

import "fmt"

func testMake() {

	// It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).

	// slices 初始化 (切片类型，初始长度，容量)
	// slices len: 10, cap: 100
	slices := make([]string, 10, 100)
	fmt.Printf("slices len: %d, cap: %d", len(slices), cap(slices))

	//var p = new([]int)       // allocates slice structure; *p == nil; rarely useful
	//var v = make([]int, 100) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	//var p *[]int = new([]int)
	//*p = make([]int, 100, 100)

	// Idiomatic:
	v := make([]int, 100)
	fmt.Printf("v len: %d, cap: %d", len(v), cap(v))
}
