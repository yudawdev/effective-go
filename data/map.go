package data

import (
	"fmt"
)

// map 初始化
// map: map[]
// ----------------------
// map: map[k1:v1 k2:v2]
func mapInit() {

	// map 初始化
	m := make(map[string]string)

	// 第二种方式初始化
	m2 := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}

	fmt.Printf("map: %v\n", m)
	fmt.Println("----------------------")
	fmt.Printf("map: %v\n", m2)
}

// map 遍历
func mapIterator() {
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"a1": "a2",
	}

	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	for k := range m {
		fmt.Printf("key: %s, value: %s\n", k, m[k])
	}

	for _, v := range m {
		fmt.Printf("value: %s\n", v)
	}
}

// map 删除元素
// key: k1, exist: false, value:
func mapDeleteElement() {
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"a1": "a2",
	}

	key := "k1"

	delete(m, key)

	value, ok := m[key]

	fmt.Printf("key: %s, exist: %v, value: %s ", key, ok, value)
}

// map 判断元素是否存在
// key: k1, exist: true
func mapContains() {
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"a1": "a2",
	}

	key := "k1"

	_, ok := m[key]

	fmt.Printf("key: %s, exist: %v", key, ok)
}
