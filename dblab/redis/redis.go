package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

//TODO:これは、localからは動かない。。。ssh経由でredisに接続できればいいのだが。。。
type RedisDB struct {
	Hostname string
	Port     string
	DB       int
	Conn     redis.Conn
}

var host string = "127.0.0.1"
var port string = "6379"

//New
func New(num int) *RedisDB {
	r := &RedisDB{
		Hostname: host,
		Port:     port,
		DB:       num,
	}
	r.Start()

	return r
}

//Start
func (r *RedisDB) Start() {
	//TODO:Localからvagrantのredisに接続する場合は、sshの設定が必要
	var err error
	r.Conn, err = redis.Dial("tcp", host+":"+port)
	if err != nil {
		panic("Redis Error")
	}
	r.Conn.Do("SELECT", r.DB)

	//defer redisObj.Close()
	//hset 1:u9a734416aa5e21bfeb53f5e731d6a497 2 1990-01-01
}

//Register(6)
//redisに名前を登録する
func (r *RedisDB) RegisterTarget(name string) {
	//hset [key_name] [field] [value]  # keyにfieldとvalueのハッシュをset
	//r.Conn.Do("HSET", "14:"+name, "target", "true")
	//Do(commandName string, args ...interface{}) (reply interface{}, err error)
	r.Conn.Do("HSET", "spontena_state:14:"+name, "target", "true")
}

//Register(4)
//redisに誕生日を登録する
func (r *RedisDB) RegisterBd(name string, bd string) {
	r.Conn.Do("HSET", "14:"+name, "2", bd)

}

//Register(10)
//redisに佐賀県xタイ案件用で、location_idをセットする
//hset 1:u9a734416aa5e21bfeb53f5e731d6a497 location_id 1
//1:u9a734416aa5e21bfeb53f5e731d6a497
func (r *RedisDB) RegisterSightId(account_id int, mid string, sight_id int) {
	r.Conn.Do("HSET", strconv.Itoa(account_id)+":"+mid, "sight_id", sight_id)
	fmt.Println("set data to redis")
}

func (r *RedisDB) RegisterLocation(mid string) {
	r.Conn.Do("HMSET", "26:hug:"+mid, "latitude", "35.640353", "longitude", "139.712969")
	fmt.Println("set location to redis")
}

func (r *RedisDB) RegisterSagaEntry(mid string) {
	r.Conn.Do("HMSET", "25:hug:"+mid, "entry", "true")
	fmt.Println("set entry to redis")
}

func (r *RedisDB) RegisterPersonId(mid string) {
	r.Conn.Do("HMSET", "25:hug:"+mid, "person_id", "1")
	fmt.Println("set entry to redis")
}

func (r *RedisDB) RegisterPersonName(mid string) {
	r.Conn.Do("HMSET", "25:hug:"+mid, "person_name", "namedayo")
	fmt.Println("set entry to redis")
}

//Finish
func (r *RedisDB) Finish() {
	r.Conn.Close()
}
