package util

import (
	// "fmt"
	"math/rand"
)

const (
	maxLevel = 64 // 足以容纳 2^64 个元素
	p        = 0.25 
)

// Element是一个跳表的元素
type Element struct {
    Score   float64
    Value   interface{}
    forward []*Element
}

// 新建元素
func newElement(score float64, value interface{}, level int) *Element {
    return &Element{
        Score:   score,
        Value:   value,
        forward: make([]*Element, level),
    }
}

// SkipList代表一个调表
// 来自SkipList的零值是一个准备使用的空skiplist。
type SkipList struct {
    header *Element // 头节点是一个虚拟元素
    len    int      // 目前的跳表长度，不包括头节点
    level  int      // 目前的跳表高度，不包括头节点
}

// 目前的跳表长度，不包括头节点
func (sl *SkipList) skipListLen() int {
	return sl.len
}

// 目前的跳表高度，不包括头节点
func (sl *SkipList) skipListLevel() int {
	return sl.level
}

// New返回一个新的调表指针
func skipListNew() *SkipList {
    return &SkipList{
        header: &Element{forward: make([]*Element, maxLevel)},
    }
}

// 随机生跳表高度
func skipListRandomLevel() int {
    level := 1
    for rand.Float32() < p && level < maxLevel {
        level++
    }
    return level
}

// 返回跳表的第一个值，可能是nil
func (sl *SkipList) skipListFront() *Element {
    return sl.header.forward[0]
}

// 返回e节点的后一个值
func (e *Element) skipListNext() *Element {
    if e != nil {
        return e.forward[0]
    }
    return nil
}

// 返回跳表在给定排位上的节点
func (sl *SkipList) skipListGetElementByBank(bank int) *Element{
	i := 0
	for  e := sl.skipListFront();e != nil ; e = e.skipListNext() { 
		i ++
		if i == bank - 1{
			return e.forward[0]
		}
	}
	return nil
}

// 给定一个float64类型分值范围，如果跳表中有至少一个节点分值在这个范围内
// 那么返回 1 否则返回 0
func (sl *SkipList) skipListIsInRange(left float64, right float64) int{
	x := sl.header
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < left {
            x = x.forward[i]
        }
    }
    x = x.forward[0]
    if x != nil && x.Score >= left && x.Score <= right {
        return 1
    }
    return 0
}

// 给定一个float64类型分值范围，返回跳表内第一个符合此范围的节点
func (sl *SkipList) skipListFirstInRange(left float64, right float64) *Element{
	x := sl.header
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < left {
            x = x.forward[i]
        }
    }
    x = x.forward[0]
    if x != nil && x.Score >= left && x.Score <= right {
        return x
    }
    return nil
}

// 给定一个float64类型分值范围，返回跳表内最后一个符合此范围的节点
func (sl *SkipList) skipListLastInRange(left float64, right float64) *Element{
	x := sl.header
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < right {
            x = x.forward[i]
        }
    }
	if !(x != nil && x.Score >= left && x.Score <= right){
		return nil
	}
    for x != nil && x.Score >= left && x.Score < right {
		x = x.forward[0]

    }
	return x
	
}

// 搜索skiplist，找出具有给定分数的元素。
// 如果给定的分数存在，返回（*Element, true），否则返回（nil, false）。
func (sl *SkipList) skipListSearch(score float64, value interface{}) (element *Element, ok bool) {
    x := sl.header
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < score {
            x = x.forward[i]
        }
    }
    x = x.forward[0]
    if x != nil && x.Score == score && x.Value == value {
        return x, true
    }
    return nil, false
}

// 在skiplist中插入（分值，值）对，并返回元素的指针。
func (sl *SkipList) skipListInsert(score float64, value interface{}) *Element {
    update := make([]*Element, maxLevel)
    x := sl.header
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < score {
            x = x.forward[i]
        }
        update[i] = x
    }
    x = x.forward[0]

    // 已经出现的分值，用新值替换，然后返回
    if x != nil && x.Score == score {
        x.Value = value
        return x
    }

    level := skipListRandomLevel()
    if level > sl.level {
        level = sl.level + 1
        update[sl.level] = sl.header
        sl.level = level
    }
    e := newElement(score, value, level)
    for i := 0; i < level; i++ {
        e.forward[i] = update[i].forward[i]
        update[i].forward[i] = e
    }
    sl.len++
    return e
}

// 删除给定值的节点，并返回给定值，如果节点不存在，返回nil
func (sl *SkipList) skipListDelete(score float64) *Element {
    update := make([]*Element, maxLevel)
    x := sl.header
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < score {
            x = x.forward[i]
        }
        update[i] = x
    }
    x = x.forward[0]

    if x != nil && x.Score == score {
        for i := 0; i < sl.level; i++ {
            if update[i].forward[i].Score != x.Score {
                return nil
            }
            update[i].forward[i] = x.forward[i]
        }
        sl.len--
    }
    return x
}

// 给定一个分值范围，删除跳跃表中所有在这个范围内的节点

func (sl *SkipList) skipListDeleteRangeByScore(left float64, right float64) int{
	num := 0
	//const s *Element = sl.skipListFirstInRange(left,right) 
	
	
		for{
            s := sl.skipListFirstInRange(left,right)
            if s == nil{
                break
            }
            sl.skipListDelete(s.Score)
            
	    	num ++

        }

		
	
	return num
}

