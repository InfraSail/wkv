package util

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"time"
	"github.com/dchest/siphash"
)

// 初始化哈希表的容量
const (
	_initialHashtableSize uint64 = 4
)

type Dict struct {
	hashTables []*hashTable
	rehashIdx  int64 // 是否进行重哈希的标记量
	iterators  uint64 
}

type hashTable struct {
	buckets  []*entry
	size     uint64
	sizemask uint64 
	used     uint64 
}

type entry struct {
	key sdshdr
	value interface{}
	next       *entry
}

// New 实例化一个字典。
func NewDict() *Dict {
	return &Dict{
		// 初始化的时候，准备两张哈希表，默认使用哈希表 1
		// 在进行扩容时，会将哈希表 1 中的所有元素迁移到
		// 哈希表 2。
		hashTables: []*hashTable{{}, {},},
		rehashIdx:  -1,
		iterators:  0,
	}
}

func (d *Dict) String() string {
	return fmt.Sprintf("Dict(len=%d, cap=%d, isRehash=%v)", d.Len(), d.Cap(), d.isRehashing())
}

// Store 向字典中添加 key-value。
func (d *Dict) Store(key string, value interface{}) {
	sdsKey := sdsNew(key)
	ent, loaded := d.loadOrStore(sdsKey, value)
	if loaded {
		ent.value = value // 直接更新 value 即可
	} // 否则，上述函数调用会自动添加 (key, value) 到字典中
}

// Load 从字典中加载指定的 key 对应的值。
func (d *Dict) Load(key *sdshdr) (value interface{}, ok bool) {
	if d.isRehashing() {
		d.rehashStep()
	}

	_, existed := d.keyIndex(key)
	if existed != nil {
		return existed.value, true
	}

	return nil, false
}

// LoadOrStore 如果 key 存在于字典中，则直接返回其对应的值。
// 否则，该函数会将给定的值添加的字典中，并将给定的默认值返回。
// 如果能够在字典中成功查找的给定的 key，则 loaded 返回 true，
// 否则返回 false。
func (d *Dict) LoadOrStore(key *sdshdr, value interface{}) (actual interface{}, loaded bool) {
	ent, loaded := d.loadOrStore(key, value)
	if loaded {
		return ent.value, true
	} else {
		return value, false
	}
}

// Delete 从字典中删除指定的 key，如果 key 不存在，则什么也
// 不做。
// 实现描述：
// 1. 遍历哈希表，定位到对应的 buckets
// 2. 删除 buckets 中匹配的 entry。
func (d *Dict) Delete(key *sdshdr) {
	if d.Len() == 0 {
		// 不要做无畏的挣扎！
		return
	}

	if d.isRehashing() {
		d.rehashStep()
	}

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, key)
	if err != nil {
		return
	}
	hash := siphash.New(bytesBuffer.Bytes())

	for i := 0; i < 2; i++ {
		ht := d.hashTables[i]
		err := binary.Write(bytesBuffer, binary.BigEndian, ht.sizemask)
		if err != nil {
			return
		}
		idx := ht.sizemask & hash.Sum64()

		var prevEntry *entry
		for ent := ht.buckets[idx]; ent != nil; ent = ent.next {
			if string(ent.key.buf) == string(key.buf) {
				// 此时需要释放 ent 节点
				if prevEntry != nil {
					prevEntry.next = ent.next
				} else {
					// 说明待释放的节点是头节点，需要调整 buckets[idx] 指向下一个节点
					ht.buckets[idx] = ent.next
				}

				ent.next = nil
				ht.used--

				return
			}

			prevEntry = ent
		}

		if !d.isRehashing() {
			break
		}
	}
}

// Len 返回字典中元素的个数
func (d *Dict) Len() uint64 {
	var _len uint64
	for _, ht := range d.hashTables {
		_len += ht.used
	}
	return _len
}

// Cap 返回字典的容量
func (d *Dict) Cap() uint64 {
	if d.isRehashing() {
		return d.hashTables[1].size
	}
	return d.hashTables[0].size
}

