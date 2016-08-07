package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

// https://godoc.org/github.com/garyburd/redigo/redis
// http://stackoverflow.com/questions/25708256/golang-selecting-db-on-a-redispool-in-redigo

type RedisInfo struct {
	Pool *redis.Pool
	Conn redis.Conn
	DbNo uint8
}

var rdInfo RedisInfo

func New(host string, port uint16, pass string) {
	if rdInfo.Pool == nil {
		rdInfo.Pool = &redis.Pool{
			MaxIdle:   80,
			MaxActive: 12000, // max number of connections
			Dial: func() (redis.Conn, error) {
				var c redis.Conn
				var err error
				if pass != "" {
					//plus password
					c, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port), redis.DialPassword(pass))
				} else {
					c, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
				}
				if err != nil {
					panic(err.Error())
				}
				return c, err
			},
		}
	}
	rdInfo.DbNo = 0
}

// singleton architecture
func GetRedisInstance() *RedisInfo {
	if rdInfo.Pool == nil {
		//panic("Before call this, call New in addtion to arguments")
		return nil
	}
	return &rdInfo
}

func (rd *RedisInfo) Close() {
	rd.Pool.Close()
}

func (rd *RedisInfo) Connection(dbNo int) {
	rd.Conn = rd.Pool.Get()
	rd.Conn.Do("SELECT", dbNo)
}

func (rd *RedisInfo) ConnectionS(dbNo int) {
	rd.Conn = rd.Pool.Get()
	rd.Conn.Send("SELECT", dbNo)
}

func (rd *RedisInfo) Flush(dbNo int) {
	rd.Conn = rd.Pool.Get()
	rd.Conn.Send("SELECT", dbNo)
	rd.Conn.Send("FLUSHALL")
	rd.Conn.Flush()
}

//TODO: not finished yet. work in progress.
func (rd *RedisInfo) GetAndCruster() (int, error) {
	data, err := redis.Int(rd.Conn.Do("GET", "key1"))

	if err != nil {
		fmt.Println(err)

		//for Redis Cluster
		splitStr := strings.Split(fmt.Sprint(err), " ")

		if splitStr[0] == "MOVED" {
			address := splitStr[2]
			fmt.Println(address)

			//Re connect
			//c, err = redis.Dial("tcp", address)
		} else {
			return 0, err
		}
		return 0, err
	}
	return data, nil
}
