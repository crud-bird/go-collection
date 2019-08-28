package queue

import (
	"errors"
	"sync"
)

//Queue 队列
type Queue interface {
	IsEmpty() bool
	IsFull() bool
	Len() int
	Next() interface{}
	Enqueue(interface{}) interface{}
	Dequeue() (interface{}, error)
}

//IQueue 自动扩容队列
type IQueue struct {
	queue                   []interface{}
	head, tail, lenth, capa int
}

//CQueue 并发情况下使用
type CQueue struct {
	q   IQueue
	mtx *sync.RWMutex
}

//New 创建一个队列，capa为初始容量，默认16
func New(capa ...int) *IQueue {
	n := 16
	if len(capa) > 0 {
		n = capa[0]
	}
	return &IQueue{
		queue: make([]interface{}, n, n),
		capa:  n,
	}
}

//NewConcurrent 创建一个并发的队列，capa为初始容量，默认16
func NewConcurrent(capa ...int) *CQueue {
	n := 16
	if len(capa) > 0 {
		n = capa[0]
	}
	return &CQueue{
		q: IQueue{
			queue: make([]interface{}, n, n),
			capa:  n,
		},
		mtx: &sync.RWMutex{},
	}
}

//IsEmpty 判空
func (q *IQueue) IsEmpty() bool {
	return q.head == q.tail
}

//IsEmpty 判空
func (cq *CQueue) IsEmpty() bool {
	cq.mtx.RLock()
	defer cq.mtx.RUnlock()

	return cq.q.IsEmpty()
}

//IsFull 是否已满
func (q *IQueue) IsFull() bool {
	return (q.tail+1)%q.capa == q.head
}

//IsFull 是否已满
func (cq *CQueue) IsFull() bool {
	cq.mtx.RLock()
	defer cq.mtx.RUnlock()

	return cq.q.IsFull()
}

//Len 队列长度
func (q *IQueue) Len() int {
	return q.lenth
}

//Len 队列长度
func (cq *CQueue) Len() int {
	cq.mtx.RLock()
	defer cq.mtx.RUnlock()

	return cq.q.Len()
}

//Next 查询下一个元素
func (q *IQueue) Next() interface{} {
	return q.queue[q.head]
}

//Next 查询下一个元素
func (cq *CQueue) Next() interface{} {
	cq.mtx.RLock()
	defer cq.mtx.RUnlock()

	return cq.q.Next()
}

//Enqueue 入队
func (q *IQueue) Enqueue(v interface{}) interface{} {
	if q.IsFull() {
		q.extend()
	}

	q.queue[q.tail] = v
	q.tail = (q.tail + 1) % q.capa
	q.lenth++

	return v
}

//Enqueue 入队
func (cq *CQueue) Enqueue(v interface{}) interface{} {
	cq.mtx.Lock()
	defer cq.mtx.Unlock()

	return cq.q.Enqueue(v)
}

//Dequeue 出队
func (q *IQueue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("dequeue on empty queue")
	}

	res := q.queue[q.head]
	q.head = (q.head + 1) % q.capa
	q.lenth--
	return res, nil
}

//Dequeue 出队
func (cq *CQueue) Dequeue() (interface{}, error) {
	cq.mtx.Lock()
	defer cq.mtx.Unlock()

	return cq.q.Dequeue()
}

//extend 扩容
func (q *IQueue) extend() {
	newQ := make([]interface{}, q.capa*2, q.capa*2)
	for i := 0; i < q.capa-1; i++ {
		idx := q.head + i
		newQ[i] = q.queue[idx%q.capa]
	}

	q.queue = newQ
	q.head = 0
	q.tail = q.capa - 1
	q.capa = len(newQ)
}
