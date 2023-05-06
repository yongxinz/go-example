package main

import "fmt"

func partition(list []int, low, high int) int {
	// 选择基准值
	pivot := list[high]
	for low < high {
		// low 指针值 <= pivot low 指针右移
		for low < high && pivot >= list[low] {
			low++
		}
		// low 指针值 > pivot low 值移到 high 位置
		list[high] = list[low]

		// high 指针值 >= pivot high 指针左移
		for low < high && pivot <= list[high] {
			high--
		}
		// high 指针值 < pivot high 值移到 low 位置
		list[low] = list[high]
	}
	// pivot 替换 high 值
	list[high] = pivot
	return high
}

func QuickSort(list []int, low, high int) {
	if high > low {
		// 分区
		pivot := partition(list, low, high)
		// 左边部分排序
		QuickSort(list, low, pivot-1)
		// 右边部分排序
		QuickSort(list, pivot+1, high)
	}
}

func main() {
	list := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
