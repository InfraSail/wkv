package util

import (
	//"fmt"
	"testing"
)

func TestNewNodeFunc(t *testing.T) {
	if ans := NewNode(0); ans.String() != "0->nil" {
		t.Errorf("expected be OK, but %s got", ans)
	}
}

func TestNewNodeWithNextFunc(t *testing.T) {
	if ans := NewNodeWithNext(2,NewNode(0)); ans.String() != "2->0->nil" {
		t.Errorf("expected be OK, but %s got", ans)
	}
}

func TestTraverseFunc(t *testing.T) {
	node1 := NewNodeWithNext(2,NewNode(0))
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	head := NewNodeWithNext(3,node5)
	if ans := Traverse(head); ans.String() != "0->2->3->55->100->-3->3->nil" {
		t.Errorf("expected be OK, but %s got", ans)
	}
}

func TestTraverse2Func(t *testing.T) {
	node1 := NewNodeWithNext(2,NewNode(0))
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	head := NewNodeWithNext(3,node5)
	if ans := Traverse2(head); ans.String() != "-3->3->55->100->2->3->0->nil" {
		t.Errorf("expected be OK, but %s got", ans)
	}
}

func TestHasCircleFunc(t *testing.T) {
	node1 := NewNodeWithNext(2,NewNode(0))
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	head := NewNodeWithNext(3,node5)
	if ans := HasCircle(head); ans != false {
		t.Errorf("expected be OK, but %v got", ans)
	}

}

func TestGetMiddleNodeFunc(t *testing.T) {
	
	node1 := NewNodeWithNext(2,NewNode(0))
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	head := NewNodeWithNext(3,node5)
	if ans := GetMiddleNode(head); ans.Value != 55 {
		t.Errorf("expected be OK, but %f got", ans.Value)
	}
}


func TestRDelNodeFunc(t *testing.T) {
	
	node1 := NewNodeWithNext(2,NewNode(0))
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	head := NewNodeWithNext(3,node5)
	ans := RDelNode(head,2).String()
	
	if  ans != "2->0->nil" {
		t.Errorf("expected be OK, but %v got", ans)
	}
}

func TestMergeFunc(t *testing.T) {
	node1 := NewNodeWithNext(2,NewNode(0))
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	head := NewNodeWithNext(3,node5)

	node6 := NewNodeWithNext(2,NewNode(0))
	node7 := NewNodeWithNext(3,node6)
	head2 := NewNodeWithNext(55, node7)

	if ans := Merge(head, head2); ans.String() != "3->-3->55->100->55->3->2->0->3->2->0->nil" {
		t.Errorf("expected be OK, but %v got", ans)
	}
}

/*
func TestNewLeastRecentlyUsedFunc(t *testing.T) {

	if ans := NewLeastRecentlyUsed(10); ans != &LeastRecentlyUsed{10}{
		t.Errorf("expected be OK, but %v got", ans)
	}
}*/
/*
func TestFindFunc(t *testing.T) {
	n :=NewNode(0)
	node1 := NewNodeWithNext(2,n)
	node2 := NewNodeWithNext(3,node1)
	node3 := NewNodeWithNext(55, node2)
	node4 := NewNodeWithNext(100, node3)
	node5 := NewNodeWithNext(-3,node4)
	NewNodeWithNext(3,node5)
	if _, cur, exist:= NewLeastRecentlyUsed(10).Find(0);  cur.Value != n.Value || exist != true {
		//fmt.Println(&pre)
		//t.Errorf("expected be OK, but %v got", pre.Value)

		t.Errorf("expected be OK, but %+v got", *cur)
		t.Errorf("expected be OK, but %v got", exist)

	}
}*/