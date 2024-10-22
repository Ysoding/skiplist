package skiplist

import "math/rand"

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string
}

type SkipList[Key Ordered, Val any] struct {
	head *node[Key, Val]

	maxLevel int
}

type node[Key Ordered, Val any] struct {
	key     Key
	val     Val
	forward []*node[Key, Val]
}

const (
	defaultMaxLevel = 10
	p               = 0.5
)

func NewSkipList[Key Ordered, Val any]() *SkipList[Key, Val] {
	return &SkipList[Key, Val]{
		head: &node[Key, Val]{
			forward: make([]*node[Key, Val], defaultMaxLevel),
		},
		maxLevel: defaultMaxLevel,
	}
}

func (sl *SkipList[Key, Val]) Search(key Key) (Val, bool) {
	current := sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}

	current = current.forward[0]
	if current != nil && current.key == key {
		return current.val, true
	}

	var zero Val
	return zero, false
}

func (sl *SkipList[Key, Val]) Insert(key Key, value Val) {
	update := make([]*node[Key, Val], sl.maxLevel)
	current := sl.head

	for i := sl.maxLevel - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	level := sl.randomLevel()
	newNode := &node[Key, Val]{
		key:     key,
		val:     value,
		forward: make([]*node[Key, Val], level),
	}

	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

func (sl *SkipList[Key, Val]) randomLevel() int {
	level := 1
	for rand.Float32() < p && level < sl.maxLevel {
		level++
	}
	return level
}

func (sl *SkipList[Key, Val]) Delete(key Key) {
	update := make([]*node[Key, Val], sl.maxLevel)
	current := sl.head

	for i := sl.maxLevel - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]
	if current != nil && current.key == key {
		for i := 0; i < sl.maxLevel; i++ {
			if update[i].forward[i] != current {
				break
			}
			update[i].forward[i] = current.forward[i]
		}
	}
}
