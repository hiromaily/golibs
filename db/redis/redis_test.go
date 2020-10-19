package redis_test

import (
	"os"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"

	conf "github.com/hiromaily/golibs/config"
	. "github.com/hiromaily/golibs/db/redis"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//http://qiita.com/nabewata07/items/10ab0008cb5e07b81a34
//http://qiita.com/taizo/items/82930518430f940721a0

//TODO:Sorted sets type on Redis, it's easy to total
//http://redis.shibu.jp/commandreference/sortedsets.html

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[REDIS]")

	c := conf.GetConf().Redis

	//New("localhost", 6379)
	New(c.Host, c.Port, c.Pass, 0)
	//if !tu.BenchFlg {
	//	GetRedis().Connection(0)
	//}
}

func teardown() {
	if !tu.BenchFlg {
		dropDatabase()
		r := GetRedis()
		r.Close()
	}
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// functions
//-----------------------------------------------------------------------------
func dropDatabase() {
	r := GetRedis()
	r.Flush(0)
	r.Flush(1)
	r.Flush(2)
	r.Flush(3)
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//EXPIRE(key, seconds), TTL(key), INFO
//-----------------------------------------------------------------------------
func TestCommonUsingDo(t *testing.T) {
	//tu.SkipLog(t)

	sleepS := 2

	c := GetRedis().Conn

	c.Do("MSET", "key1", 20, "key2", 30)
	c.Do("HMSET", "key3:subkey1", "field1", 1, "field2", 2)
	c.Do("HMSET", "key3:subkey2", "field1", 99, "field2", 100)

	val, err := redis.Int(c.Do("GET", "key1"))
	if err != nil {
		t.Errorf("key1 is 10: value is %v, error is %s", val, err)
	}

	fields, err := redis.Ints(c.Do("HMGET", "key3:subkey1", "field1", "field2"))
	if err != nil {
		t.Errorf("field1 should be 1, field2 should be 1 but result is %#v, error is %s", fields, err)
	}

	//EXPIRE
	c.Do("EXPIRE", "key1", sleepS)
	c.Do("EXPIRE", "key3:subkey1", sleepS)

	//TTL
	s, err := redis.Int(c.Do("TTL", "key1"))
	if err != nil {
		t.Errorf("redis.Int(c.Do(\"TTL\", \"key1\")) error : %s", err)
	}
	lg.Debugf("TTL is %v", s)

	//sleep
	time.Sleep(time.Duration(sleepS+1) * time.Second)

	s, _ = redis.Int(c.Do("TTL", "key1"))
	lg.Debugf("TTL is %v", s)

	//It can't access
	val, err = redis.Int(c.Do("GET", "key1"))
	if err == nil {
		t.Errorf("key1 has already expired: value is %v", val)
	}

	//It can't access
	//TODO:somehow it can access. but value is 0
	fields, err = redis.Ints(c.Do("HMGET", "key3:subkey1", "field1", "field2"))
	//if err == nil {
	if err != nil || fields[0] != 0 || fields[1] != 0 {
		t.Errorf("key3:subkey1 has already expired: value is %+v", fields)
	}

	//It's OK.
	fields, err = redis.Ints(c.Do("HMGET", "key3:subkey2", "field1", "field2"))
	if err != nil {
		//if err != nil || fields[0] != 99 || fields[1] != 100 {
		t.Errorf("field1 should be 99, field2 should be 100 but result is %#v", fields)
	}
}

//-----------------------------------------------------------------------------
//Strings (string, number, binary) up to 512MB
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//http://redis.shibu.jp/commandreference/strings.html

//SET(key, value)
//GET(key)

//MSET(key1, value1, key2, value2, ..., keyN, valueN)
//MGET(key1, key2, ..., keyN)

//DEL(key)
//-----------------------------------------------------------------------------

func TestStringsUsingDo(t *testing.T) {
	//tu.SkipLog(t)

	//GetRedisInstance().Connection(1)
	//GetRedisInstance().ConnectionS(2)

	c := GetRedis().Conn
	//c := GetRedisInstance().Pool.Get()
	c.Do("SET", "key1", 10)
	c.Do("MSET", "key2", 20, "key3", 30)

	val, err := redis.Int(c.Do("GET", "key1"))
	if err != nil || val != 10 {
		t.Errorf("key1 should be 10 but result is %d: err is %s", val, err)
	}

	vals, err2 := redis.Ints(c.Do("MGET", "key2", "key3"))
	if err2 != nil || vals[0] != 20 || vals[1] != 30 {
		t.Errorf("key2 should be 20, key2 should be 30, but result is %#v: err is %s", vals, err2)
	}
}

//send is faster than do method
func TestStringsUsingSend(t *testing.T) {
	//tu.SkipLog(t)

	GetRedis().ConnectionS(3)

	c := GetRedis().Conn
	c.Send("SET", "key1", 10)
	c.Send("MSET", "key2", 20, "key3", 30)
	c.Flush()
	for i := 0; i < 3; i++ {
		c.Receive() //OK
	}

	//GET
	c.Send("GET", "key1")
	c.Flush()
	val, err := redis.Int(c.Receive())
	if err != nil || val != 10 {
		t.Errorf("key1 should be 10 but result is %d: err is %s", val, err)
	}

	//MGET
	c.Send("MGET", "key2", "key3")
	c.Flush()
	vals, err2 := redis.Ints(c.Receive())
	if err2 != nil || vals[0] != 20 || vals[1] != 30 {
		t.Errorf("key2 should be 20, key2 should be 30, but result is %#v: err is %s", vals, err2)
	}
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//Hashes (string, number, binary) up to 512MB
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//HSET(key, field, value)
//HGET(key, field)

//HMSET(key, field1, value1, ..., fieldN, valueN)
//HMGET(key, field1, ..., fieldN)

//HGETALL(key) //get all elements.

//--DELETE--
//HDEL(key, field)

//-----------------------------------------------------------------------------

func TestHashesUsingDo(t *testing.T) {
	//tu.SkipLog(t)

	c := GetRedis().Conn
	//c := GetRedisInstance().Pool.Get()
	c.Do("HMSET", "key:subkey1", "field1", 1, "field2", 2)

	fields, err := redis.Ints(c.Do("HMGET", "key:subkey1", "field1", "field2"))
	if err != nil || fields[0] != 1 || fields[1] != 2 {
		t.Errorf("field1 should be 1, field2 should be 2 but result is %#v: err is %s", fields, err)
	}

	//HGETALL
	fields2, _ := redis.StringMap(c.Do("HGETALL", "key:subkey1"))
	lg.Debugf("HGETALL: %v, %s, %s", fields2, fields2["field1"], fields2["field2"])
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//Lists (Array for Redis Strings with order) up to 4.2 billion elements
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//http://redis.shibu.jp/commandreference/lists.html

//--SET--
//LPUSH(key, string)
//RPUSH(key, string)

//--UPDATE
//LSET(key, index, value)

//LTRIM(key, start, end)   //update lists

//--GET--
//LINDEX(key, index)       //get index.  e.g.)LINDEX("key1", 0), LINDEX("key1", -1)

//LRANGE(key, start, end)  //get range of lists

//LLEN(key)   //get length of lists

//--DELETE--
//LPOP(key)  //delete and get
//RPOP(key)  //delete and get

//複数のkeyを指定して、最初に見つかった値を保持するkeyと値を取得し、それを削除する
//BLPOP(key1, key2, ..., keyN, timeout)  //delete and get
//BRPOP(key1, key2, ..., keyN, timeout)

//LREM(key, count, value)

//-----------------------------------------------------------------------------
func TestListsUsingDo(t *testing.T) {
	//tu.SkipLog(t)

	GetRedis().Connection(1)
	//GetRedisInstance().ConnectionS(2)

	c := GetRedis().Conn
	//c := GetRedisInstance().Pool.Get()

	//RPUSH
	for i := 10; i < 101; i++ {
		c.Do("RPUSH", "key-list1", i)
	}
	vals, _ := redis.Ints(c.Do("LRANGE", "key-list1", 0, -1))
	lg.Debugf("key-list1 values is %v", vals)

	//LPUSH
	for i := 9; i > 0; i-- {
		c.Do("LPUSH", "key-list1", i)
	}
	vals, _ = redis.Ints(c.Do("LRANGE", "key-list1", 0, -1))
	lg.Debugf("key-list1 values is %v", vals)

	//LSET(key, index, value)
	c.Do("LSET", "key-list1", 0, 100)
	result, _ := redis.Int(c.Do("LINDEX", "key-list1", 0))
	lg.Debugf("result of LSET is %v", result)

	//LTRIM(key, start, end)   //update list
	c.Do("LTRIM", "key-list1", 0, 9)
	vals, _ = redis.Ints(c.Do("LRANGE", "key-list1", 0, -1))
	//LLEN(key)   //get length of lists
	length, _ := redis.Int(c.Do("LLEN", "key-list1"))
	lg.Debugf("key-list1 value is %v, length is %d", vals, length)
	vals, _ = redis.Ints(c.Do("LRANGE", "key-list1", 0, -1))
	lg.Debugf("key-list1 values is %v", vals)

	//LPOP(key)
	result, _ = redis.Int(c.Do("LPOP", "key-list1"))
	lg.Debugf("result of LPOP is %v", result)
	result, _ = redis.Int(c.Do("RPOP", "key-list1"))
	lg.Debugf("result of RPOP is %v", result)

	vals, _ = redis.Ints(c.Do("LRANGE", "key-list1", 0, -1))
	lg.Debugf("key-list1 values is %v", vals)
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//Sets (Array for Redis Strings without order) up to 4.2 billion elements
// Value is unique
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//http://redis.shibu.jp/commandreference/sets.html

//--SET--
//SADD(key, member)

//--GET--
//SMEMBERS(key)  //get all members

//SINTER(key1, key2, ..., keyN)   //get all member from designated keys??

//SCARD(key)    //get length of elements

//SUNION(key1, key2, ..., keyN)   //get all members(it's unique) from designated keys. union data.

//SISMEMBER(key, member)   //search member from key

//--UPDATE--??
//SMOVE(srckey, dstkey, member)   //move data from srckey to dstkey.

//SINTERSTORE(dstkey, key1, key2, ..., keyN)   //get all member from designated keys

//--DELETE--
//SREM(key, member)

//SPOP(key)        //delete a element at random

func TestSetsUsingDo(t *testing.T) {
	//tu.SkipLog(t)

	GetRedis().Connection(1)
	//GetRedisInstance().ConnectionS(2)

	c := GetRedis().Conn
	//c := GetRedisInstance().Pool.Get()

	key := "key-set1"

	//RPUSH
	for i := 0; i < 10; i++ {
		c.Do("SADD", key, i)
	}
	vals, _ := redis.Ints(c.Do("SMEMBERS", key))
	lg.Debugf("%s values is %v", key, vals)
}

//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
//Benchmark
//-----------------------------------------------------------------------------
//-----------------------------------------------------------------------------
func BenchmarkSetData(b *testing.B) {
	//b.Skip(fmt.Sprintf("skipping %s", r.CurrentFunc(1)))
	tu.SkipBLog(b)

	GetRedis().Connection(0)
	c := GetRedis().Conn

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Do("SET", "key1", 10)
	}
	b.StopTimer()

	dropDatabase()
	//35143 ns/op
}

func BenchmarkGetData(b *testing.B) {
	GetRedis().Connection(0)
	c := GetRedis().Conn

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		redis.Int(c.Do("GET", "key1"))
	}
	b.StopTimer()

	dropDatabase()
	//38652 ns/op
}

func BenchmarkSetGetData01(b *testing.B) {
	GetRedis().Connection(0)
	c := GetRedis().Conn

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Do("SET", "key1", 10)
		redis.Int(c.Do("GET", "key1"))
	}
	b.StopTimer()

	dropDatabase()
	//67251 ns/op
}

func BenchmarkSetGetData02(b *testing.B) {
	GetRedis().ConnectionS(3)
	c := GetRedis().Conn
	c.Flush()
	c.Receive()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Send("SET", "key1", 10)
		c.Flush()
		c.Receive() //OK

		c.Send("GET", "key1")
		c.Flush()
		redis.Int(c.Receive())
	}
	b.StopTimer()

	dropDatabase()
	//71453 ns/op
}

//Bulk
func BenchmarkSetBulkData01(b *testing.B) {
	GetRedis().Connection(0)
	c := GetRedis().Conn

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Do("SET", "key1", 10)
		c.Do("MSET", "key2", 20, "key3", 30)
		c.Do("HMSET", "key:subkey1", "field1", 1, "field2", 2)

		redis.Int(c.Do("GET", "key1"))
		redis.Ints(c.Do("MGET", "key2", "key3"))
		redis.Ints(c.Do("HMGET", "key:subkey1", "field1", "field2"))
	}
	b.StopTimer()

	dropDatabase()
	//220368 ns/op (220ms)
}

func BenchmarkSetBulkData02(b *testing.B) {
	GetRedis().ConnectionS(3)
	c := GetRedis().Conn
	c.Flush()
	c.Receive()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Send("SET", "key1", 10)
		c.Send("MSET", "key2", 20, "key3", 30)
		c.Send("HMSET", "key:subkey1", "field1", 1, "field2", 2)
		c.Flush()
		for i := 0; i < 3; i++ {
			c.Receive() //OK
		}

		//#1
		c.Send("GET", "key1")
		c.Flush()
		redis.Int(c.Receive())

		//#2
		c.Send("MGET", "key2", "key3")
		c.Flush()
		redis.Ints(c.Receive())

		//#3
		c.Send("HMGET", "key:subkey1", "field1", "field2")
		c.Flush()
		redis.Ints(c.Receive())
	}
	b.StopTimer()

	dropDatabase()
	//149114 ns/op (149ms)
}
