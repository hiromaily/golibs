package server

import (
	"net/http"

	lg "github.com/hiromaily/golibs/log"
	"github.com/hiromaily/golibs/web/server/builder"
	"github.com/hiromaily/golibs/web/server/fetcher"
)

type Handler struct {
	builder builder.Builder
	fetcher fetcher.Fetcher
}

func NewHandler(reg Registrier) *Handler {
	return &Handler{
		builder: reg.NewBuilder(),
		fetcher: reg.NewFetcher(),
	}
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			lg.Error(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	// fetch
	creative, err := h.fetcher.Fetch()
	if err != nil {
		lg.Error(err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	// dat
	datString := req.URL.Query().Get("dat")

	// build
	err = h.builder.Build(res, datString, creative)
	if err != nil {
		lg.Error(err)
		res.WriteHeader(http.StatusInternalServerError)
	}
}
