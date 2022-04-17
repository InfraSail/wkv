package util

import (
	//"fmt"
	"testing"
)

func TestLoadFunc(t *testing.T) {
	d := NewDict()
	d.Store("key","value")
	if v, exi := d.Load("key"); v != nil || exi != false {
		t.Errorf("expected be OK, but %s got", v)
		t.Errorf("expected be OK, but %v got", exi)

	}
}