// TODO :
// Range 以非安全的方式进行迭代，意味着在迭代期间不允许对字典执行额外的
// 操作，以免引起 rehash，导致重复扫描一些键。
// 用户传入的 `fn` 回调，可以通过返回 false 指示迭代器停止工作。
// 迭代完毕后，内部迭代器在释放时会自动对比当前的字典指纹和迭代器指纹是否
// 一致，如果不一致，表明执行了禁止的操作，进而引起 panic。
/*func (d *Dict) Range(fn func(key, value interface{}) bool) {
	d.rangeDict(fn, false)
}*/

// TODO :
// Range 以安全的方式进行迭代，可以在迭代期间执行 Load, Store 等操作，
// 它会在执行这些操作时，阻止字典进行 rehash 操作。但不保证新加入的键值
// 一定能够被扫描到。
// 用户传入的 `fn` 回调，可以通过返回 false 指示迭代器停止遍历。
/*func (d *Dict) RangeSafely(fn func(key, value interface{}) bool) {
	d.rangeDict(fn, true)
}*/

// Resize 让字典扩容或者缩容到一定大小。
// 注意，这里只是会准备好用于扩容的第二个哈希表，但真正的迁移还是分散
// 在多次 rehash 操作中。
func (d *Dict) Resize() error {
	if d.isRehashing() {
		return errors.New("dict is rehashing")
	}

	size := d.hashTables[0].used
	if size < _initialHashtableSize {
		size = _initialHashtableSize
	}
	return d.resizeTo(size)
}

// RehashForAWhile 执行 rehash 一段时间。
func (d *Dict) RehashForAWhile(duration time.Duration) int64 {
	tm := time.NewTimer(duration)
	defer tm.Stop()

	var rehashes int64
	for {
		select {
		case <-tm.C:
			return rehashes
		default:
			if d.rehash(100) {
				return rehashes
			}
			rehashes += 100
		}
	}
}

// loadOrStore 先尝试使用 key 查找，如果查找到则直接返回对应 entry，
// 否则，会添加新的 entry 到字典中，同时返回 nil，表示之前不存在。
func (d *Dict) loadOrStore(key *sdshdr, value interface{}) (ent *entry, loaded bool) {
	if d.isRehashing() {
		d.rehashStep()
		
	}

	_ = d.expandIfNeeded() // 这里简单起见，假设一定是可以扩容成功的，忽略了错误
	idx, existed := d.keyIndex(key)

	ht := d.hashTables[0]
	if d.isRehashing() {
		ht = d.hashTables[1]
	}

	if existed != nil {
		return existed, true
	} else {
		// 否则，需要在指定 bucket 添加新的 entry
		// 对于哈希冲突的情况，采用链地址法，在插入新的 entry 时，
		// 采用头插法，保证最近添加的在最前面
		entry := &entry{key: *key, value: value, next: ht.buckets[idx]}
		ht.buckets[idx] = entry
		ht.used++
	}

	return nil, false
}

// keyIndex 基于指定的 key 获得对应的 bucket 索引
// 如果 key 已经存在于字典中，则直接返回关联的 entry
func (d *Dict)  keyIndex(key *sdshdr) (idx uint64, existed *entry) {
	//only hash the key.buf
	hash := siphash.New(key.buf).Sum64()
	
	for i := 0; i < 2; i++ {
		ht := d.hashTables[i]
		idx = ht.sizemask & hash //?
		for ent := ht.buckets[idx]; ent != nil; ent = ent.next {
			if string(ent.key.buf) == string(key.buf) {
				return idx, ent
			}
		}

		if !d.isRehashing() {
			break
		}
	}

	// 如果字典处于 rehashing 中，上面的循环可以保证最后的 idx 一定位于
	// 第二个哈希表，从而保证依赖该接口的地方存储的新键一定进入到新的哈希表
	return idx, nil
}
//扩容
func (d *Dict) expandIfNeeded() error {
	if d.isRehashing() {
		// 此时表明扩容已经成功，正在进行迁移（rehash）
		return nil
	}

	if d.hashTables[0].size == 0 {
		// 第一次扩容，需要一定的空间存放新的 keys
		return d.resizeTo(_initialHashtableSize)
	}

	// 否则，根据负载因子判断是否需要进行扩容
	// 扩容策略简单粗暴，至少要是已有元素个数的二倍
	if d.hashTables[0].used == d.hashTables[0].size {
		return d.resizeTo(d.hashTables[0].used * 2)
	}

	return nil
}

