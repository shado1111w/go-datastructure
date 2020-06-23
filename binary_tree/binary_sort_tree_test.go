package binary_tree

import (
	"fmt"
	"testing"
)

func TestBSTree(t *testing.T) {
	tree := new(BSTree)
	CreateBST(tree)

	DeleteBST(tree, 12)

	result := SearchBST(tree, 12)
	if result != nil {
		fmt.Println(*result)
	}
}
