package interfaces

import (
	"fmt"
	"sort"
)

type Sequence []int

// 类型转换后调用方法
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
