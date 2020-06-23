package string

func IndexBF(s, t string, pos int) int {
	i := pos - 1
	j := 0

	sRune := []rune(s)
	tRune := []rune(t)

	for i < len(sRune) && j < len(tRune) {
		if sRune[i] == tRune[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j >= len(tRune) {
		return i + 1 - len(tRune)
	} else {
		return 0
	}
}

func IndexKMP(s, t string, pos int) int {
	i := pos - 1
	j := 0

	sRune := []rune(s)
	tRune := []rune(t)

	next := get_next(tRune)

	for i < len(sRune) && j < len(tRune) {
		if j == -1 || sRune[i] == tRune[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j >= len(tRune) {
		return i + 1 - len(tRune)
	} else {
		return 0
	}
}

func get_next(tRune []rune) []int {
	i := 0
	next := make([]int, len(tRune))
	next[0] = -1
	j := -1

	length := len(tRune)
	for i < length-1 {
		if j == -1 || tRune[i] == tRune[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}

	return next
}
