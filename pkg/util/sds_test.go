package util

import (
	"testing"
)


func TestSdsNewFunc(t *testing.T) {
	 if ans := sdsNew("feefds"); ans.GetString() != "feefds" {
	  t.Errorf("expected be OK, but %s got", ans.GetString())
	 }
	}


func TestSdsClearFunc(t *testing.T) {
	if ans := sdsClear(&sdshdr{i,fr, bu}); ans != "OK" {
		t.Errorf("expected be OK, but %s got", ans)
	}
}
func TestSdsCmpFunc(t *testing.T) {
	if ans := sdsCmp(&sdshdr{i,fr, bu},&sdshdr{i,fr, bu}); ans != 0 {
		t.Errorf("expected be OK, but %d got", ans)
	}
}
func TestSdsCatFunc(t *testing.T) {
	if ans := sdsCat(&sdshdr{4,4, []byte{'x','e','r','e',0,0,0,0}}, "hdi"); ans != "OK" {
		t.Errorf("expected be OK, but %s got", ans)
	}
	if ans := sdsCat(&sdshdr{4,4, []byte{'x','e','r','e',0,0,0,0}}, "hdiwdwsxcew"); ans != "OK" {
		t.Errorf("expected be OK, but %s got", ans)
	}
	if ans := sdsCat(&sdshdr{4, 4, []byte{'x', 'e', 'r', 'e', 0, 0, 0, 0}}, func() (str3 string) {
		 for i := 0; i < 512*1024; i++ {
		 str3 += "njxo"
		 }
		 return
		 }()); ans != "OK" {
		 t.Errorf("expected be OK, but %s got", ans)
		 }
}





func TestSdsGrowZeroFunc(t *testing.T) {
	if ans := sdsGrowZero(&sdshdr{4,4, []byte{'x','e','r','e',0,0,0,0}}, 6); ans != "xere  " {
		t.Errorf("expected be OK, but %s got", ans)
	}
}

func TestSdsRangeFunc(t *testing.T){
	if err := sdsRange(&sdshdr{6,6, []byte{'x','e','r','e','r','e',0,0,0,0,0,0}},1,4); err != nil{
		t.Errorf("expected be OK, but %v got", err)
	
	}
}


func TestSdsTrimFunc(t *testing.T) {
	if ans := sdsTrim(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}},"xdwf"); ans != "eee" {
	 t.Errorf("expected be ddd, but %s got", ans)
	}
	if ans := sdsTrim(&sdshdr{7,7, []byte{'a','a','d','e','d','e','e',0,0,0,0,0,0,0}},"ed"); ans != "aa" {
		t.Errorf("expected be aa, but %s got", ans)
	   }
	   if ans := sdsTrim(&sdshdr{7,7, []byte{'a','x','a','e','d','d','e',0,0,0,0,0,0,0}},"axxc"); ans != "edde" {
		t.Errorf("expected be dedee, but %s got", ans)
	   }
   }


func TestSdsSeplitLenFunc(t *testing.T) {
	   /*ans1 := [7]string{"xededed"}
	   ans2 := []string{"xe","e","e"}
	if ans := sdsSeplitLen(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}}," "); ans != ans1 {
		t.Errorf("expected be , but %v got", ans)
	   }

	if ans := sdsSeplitLen(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}},"d"); ans != ans2 {
		t.Errorf("expected be , but %v got", ans)
	   }*/
	   sdsSeplitLen(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}}," ")
   }
/*
func TestSdsFreeFunc(t *testing.T) {

 if ans := sdsFree(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}}); ans != &sdshdr{1,1,'a'} {
	 t.Errorf("expected be nil, but %v got", ans)

	}
}*/

func TestSdsDupFunc(t *testing.T) {

	if ans := sdsDup(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}}); ans.GetString() != "xededed" {
		t.Errorf("expected be xdeded, but %v got", ans)
   
	   }
   }



func TestSdsLenFunc(t *testing.T){
	if ans := sdsLen(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}}); ans != 7 {
		t.Errorf("expected be xdeded, but %d got", ans)
   
	   }
}

func TestSdsAvailFunc(t *testing.T){
	if ans := sdsAvail(&sdshdr{7,7, []byte{'x','e','d','e','d','e','d',0,0,0,0,0,0,0}}); ans != 7 {
		t.Errorf("expected be xdeded, but %d got", ans)
   
	   }
}

func TestSdsEmptyFunc(t *testing.T){
	sdsEmpty() 
}
