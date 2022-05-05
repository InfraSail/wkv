package util

import (
	"fmt"
	"reflect"

	"github.com/dchest/siphash"
)

var (
	siph = siphash.New([]byte("zGg2KpAXv9rEWJxs"))
)

// SipHash 底层使用 siphash 算法计算出 key 的 hash 值。
// 这个也是 Redis dict 默认使用的算法，这里简单封装了下，
// 以接受 string/[]byte/int* 类型。
func SipHash(v interface{}) uint64 {
	var data []byte
	switch iv := v.(type) {
	case string:
		data = []byte(iv)
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		data = []byte(fmt.Sprintf("%d", iv))
	default:
		panic(fmt.Sprintf("key type '%s' is not supported", reflect.TypeOf(v).String()))
	}
	siph.Reset()
	_, _ = siph.Write(data)
	return siph.Sum64()
}
