package string

import (
	"testing"
)

func TestIndex(t *testing.T) {
	// i := IndexBF("123456789", "789", 1)
	// t.Log(i)

	// i = IndexBF("一二三四五六七八九", "七八九", 1)
	// t.Log(i)

	// i = IndexBF("abcdefghi", "ghi", 1)
	// t.Log(i)

	i := IndexBF("acabaabaabcacaabc", "abaabcac", 1)
	t.Log(i)

	i = IndexKMP("acabaabaabcacaabc", "abaabcac", 1)
	t.Log(i)

	i = IndexKMP("一二三四五六七八九", "七八九", 1)
	t.Log(i)

	i = IndexKMP("abcdefghi", "ghi", 1)
	t.Log(i)
}
