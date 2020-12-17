package skiplist

import (
	//"container/list"
	//"fmt"
	"math/rand"
	"time"
)

const (
	// SkipListMaxLevel :
	SkipListMaxLevel = 32
	// SkipListP :
	SkipListP float64 = 0.25
)

// SkipLister :
type SkipLister interface {
	Set(score uint64, val interface{})
	GetByScore(score uint64) interface{}
	GetByIndex(index int) interface{}
	Len() int
	DelByScore(score uint64)
}

func randomLevel() int {

	var level int = 1

	rand.Seed(time.Now().Unix())

	for float64(rand.Int63n(0xFFFF)) < (SkipListP * float64(0xFFFF)) {
		level++
	}

	if level < SkipListMaxLevel {
		return level
	}

	return SkipListMaxLevel

}

type skipList struct {
	head, tail *skipListNode

	length int

	level int
}

type skipListNode struct {
	val interface{}

	score uint64

	forWard []*skipListNode
}

func createNode(level int, score uint64, val interface{}) *skipListNode {

	return &skipListNode{
		forWard: make([]*skipListNode, level, SkipListMaxLevel),
		score:   score,
		val:     val,
	}
}

// Create :
func Create() SkipLister {

	return &skipList{
		head:   createNode(SkipListMaxLevel, 0, nil),
		tail:   nil,
		length: 0,
		level:  0,
	}
}

// Set
func (sk *skipList) Set(score uint64, val interface{}) {

	sk.set(score, val)

	return
}

func (sk *skipList) set(score uint64, val interface{}) {

	update := make([]*skipListNode, SkipListMaxLevel, SkipListMaxLevel)

	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x.forWard[i] != nil && x.forWard[i].score < score {
			x = x.forWard[i]
		}
		update[i] = x
	}

	if x.forWard[0] != nil && x.forWard[0].score == score {
		x.forWard[0].val = val
		return
	}

	level := randomLevel()

	if level > sk.level {

		for i := sk.level; i < level; i++ {
			update[i] = sk.head
		}
		sk.level = level
	}

	x = createNode(level, score, val)

	for i := 0; i < level; i++ {
		x.forWard[i] = update[i].forWard[i]
		update[i].forWard[i] = x
	}

	sk.length++

	return
}

// DelByScore
func (sk *skipList) DelByScore(score uint64) {

	sk.delByScore(score)

	return
}

func (sk *skipList) delByScore(score uint64) {

	update := make([]*skipListNode, SkipListMaxLevel, SkipListMaxLevel)

	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x.forWard[i] != nil && x.forWard[i].score < score {
			x = x.forWard[i]
		}

		update[i] = x
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

	sk.length--

	return
}

// GetByScore
func (sk *skipList) GetByScore(score uint64) interface{} {

	x := sk.head

	for i := sk.level - 1; i >= 0; i-- {
		for x.forWard[i] != nil && x.forWard[i].score < score {
			x = x.forWard[i]
		}
	}

	x = x.forWard[0]

	if x != nil && x.score == score {
		return x.val
	}

	return nil
}

// GetByIndex
func (sk *skipList) GetByIndex(index int) interface{} {

	idx := 0
	for x := sk.head.forWard[0]; x != nil; x = x.forWard[0] {
		if idx == index {
			return x.val
		}
		idx++
	}

	return nil
}

// Len
func (sk *skipList) Len() int {

	return sk.length
}
