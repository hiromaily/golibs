package server

import (
	"github.com/hiromaily/golibs/config"
	"github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	"github.com/hiromaily/golibs/web/server/builder"
	"github.com/hiromaily/golibs/web/server/fetcher"
	"github.com/hiromaily/golibs/web/server/models/db"
	redismodel "github.com/hiromaily/golibs/web/server/models/redis"
)

type Registrier interface {
	NewBuilder() builder.Builder //interface
	NewFetcher() fetcher.Fetcher //interface
}

type registry struct {
	conf *config.Config
}

type redisRegistry struct {
	registry
	rd *redis.RD
}

type dbRegistry struct {
	registry
	db *mysql.MS
}

//-----------------------------------------------------------------------------
// Base Registry
//-----------------------------------------------------------------------------

func (r *registry) NewBuilder() builder.Builder {
	//return builder struct (interface)
	return builder.NewBuilder(
		r.newBannerBuilder(),
		r.newHTMLBuilder(),
	)
}

func (r *registry) newBannerBuilder() builder.BannerBuilder {
	return builder.NewBannerBuilder("random", "clickURL")
}

func (r *registry) newHTMLBuilder() builder.HTMLBuilder {
	return builder.NewHTMLBuilder("randomHTML", "clickHTMLURL")
}

//-----------------------------------------------------------------------------
// reidsRegistry
//-----------------------------------------------------------------------------

func NewRedisRegistry(conf *config.Config, rd *redis.RD) Registrier {
	return &redisRegistry{
		registry: registry{conf: conf},
		rd:       rd,
	}
}

func (r *redisRegistry) NewFetcher() fetcher.Fetcher {
	return fetcher.NewDatFetcher(
		"datString",
		redismodel.NewCreativeRepository(r.rd),
	)
}

//-----------------------------------------------------------------------------
// dbRegistry
//-----------------------------------------------------------------------------

func NewDBRegistory(conf *config.Config, db *mysql.MS) Registrier {

	return &dbRegistry{
		registry: registry{conf: conf},
		db:       db,
	}
}

func (r *dbRegistry) NewFetcher() fetcher.Fetcher {
	return fetcher.NewRadFetcher(
		"radString",
		db.NewCreativeRepository(r.db),
	)
}
