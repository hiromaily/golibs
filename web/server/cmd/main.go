package main

import (
	"os"

	"github.com/hiromaily/golibs/config"
	ms "github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	lg "github.com/hiromaily/golibs/log"
	"github.com/hiromaily/golibs/web/server"
)

func init() {
	//log
	lg.InitializeLog(lg.DebugStatus, lg.NoDateNoFile, "[JWT]", "", "hiromaily")
}

func main() {
	//config
	tomlPath := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/settings.default.toml"
	c := config.New(tomlPath, false)

	//db
	err := ms.New(c.MySQL.Host, c.MySQL.DbName, c.MySQL.User, c.MySQL.Pass, c.MySQL.Port)
	if err != nil {
		panic(err)
	}
	db := ms.GetDB()

	//redis
	rd := redis.New(c.Redis.Host, c.Redis.Port, c.Redis.Pass, 0)

	//server
	server.NewServer(c, db, rd)

}
