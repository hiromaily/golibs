package middlewares

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"

	lg "github.com/hiromaily/golibs/log"
)

// RecoverMiddleware はpanic発生時にリカバリを行い、httpリクエストの受付を継続
// させる。panicが起こった場合はlogを通してログを吐くと共に、500レスポンスを
// 返す。
func RecoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s := stacktraceMsg(1) // skip this closure
				lg.Warnf("panic on %s: %+v, stack trace:\n%s", r.RequestURI, err, s)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func stacktraceMsg(depth int) string {
	var rawbuf []byte
	buf := bytes.NewBuffer(rawbuf)

	for d := depth + 1; ; d++ {
		_, f, l, ok := runtime.Caller(d)
		if !ok {
			break
		}
		buf.WriteString(fmt.Sprintf("\t%s at line %d\n", f, l))
	}

	return buf.String()
}
