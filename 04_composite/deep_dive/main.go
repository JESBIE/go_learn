package main

import "fmt"

func main() {
	// ========== Slice 核心：len vs cap ==========
	fmt.Println("=== len vs cap ===")

	// make([]T, len, cap)
	s := make([]int, 3, 5) // 长度3（能访问的前3个），容量5（底层数组大小）
	fmt.Printf("s=%v  len=%d  cap=%d\n", s, len(s), cap(s))

	// append 在 cap 范围内：不扩容，底层数组不变
	s = append(s, 10)
	fmt.Printf("s=%v  len=%d  cap=%d\n", s, len(s), cap(s))

	s = append(s, 20)
	fmt.Printf("s=%v  len=%d  cap=%d  ← cap满了\n", s, len(s), cap(s))

	// 超过 cap：Go 自动扩容（通常翻倍），分配新数组
	s = append(s, 30)
	fmt.Printf("s=%v  len=%d  cap=%d  ← 扩容了！底层数组换了\n", s, len(s), cap(s))

	// ========== 切片共享底层数组（坑！）==========
	fmt.Println("\n=== 底层数组共享 ===")

	arr := []int{1, 2, 3, 4, 5}
	left := arr[0:3]  // [1 2 3]
	right := arr[2:5] // [3 4 5]

	fmt.Println("left:", left, "right:", right)

	// 改 left 会影响 right！因为它们共享同一个底层数组
	left[2] = 999 // left[2] 就是底层数组的第 2 个位置
	fmt.Println("改 left[2]=999 → left:", left, "right:", right)
	fmt.Println("原数组 arr:", arr)

	// ========== nil slice vs empty slice ==========
	fmt.Println("\n=== nil vs empty ===")
	var nilSlice []int        // nil（ptr=nil, len=0, cap=0）
	emptySlice := []int{}     // 空但非 nil
	makeSlice := make([]int, 0) // 空但非 nil

	fmt.Printf("nilSlice:  %#v  len=%d  cap=%d  ==nil? %v\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice: %#v  len=%d  cap=%d  ==nil? %v\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("makeSlice:  %#v  len=%d  cap=%d  ==nil? %v\n", makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)
}
