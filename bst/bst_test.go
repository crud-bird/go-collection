package bst_test

import (
	"fmt"
	"go-collection/bst"
	"testing"
)

func TestBST(t *testing.T) {
	p := bst.New(6)
	p.Insert(5)
	p.Insert(7)
	p.Insert(2)
	p.Insert(5)
	p.Insert(8, "王二虎")
	bst.PreorderTraversal(p)
	fmt.Println()
	bst.InorderTraversal(p)
	fmt.Println()
	bst.PostorderTraversal(p)
	fmt.Println()
	bst.LevelTraversal(p)
	fmt.Println()
	fmt.Printf("result: %v", p.Search(8))
}
