package validations

import (
	"net/http"

	lg "github.com/hiromaily/golibs/log"
)

// DBvalidate for delivery
func DB(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.Query()

		//validate for db
		if u.Get("data") == "" {
			lg.Error("data should be required")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Redis validate for redis
func Redis(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		u := r.URL.Query()

		//rad_preview
		if u.Get("data") == "" {
			lg.Error("rad_preview should be required")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
