package util

import "time"

//输入的时间是否以秒为单位

const (
	UNIT_SECONDS = 0
	UNIT_MILLISECONDS = 1
)

type redisDB struct{
	dict *Dict //数据库键空间，存放着所有的键值对
	expires *Dict //键的过期时间

}

// 设置key过期时间命令,basetime为当前时间，unit为传入过期时间

func expireGenericCommand(baseTime int, unit int, db *redisDB)   {
	var k interface{} = "aaa"
	//when应该是用户传入吧？？
	when := 1234
	//如果传入的过期时间是以秒为单位的，那么将它转换为毫秒
	if unit == UNIT_SECONDS {
		when *= 1000
	}
	unit += baseTime

	//查询一下该键是否存在，不存在给客户端返回信息
	if _, env := lookUpKeyWrite(k, db);env == nil {
		return 
	}
	
	 /*TODO:
      * 在载入AOF数据时，或者服务器为附属节点时，
      * 即使 EXPIRE 的 TTL 为负数，或者 EXPIREAT 提供的时间戳已经过期，
      * 服务器也不会主动删除这个键，而是等待主节点发来显式的 DEL 命令。
     */
       //进入这个函数的条件：when 提供的时间已经过期，未载入数据且服务器为主节点（注意主服务器的masterhost==NULL）}

//设置expire
func setExpire(db *redisDB, key interface{},when int) {
	idx, ent:= db.dict.keyIndex(key)

}

// 找key，如果不存在，就返回nil
func lookUpKeyWrite(key interface{},db *redisDB) (idx uint64, existed *entry){
	 idx, ent:= db.dict.keyIndex(key)
	if(ent == nil){
		return 
	}
	return idx, ent

}
func expireCommand() {
	
	expireGenericCommand(mstime(), UNIT_SECONDS, &redisDB{})
}

func expireatCommand() {
	expireGenericCommand(0, UNIT_SECONDS, &redisDB{})
}
func pexpireCommand() {
	expireGenericCommand(mstime(), UNIT_MILLISECONDS, &redisDB{})
}
func pexpireatCommand() {
	expireGenericCommand(0, UNIT_MILLISECONDS, &redisDB{})
}

//返回UNIT当前时间
func mstime() int {
	return time.Now().Nanosecond()
}
