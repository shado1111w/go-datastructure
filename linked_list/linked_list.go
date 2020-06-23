package linked_list

import "fmt"

type LNode struct {
	data int
	next *LNode
}

// 1.初始化
func (l *LNode) Init() bool {
	l.next = nil
	return true
}

// 2.取值
func (l *LNode) GetElem(i int) (e int, ok bool) {
	j := 1
	p := l.next
	for p != nil && j < i {
		p = p.next
		j++
	}

	if j == i {
		e = p.data
		ok = true
	}

	return
}

// 3.查找
func (l *LNode) LocateElem(e int) (index int) {
	j := 1
	p := l.next
	for p != nil {
		if e == p.data {
			index = j
			break
		}

		p = p.next
		j++
	}

	return
}

// 4.插入
func (l *LNode) Insert(i, e int) (ok bool) {
	j := 0
	p := l
	for p != nil && j < i-1 {
		p = p.next
		j++
	}

	if p == nil && j > i-1 {
		ok = false
		return
	}

	s := new(LNode)
	s.data = e
	s.next = p.next
	p.next = s
	ok = true

	return
}

// 5.删除
func (l *LNode) Delete(i int) (ok bool) {
	j := 0
	p := l

	for p.next != nil && (j < i-1) {
		p = p.next
		j++
	}

	if p.next == nil || j > i-1 {
		ok = false
		return
	}

	q := p.next
	p.next = q.next
	ok = true

	return
}

// 6.创建单链表（前插法）
func (l *LNode) CreatedH(n int) {
	l.next = nil
	for i := 0; i < n; i++ {
		p := new(LNode)
		fmt.Scanln(&p.data)
		p.next = l.next
		l.next = p
	}
}

// 7.创建单链表（后插法）
func (l *LNode) CreatedR(n int) {
	l.next = nil
	r := l
	for i := 0; i < n; i++ {
		p := new(LNode)
		fmt.Scanln(&p.data)
		p.next = nil
		r.next = p
		r = p
	}
}
