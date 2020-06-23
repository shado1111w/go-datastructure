package search

func SearchSeq(ss []int, key int) int {
	for i, val := range ss {
		if val == key {
			return i
		}
	}

	return -1
}

func SearchBin(ss []int, key int) int {
	low := 0
	high := len(ss) - 1
	for low <= high {
		mid := (low + high) / 2
		if ss[mid] == key {
			return mid
		} else if key > ss[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func SearchBin2(ss []int, key int) int {
	low := 0
	high := len(ss) - 1
	mid := (low + high) / 2
	if ss[mid] == key {
		return mid
	}

	if high == 0 {
		return -1
	}

	if key > ss[mid] {
		return SearchBin2(ss[mid+1:high+1], key)
	} else {
		return SearchBin2(ss[low:mid], key)
	}
}
