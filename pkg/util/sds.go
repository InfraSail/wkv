package pkg

import (
	"errors"
	"fmt"
	"strings"
)

const SDS_MAX_PREALLOC int = 1024 * 1024

var (
	str string
	i   int = 0
	fr  int = 0

	bu         []byte
	sdslen     int
	choose     string = ""
	cat_or_not int    = 0
)

type sdshdr struct {
	// 等于SDS所保存字符串的长度

	len int
	// 记录buf数组中未使用字节的数量
	free int
	// 字节数组，用于保存字符串
	buf []byte
}

//TODO:makeroom
func sdsMakeRoomFor(s *sdshdr) {
	if s.len > SDS_MAX_PREALLOC {
		fr *= 2
	} else {
		fr += SDS_MAX_PREALLOC
	}
}

func SetFunc(key string, value string) {

	sdsNew(value)

}

func GetFunc(key string) {

}

/*func split(s *sdshdr) string{
	s.buf = []byte(str)
	for i, value:=range s.buf{
	fmt.Printf("buf%d:'%c'",i,value)
	}
	return "OK"
 }*/

func UserAsk(ask_str string) {
	ask_split := strings.Split(ask_str, " ")
	if ask_split[0] == "set" && ask_split[0] == "SET" {
		//存储键值对
		SetFunc(ask_split[1], ask_split[2])

	} else if ask_split[0] == "get" && ask_split[0] == "GET" {
		GetFunc(ask_split[1])
	} else {
		fmt.Println("error")
	}
}

//func sdsnewlen(s sdshdr) float32{

//}

//创建一个包含给定go字符串的SDS
func sdsNew(str string) string {
	s := sdshdr{i, fr, bu}
	s.buf = []byte(str)
	for i, value := range s.buf {
		fmt.Printf("buf%d:'%c'", i, value)
	}

	sdslen = sdsLen(&s)
	/*
		if(sdslen <= fr){
			fr *= 2
			s.free = fr - sdslen + 1
		}else{
			fr +=
		}*/
	sdsAvail(&s)

	return "OK"
}

//创建一个不包含任何内容的空 SDS
func sdsEmpty() sdshdr {
	s := sdshdr{0, 0, bu}
	return s
}

//返回SDS已使用的字节数
func sdsLen(s *sdshdr) int {
	s.len = len(s.buf)
	fmt.Printf("sdslen: %d\n", s.len)
	return s.len
}

//返回SDS未使用的字节数
func sdsAvail(s *sdshdr) int {
	if cat_or_not == 0 {
		s.free = 0
	} else if cat_or_not == 1 {
		s.free, fr = s.len, s.len*2

	} else if cat_or_not > 1 {
		if s.len > fr {
			sdsMakeRoomFor(s)
		}
		s.free = fr - sdsLen(s) + 1

	}
	//s.free = fr - sdsLen(s) + 1
	fmt.Printf("sdsavail: %d\n", s.free)
	return s.free
}

//清空SDS保存的字符串内容
func sdsClear(s *sdshdr) string {
	s.free = fr
	s.len = i
	s.buf = bu
	return "OK"
}

//对比两个 SDS 字符串是否相同,-1不同,0相同
func sdsCmp(s1 *sdshdr, s2 *sdshdr) int {

	return strings.Compare(string(s1.buf), string(s2.buf))
}

//增长字符串
func sdsCat(s *sdshdr, cat_str string) string {
	//var cat_buf[] byte
	cat_or_not++

	cat_buf := []byte(cat_str)
	for i, value := range cat_buf {
		s.buf[i+s.len] = value
	}
	sdsAvail(s)
	return "OK"
}

//缩减字符串
//

func sdsTrim(s *sdshdr, trim_str string) string {
	trim_buf := []byte(trim_str)
	for _, trim_value := range trim_buf {
		for _, s_value := range s.buf {
			if trim_value == s_value {
				//TODO:移除特定的字符
			}
		}
	}
	return string(s.buf)
}

//保留 SDS 给定区间内的数据， 不在区间内的数据会被覆盖或清除。
func sdsRange(s *sdshdr, left int, right int) (str string, err error) {
	if right < left || left < 0 || right > s.len {
		err = errors.New("ERROR!")
		str = ""
		return
	}
	s.buf = s.buf[left:]
	s.buf = s.buf[:right]
	//s.len -= (right - left)
	s.free = fr - sdsLen(s) + 1
	return
}

//
func sdsGrowZero(s *sdshdr, self_defined_len int) string {
	cat_or_not++
	for i := 0; i < self_defined_len; i++ {
		s.buf[s.len+i] = byte(' ')
	}
	sdsAvail(s)
	return string(s.buf)
}

func main() {
	fmt.Println("hello")
	fmt.Scanln(&str)
	UserAsk(str)

	sdsNew(str)

	//s.free = fr - s.len + 1
	//sdslen(&s)
	//sdsavail(&s)
}
