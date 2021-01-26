package code

import "fmt"

/**
 * 冒泡排序
 * 说明：就是第一个位置上的数与他相邻第二个位置上的数比较，
 * 如果比他相邻的数大，则两者交换位置，否则不交换。
 * 接着第二个位置上的数与第三个位置上的数比较大小，也是大则交换，
 * 一直到和最后一个位置的数比较交换完毕。
 * 然后，是下一个循环，就是第二个位置上的数重复上面的比较交换操作，
 * 直到把整个数组变成是一个从小到大的有序序列。
 */
func BubbleSort(arr []int) []int {
	l := len(arr)

	for i := 0; i < l; i++ {
		// 标记此轮是否有交换，没有表示已经有序
		btn := true
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				btn = false
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)

		if btn {
			break
		}
	}

	return arr
}

/**
 * 插入排序
 * 说明：从待排序的数组中选出来一个最小值(可以认为第一个数就是已排序的数组)，
 * 然后从剩余数中选出来最小值有序放到已排序的数组中，
 * 依次操作，直到最后的数组都是一个从小到大的有序数组为止
 */
func InsertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		v := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > v {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = v
		fmt.Println(arr)
	}

	return arr
}

/**
 * 选择排序
 * 说明：从待排序的数组中选出来一个最小值，放到新的数组的第一个位置，
 * 继续从剩余的数组中选取最小值放入到数组中，重复上面的步骤，将数字都取出来排成新的有序数组
 */
func SelectSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = tmp
		fmt.Println(arr)
	}

	return arr
}
