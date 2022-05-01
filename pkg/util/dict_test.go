package util

import (
	//"fmt"
	"testing"
)



func TestLoadFunc(t *testing.T) {
	d := NewDict()
	d.Store("key","value")
	d.Store("aaa","bbb")
	d.Store("aaa","bbb")
	d.Store("aaa","bbb")
	d.Store("aaa","bbb")
	d.Store("aaa",33)

	if v, exi := d.Load(sdsNew("aaa")); v != nil || exi != false {
		t.Errorf("expected be OK, but %s got", v)
		t.Errorf("expected be OK, but %v got", exi)

	}
	if ans := d.String(); ans != "Dict(len=6, cap=8, isRehash=true)" {
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
	
	d.loadOrStore(sdsNew("aaa"),"bbb")
	d.loadOrStore(sdsNew("aaa"),"bbb")
	d.loadOrStore(sdsNew("aaa"),"bbb")

	if ent, loaded := d.loadOrStore(sdsNew("aaa"),"bbb"); ent != nil || loaded != false {
		t.Errorf("expected be OK, but %v got", ent)
		t.Errorf("expected be OK, but %v got", loaded)
		
	}
	if ent, loaded := d.loadOrStore(sdsNew("aaa"),"bbb"); ent != nil || loaded != false {
		t.Errorf("expected be OK, but %v got", ent)
		t.Errorf("expected be OK, but %v got", loaded)
		
	}
}

func TestKeyIndexFunc(t *testing.T) {
	d := NewDict()
	d.Store("aaa","val")
	if idx, exi := d.keyIndex(sdsNew("aaa")); idx != 0 || exi != nil {
		t.Errorf("expected be OK, but %d got", idx)
		t.Errorf("expected be OK, but %v got", exi)
		
	}
	if idx, exi := d.keyIndex(sdsNew("aaa")); idx != 0 || exi != nil {
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