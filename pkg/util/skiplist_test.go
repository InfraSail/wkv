package util

import (
	//"fmt"
	"testing"
)

func TestSkipListInsertFunc(t *testing.T) {
	sl := skipListNew()
		if ans := sl.skipListInsert(1, "O1"); ans.Score != float64(1) || ans.Value != "O1" {
		 t.Errorf("expected be OK, but %f got", ans.Score)
		 t.Errorf("expected be OK, but %v got", ans.Value)

		}
	   
}

func TestSkipListSearchFunc(t *testing.T) {
	sl := skipListNew()
	sl.skipListInsert(1, "O1")
		if e, ok := sl.skipListSearch(1,"O1"); e.Score != float64(1) || ok != true {
		 t.Errorf("expected be OK, but %f got", e.Score)
		 t.Errorf("expected be OK, but %v got", ok)

		}
		if e, ok := sl.skipListSearch(100,"aa"); e != nil || ok != false {
			t.Errorf("expected be OK, but %v got", e)
			t.Errorf("expected be OK, but %v got", ok)
   
		   }
}

func TestSkipListLentFunc(t *testing.T) {
	sl := skipListNew()
	
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(30, 38)
	
		if ans := sl.skipListLen(); ans != 3{
		 t.Errorf("expected be OK, but %d got", ans)

		}
	   
}

func TestSkipListDeleteFunc(t *testing.T) {
	sl := skipListNew()
	
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(30, 38)
		if ans := sl.skipListDelete(float64(20.5)); ans.Value != "O1"{
		 t.Errorf("expected be OK, but %v got", ans.Value)

		}   
}

func TestSkipListGetElementByBankFunc(t *testing.T) {
	sl := skipListNew()
	
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(30, 38)
		if ans := sl.skipListGetElementByBank(2); ans.Value != "O1"{
		 t.Errorf("expected be OK, but %v got", ans.Value)

		}   
}

func TestSkipListIsInRangeFunc(t *testing.T) {
	sl := skipListNew()
	
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(30, 38)
		if ans := sl.skipListIsInRange(0,1); ans != 1{
		 t.Errorf("expected be OK, but %d got", ans)

		}   
		if ans := sl.skipListIsInRange(-2,-1); ans != 0{
			t.Errorf("expected be OK, but %d got", ans)
   
		   } 
}

func TestSkipListFirstInRangeFunc(t *testing.T) {
	sl := skipListNew()
	sl.skipListInsert(1, "O3")
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(0, "O2")

	sl.skipListInsert(30, 38)
		if ans := sl.skipListFirstInRange(0,1); ans.Score != 0{
		 t.Errorf("expected be OK, but %f got", ans.Score)

		}   
		if ans := sl.skipListFirstInRange(-2,-1); ans != nil{
			t.Errorf("expected be OK, but %v got", ans)
   
		   } 
}

func TestSkipListLastInRangeFunc(t *testing.T) {
	sl := skipListNew()
	sl.skipListInsert(1, "O4")

	sl.skipListInsert(1, "O3")
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(0, "O0")

	sl.skipListInsert(30, 38)
		if ans := sl.skipListLastInRange(0,1); ans.Value != "O2"{
		 t.Errorf("expected be OK, but %s got", ans.Value)

		}   
		if ans := sl.skipListLastInRange(-2,-1); ans != nil{
			t.Errorf("expected be OK, but %v got", ans)
   
		   } 
}

/*
func TestSkipListDeleteRangeByScoreFunc(t *testing.T) {
	sl := skipListNew()
	
	sl.skipListInsert(20.5, "O1")	
	sl.skipListInsert(1, "O2")
	sl.skipListInsert(30, 38)
		if ans := sl.skipListDeleteRangeByScore(0,20); ans != 2{
		 t.Errorf("expected be OK, but %d got", ans)

		}   
}*/