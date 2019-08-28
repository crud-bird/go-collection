package bst

import (
	"fmt"
	"go-collection/queue"
	"go-collection/stack"
)

//BST 二叉搜索树
type BST interface {
	Search() Data
	Minimum() Data
	Maximum() Data
	Predecessor() Data
	Successor() Data
	Insert() Data
	Delete() Data
}

//Data 节点数据
type Data struct {
	Key   int
	Value interface{}
}

//DataNode 数据节点
type DataNode struct {
	r, l *DataNode
	flag bool //后序遍历辅助标记
	Data
}

//New 创建一棵二叉树
func New(key int, val ...interface{}) *DataNode {
	var v interface{}
	if len(val) > 0 {
		v = val[0]
	}
	return &DataNode{
		Data: Data{
			Key:   key,
			Value: v,
		},
	}
}

//Insert 插入
func (t *DataNode) Insert(key int, val ...interface{}) Data {
	var v interface{}
	if len(val) > 0 {
		v = val[0]
	}

	node := &DataNode{
		Data: Data{
			Key:   key,
			Value: v,
		},
	}

	for {
		if key < t.Key {
			if t.l == nil {
				t.l = node
				break
			}
			t = t.l
		} else {
			if t.r == nil {
				t.r = node
				break
			}
			t = t.r
		}
	}

	return node.Data
}

//Search 查找
func (t *DataNode) Search(key int) Data {
	pNode := t
	for pNode != nil {
		if pNode.Key == key {
			return pNode.Data
		}

		if key < pNode.Key {
			pNode = pNode.l
		} else {
			pNode = pNode.r
		}
	}

	return Data{}
}

//PreorderTraversal 前序遍历
func PreorderTraversal(root *DataNode) {
	stack := stack.New()
	pNode := root
	for pNode != nil || !stack.IsEmpty() {
		if pNode != nil {
			fmt.Printf(" %d", pNode.Key)
			stack.Push(pNode)
			pNode = pNode.l
		} else { //pNode == nil && !stack.IsEmpty() 左子树遍历完成，开始遍历右子树
			tmp, _ := stack.Pop()
			pNode = tmp.(*DataNode).r
		}
	}
}

//InorderTraversal 中序遍历
func InorderTraversal(root *DataNode) {
	stack := stack.New()
	pNode := root

	for pNode != nil || !stack.IsEmpty() {
		if pNode != nil {
			stack.Push(pNode)
			pNode = pNode.l
		} else { //pNode == nil && !stack.IsEmpty()
			tmp, _ := stack.Pop()
			pNode = tmp.(*DataNode)
			fmt.Printf(" %d", pNode.Key)
			pNode = pNode.r
		}
	}
}

//PostorderTraversal 后序遍历
func PostorderTraversal(root *DataNode) {
	stack := stack.New()
	pNode := root

	for pNode != nil || !stack.IsEmpty() {
		if pNode != nil {
			stack.Push(pNode)
			pNode = pNode.l
		} else {
			pNode = stack.Next().(*DataNode)
			if pNode.flag { //已经遍历过该节点的右子树
				stack.Pop()
				fmt.Printf(" %d", pNode.Key)
				pNode = nil
			} else {
				pNode.flag = true
				pNode = pNode.r
			}
		}
	}
}

//LevelTraversal 层次遍历
func LevelTraversal(root *DataNode) {
	q := queue.New()
	q.Enqueue(root)
	var pNode *DataNode
	for !q.IsEmpty() {
		tmp, _ := q.Dequeue()
		pNode = tmp.(*DataNode)
		fmt.Printf(" %d", pNode.Key)

		if pNode.l != nil {
			q.Enqueue(pNode.l)
		}

		if pNode.r != nil {
			q.Enqueue(pNode.r)
		}
	}
}
