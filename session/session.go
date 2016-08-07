package session

//TODO:work in progress
import (
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

//1.グローバルでユニークなIDの生成（sessionid）
//2.データの保存スペースを作成
//3.sessionのグローバルでユニークなIDをクライアントサイドに送信
func cookieJar() {
	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar}
	lg.Debugf("client: %v", client)
}

func Set(w http.ResponseWriter) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func Get(r *http.Request, w http.ResponseWriter) {
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)
}

func Get2(r *http.Request, w http.ResponseWriter) {
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}
