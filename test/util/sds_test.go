package test

import (
	"testing"
)

/*func TestUserAskFunc(t *testing.T){
	if ans := sds.UserAsk("set a gvriu"); ans != "OK" {
		t.Errorf("expected be OK, but %s got", ans)
	}

	if ans := sds.UserAsk("get a"); ans != "gvriu" {
		t.Errorf("expected be gviru, but %s got", ans)

	}
}*/

func TestsdsNewFunc(t *testing.T) {
	if ans := sds.sdsNew("gvriu"); ans != "OK" {
		t.Errorf("expected be OK, but %s got", ans)
	}
}

/*
func main() {
	fmt.Println("hello")
	fmt.Scanln(&str)
	UserAsk(str)

	sdsNew(str)

	//s.free = fr - s.len + 1
	//sdslen(&s)
	//sdsavail(&s)

}*/
