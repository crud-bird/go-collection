package stack

import (
	"errors"
	"sync"
)

//Stack 栈
type Stack interface {
	Len() int
	IsEmpty()
	Push(v interface{}) interface{}
	Pop() interface{}
}

type node struct {
	pre, next *node
	v         interface{}
}

//IStack 栈
type IStack struct {
	top   *node
	lenth int
}

//CStack ...
type CStack struct {
	s   IStack
	mtx *sync.RWMutex
}

//New 创建一个栈
func New() *IStack {
	return &IStack{
		top: new(node),
	}
}

//NewConcurrent 创建一个并发的栈
func NewConcurrent() *CStack {
	return &CStack{
		s: IStack{
			top: new(node),
		},
		mtx: &sync.RWMutex{},
	}
}

//Len 长度
func (s *IStack) Len() int {
	return s.lenth
}

//Len 长度
func (cs *CStack) Len() int {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()

	return cs.s.Len()
}

//IsEmpty 判空
func (s *IStack) IsEmpty() bool {
	return s.lenth == 0
}

//IsEmpty 判空
func (cs *CStack) IsEmpty() bool {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()

	return cs.IsEmpty()
}

//Push 压栈
func (s *IStack) Push(v interface{}) interface{} {
	tmp := &node{
		v: v,
	}

	s.top.next = tmp
	tmp.pre = s.top
	s.top = tmp
	s.lenth++

	return v
}

//Push 压栈
func (cs *CStack) Push(v interface{}) interface{} {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	return cs.s.Push(v)
}

//Pop 出栈
func (s *IStack) Pop() (interface{}, error) {
	if s.lenth <= 0 {
		return nil, errors.New("pop on Empty stack")
	}

	tmp := s.top

	s.lenth--
	s.top = tmp.pre
	s.top.next = nil

	return tmp.v, nil
}

//Pop 出栈
func (cs *CStack) Pop() (interface{}, error) {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	return cs.s.Pop()
}
