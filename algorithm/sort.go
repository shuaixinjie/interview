package algorithm

// 排序算法

// BubbleSort 冒泡排序
func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// SelectSort 选择排序
func SelectSort(s []int) {
	for i := 0; i < len(s); i++ {
		minIndex := i
		for j := i; j < len(s); j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}

// InsertSort 插入排序
func InsertSort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
}

// QuickSort 快速排序
func QuickSort(s []int) {
	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, start, end int) {
	if start < end {
		i, j := start, end
		key := s[(start+end)/2]
		for i <= j {
			for s[i] < key {
				i++
			}
			for s[j] > key {
				j--
			}
			if i <= j {
				s[i], s[j] = s[j], s[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(s, start, j)
		}
		if end > i {
			quickSort(s, i, end)
		}
	}
}
