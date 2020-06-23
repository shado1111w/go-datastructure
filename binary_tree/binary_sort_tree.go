package binary_tree

import (
	"fmt"
)

type BSTree struct {
	data           int
	lChild, rChild *BSTree
}

// 1.二叉排序树的查找
func SearchBST(t *BSTree, key int) *BSTree {
	if t == nil || (key == t.data) {
		return t
	}

	if key < t.data {
		return SearchBST(t.lChild, key)
	} else {
		return SearchBST(t.rChild, key)
	}
}

// 2.二叉排序树的插入
func InsertBST(t *BSTree, e int) bool {
	if t == nil {
		return true
	} else if e < t.data {
		ok := InsertBST(t.lChild, e)
		if ok {
			s := new(BSTree)
			s.data = e
			s.rChild = nil
			s.lChild = s.rChild
			t.lChild = s
		}

		return false
	} else if e > t.data {
		ok := InsertBST(t.rChild, e)
		if ok {
			s := new(BSTree)
			s.data = e
			s.rChild = nil
			s.lChild = s.rChild
			t.rChild = s
		}

		return false
	}

	return false
}

// 3.二叉排序树的创建
func CreateBST(t *BSTree) {
	var e int
	fmt.Scan(&e)
	t.data = e

	fmt.Println(*t)

	fmt.Scan(&e)
	for e != 520 {
		InsertBST(t, e)
		fmt.Println(*t)
		fmt.Scan(&e)
	}
}

// 4.二叉排序树的删除
func DeleteBST(t *BSTree, key int) {
	p := t
	f := new(BSTree)

	for p != nil {
		if p.data == key {
			break
		}

		// f为p的双亲结点
		f = p
		if key < p.data {
			p = p.lChild
		} else {
			p = p.rChild
		}
	}

	// 找不到被删除结点
	if p == nil {
		return
	}

	q := p
	if p.lChild != nil && p.rChild != nil {
		s := p.lChild
		for s.rChild != nil {
			q = s
			s = s.rChild
		}

		p.data = s.data
		if q != p {
			q.rChild = s.lChild
		} else {
			q.lChild = s.lChild
		}
		return
	} else if p.rChild == nil {
		p = p.lChild
	} else if p.lChild == nil {
		p = p.rChild
	}

	if f == nil {
		t = p
	} else if q == f.lChild {
		t.lChild = p
	} else {
		f.rChild = p
	}
}
