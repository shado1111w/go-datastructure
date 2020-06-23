package binary_tree

import (
	"fmt"

	"shado1111w.com/algorithm/stack"
)

var (
	pre *BitNode
)

type BitNode struct {
	data           string
	lChild, rChild *BitNode
	lTag, rTag     int
}

// 1.中序遍历（递归）
func InOrderTraverse1(b *BitNode) {
	if b != nil {
		InOrderTraverse1(b.lChild)
		fmt.Printf("%s", b.data)
		InOrderTraverse1(b.rChild)
	}
}

// 2.中序遍历（非递归）
func InOrderTraverse2(b *BitNode) {
	stack := new(stack.IntStack)
	p := b

	for p != nil || !stack.IsEmpty() {
		if p != nil {
			stack.Push(p)
			p = p.lChild
		} else {
			q, _ := stack.Pop()
			fmt.Printf("%s", q.(*BitNode).data)
			p = q.(*BitNode).rChild
		}
	}
}

// 3.先序遍历的顺序建立二叉链表
func CreateBiTree(b *BitNode) bool {
	var ch string
	fmt.Scanln(&ch)
	if ch == "#" {
		return false
	} else {
		b.data = ch
		b.lChild = new(BitNode)
		b.rChild = new(BitNode)
		ok := CreateBiTree(b.lChild)
		if !ok {
			b.lChild = nil
		}
		ok = CreateBiTree(b.rChild)
		if !ok {
			b.rChild = nil
		}
		return true
	}
}

// 4.复制二叉树
func Copy(b *BitNode, newB *BitNode) bool {
	if b == nil {
		return false
	} else {
		newB.data = b.data
		newB.lChild = new(BitNode)
		newB.rChild = new(BitNode)
		ok := Copy(b.lChild, newB.lChild)
		if !ok {
			newB.lChild = nil
		}
		ok = Copy(b.rChild, newB.rChild)
		if !ok {
			newB.rChild = nil
		}
		return true
	}
}

// 5.计算二叉树的深度
func Depth(b *BitNode) int {
	if b == nil {
		return 0
	} else {
		m := Depth(b.lChild)
		n := Depth(b.rChild)

		if m > n {
			return m + 1
		} else {
			return n + 1
		}
	}
}

// 6.计算二叉树中结点的个数
func NodeCount(b *BitNode) int {
	if b == nil {
		return 0
	} else {
		return NodeCount(b.lChild) + NodeCount(b.rChild) + 1
	}
}

// 7.以结点p为根的子树中序线索化
func InThreading(p *BitNode) {
	if p != nil {
		InThreading(p.lChild)

		if p.lChild == nil {
			p.lTag = 1
			p.lChild = pre
		}

		if pre.rChild == nil {
			pre.rTag = 1
			pre.rChild = p
		}

		pre = p
		InThreading(p.rChild)
	}
}

// 8.带头结点的二叉树中序线索化
func InOrderThreading(thrt, b *BitNode) {
	// 头结点的右孩子指针为右线索
	thrt.rTag = 1
	// 初始化时右指针指向自己
	thrt.rChild = thrt

	// 头结点有孩子，若树非空，则其左孩子为树根；若树为空，则其左指针指向自己
	thrt.lTag = 0
	if b == nil {
		// 头结点左指针指向自己
		thrt.lChild = thrt
	} else {
		// 头结点的左孩子指向根
		thrt.lChild = b
		// pre初值指向头结点
		pre = thrt
		// 中序线性化
		InThreading(b)
		// 中序线索化结束时，pre为最右节点，pre的右线索指向头结点
		pre.rChild = thrt
		pre.rTag = 1
		// 头结点的右线索指向pre
		thrt.rChild = pre
	}
}

// 遍历中序线索二叉树
func InorderTraverseThr(b *BitNode) {
	p := b.lChild

	for p != b {
		for p.lTag == 0 && p.lChild != nil {
			p = p.lChild
		}

		fmt.Print(p.data)

		for p.rTag == 1 && p.rChild != b {
			p = p.rChild
			fmt.Print(p.data)
		}

		p = p.rChild
	}
}
