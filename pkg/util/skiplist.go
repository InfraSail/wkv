package util
/*
import "math/rand"

const (
	MaxLevel = 64 // 足以容纳 2^64 个元素
	P        = 0.25 
)

// 跳跃表节点的结构体
type SkipListNode struct {
	elem     string // 成员对象
	score    float64 // 分值，节点按分值的从小到大排列
	backward *SkipListNode // 后退指针
	level    []skipLevel // 标记此节点的层级
}

// 跳跃表层级的结构体
type skipLevel struct {
	// forward 每层都要有指向下一个节点的指针
	forward *SkipListNode
	// span 间隔定义为：从当前节点到 forward 指向的下个节点之间间隔的节点数
	span int
}

// 跳跃表总体的结构体
type Skiplist struct {
	header, tail *SkipListNode // header指向第一个节点， tail指向末尾节点
	level        int // 记录跳表的实际高度
	length       int // 记录跳表的长度（不含头节点）
}

// 比较节点的分值
func (node *SkipListNode) skipListCompare(other *SkipListNode) int {
	if node.score < other.score || (node.score == other.score && node.elem < other.elem) {
		return -1
	} else if node.score > other.score || (node.score == other.score && node.elem > other.elem) {
		return 1
	} else {
		return 0
	}
}

func (node *SkipListNode) Lt(other *SkipListNode) bool {
	return node.skipListCompare(other) < 0
}

func (node *SkipListNode) Lte(other *SkipListNode) bool {
	return node.skipListCompare(other) <= 0
}

func (node *SkipListNode) Gt(other *SkipListNode) bool {
	return node.skipListCompare(other) > 0
}

func (node *SkipListNode) Eq(other *SkipListNode) bool {
	return node.skipListCompare(other) == 0
}

// 插入节点
func (sl *Skiplist) skipListInsert(score float64, elem string) *SkipListNode {
	var (
		// update 用于记录每层待更新的节点
		update [MaxLevel]*SkipListNode
		// rank 用来记录每层经过的节点记录（可以看成到头节点的距离）
		rank [MaxLevel]int
		// 构建一个新节点，用于下面的大小判断，其 level 在后面设置
		node = &SkipListNode{score: score, elem: elem}
	)
	cur := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		if cur == sl.header {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for cur.level[i].forward != nil && cur.level[i].forward.Lt(node) {
			rank[i] += cur.level[i].span
			// 同层继续往后查找
			cur = cur.level[i].forward
		}
		update[i] = cur
	}
	// 调整跳表高度
	level := sl.randomLevel()
	if level > sl.level {
		// 初始化每层
		for i := level - 1; i >= sl.level; i-- {
			rank[i] = 0
			update[i] = sl.header
			update[i].level[i].span = sl.length
		}
		sl.level = level
	}
	// 更新节点 level，并插入新节点
	// TODO: node.SetLevel(level)
	for i := 0; i < level; i++ {
		// 更新每层的节点指向
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node
		// 更新 span 信息
		node.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}
	// 针对新增节点 level < sl.level 的情况，需要更新上面没有扫到的层 span
	for i := level; i < sl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] != sl.header {
		// update[0] 就是和新增节点相邻的前一个节点
		node.backward = update[0]
	}
	// 如果新增节点是最后一个，则需要更新 tail 指针
	if node.level[0].forward == nil {
		sl.tail = node
	} else {
		// 中间节点，需要更新后一个节点的回退指针
		node.level[0].forward.backward = node
	}
	sl.length++
	return node
}

// 
func (sl *Skiplist) randomLevel() int {
	level := 1
	for rand.Float64() < P && level < MaxLevel {
		level++
	}
	return level
}

// Delete 用于删除跳表中指定的节点。
func (sl *Skiplist) Delete(score float64, elem string) *SkipListNode {
	// 第一步，找到需要删除节点
	var (
		update     [MaxLevel]*SkipListNode
		targetNode = &SkipListNode{elem: elem, score: score}
	)
	cur := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for cur.level[i].forward != nil && cur.level[i].forward.Lt(targetNode) {
			cur = cur.level[i].forward
		}
		update[i] = cur
	}

	nodeToBeDeleted := update[0].level[0].forward
	if nodeToBeDeleted == nil || !nodeToBeDeleted.Eq(targetNode) {
		return nil
	}
	sl.deleteNode(update, nodeToBeDeleted)
	return nodeToBeDeleted
}

// 删除节点
func (sl *Skiplist) deleteNode(update [64]*SkipListNode, nodeToBeDeleted *SkipListNode) {

	// 调整每层待更新节点，修改 forward 指向
	for i := 0; i < sl.level; i++ {
		if update[i].level[i].forward == nodeToBeDeleted {
			update[i].level[i].forward = nodeToBeDeleted.level[i].forward
			update[i].level[i].span += nodeToBeDeleted.level[i].span - 1
		} else {
			update[i].level[i].span--
		}
	}
	// 调整回退指针：
	// 1. 如果被删除的节点是最后一个节点，需要更新 sl.tail
	// 2. 如果被删除的节点位于中间，则直接更新后一个节点 backward 即可
	if sl.tail == nodeToBeDeleted {
		sl.tail = nodeToBeDeleted.backward
	} else {
		nodeToBeDeleted.level[0].forward.backward = nodeToBeDeleted.backward
	}
	// 调整层数
	for sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}
	// 减少节点计数
	sl.length--
	nodeToBeDeleted.backward = nil
	nodeToBeDeleted.level[0].forward = nil
}

// 更新节点
func (sl *Skiplist) UpdateScore(curScore float64, elem string, newScore float64) *SkipListNode {
	var (
		update     [MaxLevel]*SkipListNode
		targetNode = &SkipListNode{elem: elem, score: curScore}
	)
	cur := sl.header
	// 第一步，找到符合条件的目标节点
	for i := sl.level - 1; i >= 0; i-- {
		for cur.level[i].forward != nil && cur.level[i].forward.Lt(targetNode) {
			cur = cur.level[i].forward
		}
		update[i] = cur
	}
	node := cur.level[0].forward
	if node == nil || !node.Eq(targetNode) {
		return nil
	}
	if sl.canUpdateScoreFor(node, newScore) {
		node.score = newScore
		return node
	} else {
		// 需要删除旧节点，增加新节点
		sl.deleteNode(update, node)
		return sl.Insert(newScore, node.elem)
	}
}

func (sl *Skiplist) canUpdateScoreFor(node *SkipListNode, newScore float64) bool {
	if (node.backward == nil || node.backward.score < newScore) &&
		(node.level[0].forward == nil || node.level[0].forward.score > newScore) {
		return true
	}

	return false
}
*/