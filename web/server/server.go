package server

import (
	"net/http"

	"github.com/hiromaily/golibs/config"
	"github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/db/redis"
	"github.com/hiromaily/golibs/web/server/handlers"
	"github.com/hiromaily/golibs/web/server/middlewares"
	"github.com/hiromaily/golibs/web/server/validations"
)

type Server struct {
	srvMux *http.ServeMux
}

func NewServer(conf *config.Config, db *mysql.MS, rd *redis.RD) (*Server, error) {
	s := &Server{
		srvMux: http.NewServeMux(),
	}

	redisRegistry := NewRedisRegistry(conf, rd)
	dbRegistry := NewDBRegistory(conf, db)

	//redis
	redisHandler := NewHandler(redisRegistry)
	//db
	dbHandler := NewHandler(dbRegistry)

	//health check
	hcHandler := handlers.NewHealthChecker(true)

	s.srvMux.Handle("/redis", middlewares.RecoverMiddleware(validations.Redis(redisHandler)))
	s.srvMux.Handle("/db", middlewares.RecoverMiddleware(validations.DB(dbHandler)))
	s.srvMux.Handle("/hc", hcHandler)

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.srvMux.ServeHTTP(w, req)
}
