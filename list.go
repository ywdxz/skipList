package skipList

import (
	//"container/list"
	//"fmt"
	"math/rand"
	"time"
)

const (
	SKIPLIST_MAXLEVEL         = 32
	SKIPLIST_P        float64 = 0.25
)

func randomLevel() int {

	var level int = 1

	rand.Seed(time.Now().Unix())

	for float64(rand.Int63n(0xFFFF)) < (SKIPLIST_P * float64(0xFFFF)) {
		level++
	}

	if level < SKIPLIST_MAXLEVEL {
		return level
	}

	return SKIPLIST_MAXLEVEL

}

type skipList struct {
	head, tail *skipListNode

	length uint64

	level int
}

type skipListNode struct {
	val interface{}

	score uint64

	backWard *skipListNode

	forWard []*skipListNode
}

func CreateNode(level int, score uint64, val interface{}) *skipListNode {

	return &skipListNode{
		forWard:  make([]*skipListNode, level),
		score:    score,
		val:      val,
		backWard: nil,
	}
}

func Create() *skipList {

	return &skipList{
		head:   CreateNode(SKIPLIST_MAXLEVEL, 0, nil),
		tail:   nil,
		length: 0,
		level:  0,
	}
}

func (sk *skipList) Insert(score uint64, val interface{}) {

	update := make([]*skipListNode, SKIPLIST_MAXLEVEL)

	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x.forWard[i] != nil && x.forWard[i].score < score {
			x = x.forWard[i]
		}
		update[i] = x
	}

	level := randomLevel()

	if level > sk.level {

		for i := sk.level; i < level; i++ {
			update[i] = sk.head
		}
		sk.level = level
	}

	x = CreateNode(level, score, val)

	for i := 0; i < level; i++ {
		x.forWard[i] = update[i].forWard[i]
		update[i].forWard[i] = x
	}

	x.backWard = update[0]
	if update[0] == sk.head {
		x.backWard = nil
	}

	if x.forWard[0] != nil {
		x.forWard[0].backWard = x
	} else {
		sk.tail = x
	}

	sk.length++

	return
}

func (sk *skipList) Delete(score uint64) {

	update := make([]*skipListNode, SKIPLIST_MAXLEVEL)

	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x != nil && x.score < score {
			x = x.forWard[i]
		}
		update = append(update, x)
	}

	x = x.forWard[0]

	if x != nil && score == x.score {
		sk.deleteNode(x, update)
	}

	return
}

func (sk *skipList) deleteNode(x *skipListNode, update []*skipListNode) {

	for i := 0; i < sk.level; i++ {
		if update[i].forWard[i] == x {
			update[i].forWard[i] = x.forWard[i]
		}
	}

	for sk.level > 1 && sk.head.forWard[sk.level-1] == nil {
		sk.level--
	}

	sk.level--

	return
}

func (sk *skipList) GetFirstInRange(score uint64) (ret *skipListNode) {

	ret = nil
	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x != nil && x.score < score {
			x = x.forWard[i]
		}
	}

	if x != nil && x.score == score {
		return x
	}

	return
}

func (sk *skipList) GetLastInRange(score uint64) (ret *skipListNode) {

	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x != nil && x.score > score {
			x = x.forWard[i]
		}
	}

	if x != nil && x.score == score {
		return x
	}

	return
}
