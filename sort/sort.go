package sort

// 直接插入排序
func InsertSort(ss []int) []int {
	end := len(ss) - 1
	for i := 1; i <= end; i++ {
		j := 0
		for ; j < i; j++ {
			if ss[end] <= ss[j] {
				break
			}
		}
		t := ss[end]
		ss = append(ss[:j+1], ss[j:len(ss)-1]...)
		ss[j] = t
	}
	return ss
}

// 直接插入排序
func InsertSort2(ss []int) {
	for i := 1; i < len(ss); i++ {
		t := ss[i]
		j := i - 1
		for ; j >= 0 && t < ss[j]; j-- {
			ss[j+1] = ss[j]
		}
		ss[j+1] = t
	}
}

// 折半插入排序
func BInsertSort(ss []int) {
	for i := 1; i < len(ss); i++ {
		t := ss[i]
		low := 0
		high := i - 1
		for low <= high {
			m := (low + high) / 2
			if t < ss[m] {
				high = m - 1
			} else {
				low = m + 1
			}
		}

		for j := i - 1; j >= high+1; j-- {
			ss[j+1] = ss[j]
		}
		ss[high+1] = t
	}
}

// 希尔排序
func ShellInsert(ss []int, dk int) {
	for i := dk; i < len(ss); i++ {
		if ss[i] < ss[i-dk] {
			t := ss[i]
			j := i - dk
			for ; j >= 0 && t < ss[j]; j -= dk {
				ss[j+dk] = ss[j]
			}
			ss[j+dk] = t
		}
	}
}

func ShellSort(ss, dt []int) {
	for k := 0; k < len(dt); k++ {
		ShellInsert(ss, dt[k])
	}
}

// 冒泡排序
func BubbleSort(ss []int) {
	m := len(ss) - 1
	flag := 1
	for m > 0 && flag == 1 {
		flag = 0
		for j := 0; j < m; j++ {
			if ss[j] > ss[j+1] {
				flag = 1
				t := ss[j]
				ss[j] = ss[j+1]
				ss[j+1] = t
			}
		}
		m--
	}
}

// 快速排序
func QuickSort(ss []int) {
	qSort(ss, 0, len(ss)-1)
}

func qSort(ss []int, low, high int) {
	if low < high {
		pivotloc := partition(ss, low, high)
		qSort(ss, low, pivotloc-1)
		qSort(ss, pivotloc+1, high)
	}
}

func partition(ss []int, low, high int) int {
	t := ss[low]
	pivotkey := ss[low]
	for low < high {
		for low < high && ss[high] >= pivotkey {
			high--
		}
		ss[low] = ss[high]
		for low < high && ss[low] <= pivotkey {
			low++
		}
		ss[high] = ss[low]
	}
	ss[low] = t
	return low
}

// 选择排序
func SelectSort(ss []int) {
	for i := 0; i < len(ss)-1; i++ {
		k := i
		for j := i + 1; j < len(ss); j++ {
			if ss[j] < ss[k] {
				k = j
			}
		}

		if k != i {
			t := ss[i]
			ss[i] = ss[k]
			ss[k] = t
		}
	}
}

// 堆排序
func HeapSort(ss []int) {
	createHeap(ss)
	for i := len(ss) - 1; i > 0; i-- {
		x := ss[0]
		ss[0] = ss[i]
		ss[i] = x
		heapAdjust(ss, 0, i-1)
	}
}

func createHeap(ss []int) {
	n := len(ss) - 1
	for i := n / 2; i >= 0; i-- {
		heapAdjust(ss, i, n-1)
	}
}

func heapAdjust(ss []int, s, m int) {
	rc := ss[s]
	for j := 2 * s; j <= m; j *= 2 {
		if j < m && ss[j] < ss[j+1] {
			j++
		}

		if rc >= ss[j] {
			break
		}

		ss[s] = ss[j]
		s = j
	}
	ss[s] = rc
}

// 归并排序
func MergeSort(ss []int) []int {
	t := make([]int, len(ss))
	mSort(ss, t, 0, len(ss)-1)
	return t
}

func merge(r, t []int, low, mid, high int) {
	i := low
	j := mid + 1
	k := low
	for i <= mid && j <= high {
		if r[i] <= r[j] {
			t[k] = r[i]
			i++
		} else {
			t[k] = r[j]
			j++
		}
		k++
	}
	for i <= mid {
		t[k] = r[i]
		k++
		i++
	}
	for j <= high {
		t[k] = r[j]
		k++
		j++
	}
}

func mSort(r, t []int, low, high int) {
	if low == high {
		t[low] = r[low]
	} else {
		mid := (low + high) / 2
		s := make([]int, len(r))
		mSort(r, s, low, mid)
		mSort(r, s, mid+1, high)
		merge(s, t, low, mid, high)
	}
}
