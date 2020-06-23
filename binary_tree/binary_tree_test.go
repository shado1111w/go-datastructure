package binary_tree

import (
	"testing"
)

func TestBitNode(t *testing.T) {
	b := new(BitNode)
	CreateBiTree(b)

	newB := new(BitNode)
	Copy(b, newB)

	t.Logf("nodeCount:%d\n", NodeCount(newB))

	t.Logf("depth:%d\n", Depth(newB))

	InOrderTraverse1(b)
	InOrderTraverse2(b)

	thrt := new(BitNode)
	InOrderThreading(thrt, b)
	InorderTraverseThr(thrt)
}