func (d *Dict) resizeTo(size uint64) error {
	// 这里主要是要保证扩容大小符合要求，至少要比现有元素个数多
	if d.isRehashing() || d.hashTables[0].used > size {
		return errors.New("failed to resize")
	}

	size = d.nextPower(size)
	if size == d.hashTables[0].size {
		return nil
	}

	// 准备开始扩容
	var ht *hashTable
	if d.hashTables[0].size == 0 {
		// 第一次执行扩容，给 ht[0] 准备好，接下来 keys 可以直接放进来
		ht = d.hashTables[0]
	} else {
		ht = d.hashTables[1]
		// 表明需要开始进一步扩容，迁移 ht[0] -> ht[1]
		d.rehashIdx = 0
	}
	ht.size = size
	ht.sizemask = size - 1
	ht.buckets = make([]*entry, ht.size)
	return nil
}

// nextPower 找到匹配 size 的扩容大小
// 2^2 -> 2^3 -> 2^4 -> 2^5 -> ...
func (d *Dict) nextPower(size uint64) uint64 {
	if size >= math.MaxUint64 {
		return math.MaxUint64
	}

	i := _initialHashtableSize
	for i < size {
		i <<= 1 // i*= 2
	}

	return i
}

func (d *Dict) rehashStep() {
	if d.iterators == 0 {
		d.rehash(1)
	}
}

// rehash 实现渐进式 rehash 策略。基本思想就是，每次对最多
// steps 个 buckets 进行迁移。另外，考虑到可能旧的哈希表中
// 会连续遇到较多的空 buckets，导致耗费时间不受限制，这里还
// 限定最多遇到 10 * steps 个空 buckets 就退出。
func (d *Dict) rehash(steps uint64) (finished bool) {
	if !d.isRehashing() {
		return true
	}

	maxEmptyBucketsMeets := 10 * steps

	src, dst := d.hashTables[0], d.hashTables[1]
	//如果连续遇到10 * steps个空buckets
	for ; steps > 0 && src.used != 0; steps-- {
		// 扫描哈希表直到遇到非空的 bucket
		for src.buckets[d.rehashIdx] == nil {
			d.rehashIdx++
			maxEmptyBucketsMeets--
			if maxEmptyBucketsMeets <= 0 {
				return false
			}
		}

		// 把整个 bucket 上所有的 entry 都迁移走
		for ent := src.buckets[d.rehashIdx]; ent != nil; {
			next := ent.next
			bytesBuffer := bytes.NewBuffer([]byte{})
			err := binary.Write(bytesBuffer, binary.BigEndian, ent.key)
			if err != nil{
				return 
			}
			
			idx := siphash.New(bytesBuffer.Bytes()).Sum64() & dst.sizemask
			ent.next = dst.buckets[idx]
			dst.buckets[idx] = ent
			src.used--
			dst.used++

			ent = next

		}

		src.buckets[d.rehashIdx] = nil // 清空旧的 bucket
		d.rehashIdx++
	}

	// 如果迁移完毕，需要将 ht[0] 指向迁移后的哈希表
	if src.used == 0 {
		d.hashTables[0] = dst
		d.hashTables[1] = &hashTable{}
		d.rehashIdx = -1
		return true
	}

	return false
}

func (d *Dict) isRehashing() bool {
	return d.rehashIdx >= 0
}
