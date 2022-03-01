package util
import "fmt"
import "strings"
var(
	str string
	i int = 0
	fr int = 10
	bu[] byte

	choose string = ""
)
type sdshdr struct {
// 等于SDS所保存字符串的长度

	len int
 // 记录buf数组中未使用字节的数量
	free int
 // 字节数组，用于保存字符串
	buf[] byte
 }


 //TODO:makeroom
 func sdsMakeRoomFor(s sdshdr){

 }

 func set_func(key string, value string){
	sdsnew(value)

 }

 func get_func(key string){

 }


 /*func split(s *sdshdr) string{
	s.buf = []byte(str)
	for i, value:=range s.buf{
	fmt.Printf("buf%d:'%c'",i,value)
	}
	return "OK" 
 }*/

 func user_ask(ask_str string){
	ask_split := strings.Split(ask_str, " ")
	if ask_split[0] == "set" && ask_split[0] == "SET"{
		//存储键值对
		set_func(ask_split[1],ask_split[2])
		
	}else if ask_split[0] == "get" && ask_split[0] == "GET"{
		get_func(ask_split[1])
	}else{
		fmt.Println("error")
	}
 }

//func sdsnewlen(s sdshdr) float32{
	
//}

//创建一个包含给定go字符串的SDS
func sdsnew(str string) string{
	s := sdshdr{i,fr,bu}
	s.buf = []byte(str)
	for i, value:=range s.buf{
	fmt.Printf("buf%d:'%c'",i,value)
	}
	sdslen(&s)
	sdsavail(&s)

	return "OK"
}

//返回SDS已使用的字节数
func sdslen(s *sdshdr) int{
	s.len = len(s.buf)
	fmt.Printf("sdslen: %d\n", s.len)
	return s.len
}

//返回SDS未使用的字节数
func sdsavail(s *sdshdr) int{
	s.free = fr - sdslen(s) + 1
	fmt.Printf("sdsavail: %d\n", s.free)
	return s.free
}

//清空SDS保存的字符串内容
func sdsclear(s *sdshdr) string{
	s.free = fr
	s.len = i
	s.buf = bu
	return "OK"
}

func sdscmo(){}

//增长字符串
//
func sdscat()  {
	
}
//缩减字符串
//
func sdstrim() {}

func sdsrange() {}

