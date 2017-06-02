package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func New(dbname, user, pass string) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable", dbname, user, pass))
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
