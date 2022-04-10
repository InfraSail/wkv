package util

import (
	"errors"
	"strings"

	
)

const SDS_MAX_PREALLOC int = 1024 * 1024

var (
	i   int = 0
	fr  int = 0

	bu         []byte
)

type sdshdr struct {
	// 等于SDS所保存字符串的长度

	len int
	// 记录buf数组中未使用字节的数量
	free int
	// 字节数组，用于保存字符串
	buf []byte
}

// TODO: 放入键值对

func SetFunc(key string, value string) {


}
// TODO: 根据键取值
func GetFunc(key string) {

}


func UserAsk(ask_str string) (err error) {
	ask_split := strings.Split(ask_str, " ")
	if x := ask_split[0]; x == "set" || x == "SET" {
		// 存储键值对
		SetFunc(ask_split[1], ask_split[2])
		return nil

	} else if x := ask_split[0]; x == "get" || x == "GET" {
		GetFunc(ask_split[1])
		return nil
	} else {
		err := errors.New("error")
		return err
	}
	
}


// 创建一个包含给定go字符串的SDS
func sdsNew(str string) (s *sdshdr) {
	s = &sdshdr{i, fr, bu}
	sdsCat(s, str)

	return 
}

func (s *sdshdr) GetString() (str string) {
	 str = string(s.buf[:s.len])
	 return
	}


// 创建一个不包含任何内容的空 SDS

func sdsEmpty() *sdshdr {
	s := sdshdr{0, 0, bu}
	return &s
}


func sdsCatSpace(s *sdshdr, cat_str string) string {

	cat_buf := []byte(cat_str)
	bufLength := len(cat_buf)
	l := s.len
	if(bufLength <= s.free){
		s.free -= bufLength
		s.len += bufLength
		copy(s.buf[l:], cat_buf[:])

	}else{
		s.len += bufLength

		if(bufLength + s.len <= SDS_MAX_PREALLOC){

			s.free = s.len

		}else{

			s.free = SDS_MAX_PREALLOC

		}
		new_len := s.len + s.free

		newBuf := make([]byte,new_len)
		copy(newBuf[0: l], s.buf[0:l])
		copy(newBuf[l:], cat_buf[:])
		s.buf = newBuf
		
	}
	
	return "OK"
}

// 返回SDS已使用的字节数
func sdsLen(s *sdshdr) int {

	return s.len
}

// 返回SDS未使用的字节数
func sdsAvail(s *sdshdr) int {

	return s.free
}

// 清空SDS保存的字符串内容
func sdsClear(s *sdshdr) string {
	s.free = fr
	s.len = i
	s.buf = bu
	return "OK"
}

// 对比两个 SDS 字符串是否相同,-1不同,0相同
func sdsCmp(s1 *sdshdr, s2 *sdshdr) int {

	return strings.Compare(string(s1.buf), string(s2.buf))
}



// 增长字符串
func sdsCat(s *sdshdr, cat_str string) string {
	//var cat_buf[] byte	
	sdsCatSpace(s, cat_str)
	//sdsAvail(s)
	return "OK"
}


func removeDuplication_map(arr []byte) []byte {
    set := make(map[byte]struct{}, len(arr))
    j := 0
    for _, v := range arr {
        _, ok := set[v]
        if ok {
            continue
        }
        set[v] = struct{}{}
        arr[j] = v
        j++
    }

    return arr[:j]
}


// 缩减字符串


func sdsTrim(s *sdshdr, trim_str string)  string {
	/*strTrim := string(s.buf)
	trimBuf := []byte(trim_str)
	newTrimBuf := string(removeDuplication_map(trimBuf))
	fmt.Printf("%s",newTrimBuf)

	strTrim = strings.Trim(strTrim, "xe")
	s.buf = []byte(strTrim)
	trimLength := s.len - len(s.buf)
	s.len -= trimLength
	s.free += trimLength
	fmt.Printf("%s",strTrim)*/
	trimBuf := []byte(trim_str)
	left := 0
	newTrimBuf := removeDuplication_map(trimBuf)
	for _, trim_value := range newTrimBuf {
		for j, s_value := range s.buf {
			if trim_value == s_value {
				s.buf = s.buf[:j + copy(s.buf[j:], s.buf[j + 1:])]
				s.len --
				s.free ++
			}
		}
	}
	for _, trim_value := range newTrimBuf {
		for j, s_value := range s.buf {
			
			if trim_value == s_value {
				s.buf[j] = 0
				s.len --
				s.free ++
			}
		}
	}

	for j := range s.buf {
	if s.buf[j] == 0{
		left ++
	}else{
		break
	}
}
	newstr := make([]byte,s.len - left )
	copy(newstr[:],s.buf[left:])
	s.buf = newstr
	return string(s.buf)
}

// 保留 SDS 给定区间内的数据， 不在区间内的数据会被覆盖或清除。


func sdsRange(s *sdshdr, left int, right int) ( err error) {
	
	if right < left {
		err = errors.New("ERROR!")
		return err
	}
	if left < 0{
		left = 0
	}
	if right > s.len{
		right = s.len
	}

	ran := right - left
	
	s.free += (s.len - ran)
	s.len = ran
	newBuf := make([]byte,ran)
	copy(newBuf[:], s.buf[left:right])
	
	s.buf = newBuf
	
	return nil
}

// 指定sds长度
func sdsGrowZero(s *sdshdr, self_defined_len int) (ret string) {

	var tmp_str string
	for i, iter_num := 0, self_defined_len-s.len; i < iter_num; i++ {
	tmp_str += " "
	  }
	 sdsCat(s, tmp_str)
	 ret = string(s.buf[:self_defined_len])
	 return
}

//对给定的字符串按给定的sep 分隔符来切割
func sdsSeplitLen(s *sdshdr, sep string ) []string{
	newSep := removeDuplication_map([]byte(sep))
	for _, sep_value := range newSep {
		for j, s_value := range s.buf {
			if sep_value == s_value {
				s.buf[j] = 0
			}
		}
	}
	arraySeped := strings.Split(string(newSep), sep)
	return arraySeped
}

// 释放给定的sds
/*
func sdsFree(s *sdshdr) *sdshdr{
	
	s = &sdshdr{1,1,[]byte{'a'}}
	return s
}*/

// 返回一个新的sds，内容与给定的s 相同。
func sdsDup(s *sdshdr) *sdshdr{
	newSds := sdsNew(s.GetString())
	
	return newSds
}