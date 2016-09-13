package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

// https://godoc.org/github.com/garyburd/redigo/redis
// http://stackoverflow.com/questions/25708256/golang-selecting-db-on-a-redispool-in-redigo

// RD is struct for Redis
type RD struct {
	Pool *redis.Pool
	Conn redis.Conn
}

var rdInfo RD

// New is to create instance for singleton
func New(host string, port uint16, pass string, dbNo int) *RD {
	if rdInfo.Pool == nil {
		rdInfo.setPool(host, port, pass)
	}
	rdInfo.Connection(dbNo)

	return &rdInfo
}

// NewIns make a new instance
func NewIns(host string, port uint16, pass string) *RD {
	rd := &RD{}
	rd.setPool(host, port, pass)
	return rd
}

// GetRedis is to get instance. singleton architecture
//  This is for singleton design pattern
func GetRedis() *RD {
	if rdInfo.Pool == nil {
		//panic("Before call this, call New in addition to arguments")
		return nil
	}
	return &rdInfo
}

// Connection is to connect Redis server
func (rd *RD) setPool(host string, port uint16, pass string) {
	rd.Pool = &redis.Pool{
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

// Connection is to connect Redis server
func (rd *RD) Connection(dbNo int) {
	rd.Conn = rd.Pool.Get()
	rd.Conn.Do("SELECT", dbNo)
}

// ConnectionS is to connect Redis server using Send func
func (rd *RD) ConnectionS(dbNo int) {
	rd.Conn = rd.Pool.Get()
	rd.Conn.Send("SELECT", dbNo)
}

// Flush is to flush data
func (rd *RD) Flush(dbNo int) {
	rd.Conn = rd.Pool.Get()
	rd.Conn.Send("SELECT", dbNo)
	rd.Conn.Send("FLUSHALL")
	rd.Conn.Flush()
}

// GetAndCluster is for cluster
//TODO: not finished yet. work in progress.
func (rd *RD) GetAndCluster() (int, error) {
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

// Close is to close connection
func (rd *RD) Close() {
	rd.Pool.Close()
}
