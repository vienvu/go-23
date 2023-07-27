package sorting

import (
	"sort"
)

func init() {}

func SortIntArray(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func SortFloatArray(arr []float64) []float64 {
	sort.Float64s(arr)
	return arr
}

func SortStringArray(arr *[]string) []string {
	sort.Strings(arr)
	return arr
}
