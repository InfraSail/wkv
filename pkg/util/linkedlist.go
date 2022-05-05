package util

import (
	"fmt"
	"sync"
)

type Node struct {
	Value float64
	Next  *Node
}

// 以当前节点开头，以字符串格式输出这个链表
func (n *Node) String() string {
	if n == nil {
		return ""
	}

	ret := ""
	current := n
	for {
		if current == nil {
			ret += "nil"
			return ret
		}

		ret += fmt.Sprintf("%v->", current.Value)
		current = current.Next
	}
}

// 新建节点
func NewNode(value float64) *Node {
	return &Node{
		Value: value,
	}

}

// 新建一个有指定next节点的节点
func NewNodeWithNext(value float64, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

// 遍历链表
func Traverse(head *Node) *Node {
	var pre, current *Node = nil, head
	for {
		if current == nil {
			return pre
		}
		current.Next, pre, current = pre, current, current.Next
	}
}

// 遍历链表2
func Traverse2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head.Next
	cur := NewNodeWithNext(0, head)
	for {
		if cur.Next == nil || cur.Next.Next == nil {
			return pre
		}

		a := cur.Next
		b := cur.Next.Next
		cur.Next, a.Next, b.Next = b, b.Next, a
		cur = a
	}
}

// 检测链表是否有环（快慢指针）
func HasCircle(head *Node) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for {
		if fast.Next == nil || fast.Next.Next == nil || slow.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
}

// 获取链表的中间节点
func GetMiddleNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil && slow.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			continue
		}

		return slow
	}
}

// 删除倒数第n个结点，如果倒数n个结点为头结点则不可删除
func RDelNode(head *Node, n int) *Node {
	slow, fast := head, head

	// 1 2 3 4 5
	for ; n > 0; n-- {
		if fast.Next == nil {
			break
		}
		fast = fast.Next
	}

	// 如果这里 n == 1 则表示删除头结点，目前可以考虑不删除头结点
	if n > 0 {
		return nil
	}

	for {
		if fast.Next != nil && slow.Next != nil {
			fast = fast.Next
			slow = slow.Next
			continue
		}
		break
	}

	if slow != nil && slow.Next != nil {
		ret := slow.Next
		slow.Next = slow.Next.Next
		return ret
	}
	return nil
}

// 链表合并,head1在前head2在后
func Merge(head1, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	head, insert := head1, head2
	if head2.Value < head1.Value {

		head, insert = head2, head1

	}

	cur := head
	for {
		if insert == nil {
			return head
		}

		if cur.Next == nil {
			cur.Next = insert
			return head
		}

		for {
			if insert.Value > cur.Value && cur.Next != nil && insert.Value < cur.Next.Value {
				tmp := insert.Next
				insert.Next = cur.Next
				cur.Next = insert

				cur = cur.Next
				insert = tmp
				break
			}

			if cur.Next == nil {
				break
			}
			cur = cur.Next
		}
	}
}

// 当前链表可使用的最小容量
type LeastRecentlyUsed struct {
	Capacity uint64       // 容量
	Number   uint64       // 当前链表数量
	Head     *Node        // 链表头节点
	mu       sync.RWMutex // 读写锁
}

// 新建一个LeastRecentlyUsed结构体，capacity为链表的容量
func NewLeastRecentlyUsed(capacity uint64) *LeastRecentlyUsed {
	return &LeastRecentlyUsed{
		Capacity: capacity,
	}
}

// 返回要查找的结点的前一个节点跟自己的节点，如果pre为空则要查找的结点就是头结点，并判断是否存在
func (l *LeastRecentlyUsed) Find(value interface{}) (pre, cur *Node, exist bool) {
	l.mu.RLock()
	// 方法结束之后释放读写锁
	defer l.mu.RUnlock()

	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		cur = l.Head
		exist = true
		return
	}
	pre = l.Head
	for {
		if pre.Next == nil {
			return
		}

		if pre.Next.Value == value {
			cur = pre.Next
			exist = true
			return
		}

		pre = pre.Next
	}
}

// 值为value的节点插入链表
func (l *LeastRecentlyUsed) Use(value float64) {
	pre, cur, ok := l.Find(value)
	l.mu.Lock()

	// 如果当前节点已经存在则将当前节点放在第一个节点
	if ok && pre != nil {
		pre.Next = cur.Next
		cur.Next = l.Head
		l.Head = cur
		l.mu.Unlock()
		return
	}

	// 如果当前的缓存容量已经满了，则将当前节点覆盖最后一个节点
	if l.Capacity <= l.Number {
		if l.Capacity <= 1 {
			l.Head = NewNode(value)
			l.mu.Unlock()
			return
		}
		pre := l.Head
		for {
			if pre.Next.Next != nil {
				pre = pre.Next
				continue
			}

			pre.Next = NewNode(value)
			l.mu.Unlock()
			return
		}
	} else {
		l.Number++
		newNode := NewNodeWithNext(value, l.Head)
		l.Head = newNode
		l.mu.Unlock()
		return
	}
}
