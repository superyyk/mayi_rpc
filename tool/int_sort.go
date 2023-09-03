package tool

import "sort"

// 升序
func IntSort(ll []int) []int {
	sort.Ints(ll)
	return ll
}

// 降序
func IntRsort(ll []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(ll)))
	return ll
}
