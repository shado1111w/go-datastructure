package sort

import (
	"testing"
)

func TestInsertSort2(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 27, 200}
	t.Log(ss)
	InsertSort2(ss)
	t.Log(ss)
}

func TestInsertSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 27, 200}
	t.Log(ss)
	ss = InsertSort(ss)
	t.Log(ss)
}

func TestBInsertSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 27, 200}
	t.Log(ss)
	BInsertSort(ss)
	t.Log(ss)
}

func TestShellSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 04}
	dt := []int{5, 3, 1}
	t.Log(ss)
	ShellSort(ss, dt)
	t.Log(ss)
}

func TestBubbleSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 27, 200}
	t.Log(ss)
	BubbleSort(ss)
	t.Log(ss)
}

func TestQuickSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 27, 200}
	t.Log(ss)
	QuickSort(ss)
	t.Log(ss)
}

func TestSelectSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 200, 27}
	t.Log(ss)
	SelectSort(ss)
	t.Log(ss)
}

func TestHeapSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 200, 27}
	t.Log(ss)
	HeapSort(ss)
	t.Log(ss)
}

func TestMergeSort(t *testing.T) {
	ss := []int{49, 38, 65, 97, 76, 13, 200, 27}
	t.Log(ss)
	s := MergeSort(ss)
	t.Log(s)
}
