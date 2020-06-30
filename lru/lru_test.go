package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	fmt.Println(c.Bucket, c.Tail)
	c.Put(4, 4)
	fmt.Println(c.Bucket, c.Tail)
	val := c.Get(3)
	fmt.Println(val)
	fmt.Println(c.Bucket, c.Tail)
}
