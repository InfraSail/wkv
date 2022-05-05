package util

import (
	//"fmt"
	"testing"
)



func TestLoadFunc(t *testing.T) {
	d := NewDict()
	d.Store("key","value")

	d.Store("aaa",33)
	type value struct{
		a int
		b float32
		c []int
	}
	v := value{1, 2.5, []int{1,2}}
	d.Store("bbb",v)
	if v, exi := d.Load(SdsNew("aaa")); v != 33 || exi != true {
		t.Errorf("expected be OK, but %s got", v)
		t.Errorf("expected be OK, but %v got", exi)

	}
	if v, exi := d.Load(SdsNew("key")); v != "value" || exi != true {
		t.Errorf("expected be OK, but %s got", v)
		t.Errorf("expected be OK, but %v got", exi)

	}

	if ans := d.String(); ans != "Dict(len=3, cap=4, isRehash=false)" {
		t.Errorf("expected be OK, but %s got", ans)

	}
}

func TestNextPowerFunc(t *testing.T) {
	d := NewDict()
	if ans := d.nextPower(5); ans != 8 {
		t.Errorf("expected be OK, but %v got", ans)

	}
}

func TestRehashingFunc(t *testing.T) {
	d := NewDict()
	if ans := d.rehash(5); ans != true {
		t.Errorf("expected be OK, but %v got", ans)

	}
}

func TestLoadOrStoreFunc(t *testing.T) {
	d := NewDict()
	
	d.loadOrStore(SdsNew("aaa"),"bbb")
	d.loadOrStore(SdsNew("aaa"),"bbb")
	d.loadOrStore(SdsNew("aaa"),"bbb")

	if ent, loaded := d.loadOrStore(SdsNew("aaa"),"bbb"); ent.key.GetString() != "aaa" || loaded != true {
		t.Errorf("expected be OK, but %v got", ent)
		t.Errorf("expected be OK, but %v got", loaded)
		
	}
	if ent, loaded := d.loadOrStore(SdsNew("aaa"),"bbb"); ent.key.GetString() != "aaa" || loaded != true {
		t.Errorf("expected be OK, but %v got", ent)
		t.Errorf("expected be OK, but %v got", loaded)
		
	}
}

func TestKeyIndexFunc(t *testing.T) {
	d := NewDict()
	d.Store("aaa","val")
	if idx, exi := d.KeyIndex(SdsNew("aaa")); idx != 0 || exi.value != "val" {
		t.Errorf("expected be OK, but %d got", idx)
		t.Errorf("expected be OK, but %v got", exi)
		
	}
	if idx, exi := d.KeyIndex(SdsNew("aaa")); idx != 0 || exi.value != "val" {
		t.Errorf("expected be OK, but %d got", idx)
		t.Errorf("expected be OK, but %v got", exi)
		
	}
	
}

func TestRehashForAWhileFunc(t *testing.T){
	d := NewDict()
	if ans := d.RehashForAWhile(10); ans != 0 {
		t.Errorf("expected be OK, but %d got", ans)

	}
}

func TestRehashFunc(t *testing.T){
	d := NewDict()
	d.Store("key","value")
	d.Store("aaa",33)
	d.Store("bbb",33)
	d.Store("ccc","wfq")
	d.Store("ddd","dqw")
	if ans := d.rehash(1); ans != false {
		t.Errorf("expected be OK, but %v got", ans)
	}
}

func TestIsRehashingFunc(t *testing.T){
	d := NewDict()
	d.Store("key","value")
	if ans := d.isRehashing(); ans != false {
		t.Errorf("expected be OK, but %v got", ans)
	}
}