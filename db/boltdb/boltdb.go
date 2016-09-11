package boltdb

import (
	"errors"
	"github.com/boltdb/bolt"
	lg "github.com/hiromaily/golibs/log"
	//"github.com/hiromaily/golibs/times"
	//"time"
)

var db *bolt.DB

// New is to create bold instance
func New(path string) {
	var err error
	//db, err = bolt.Open("/Users/hy/work/go/src/github.com/hiromaily/golibs/datafile", 0600, nil)
	db, err = bolt.Open(path, 0600, nil)
	if err != nil {
		lg.Fatal(err)
	}
}

// Close is to close connection
func Close() {
	db.Close()
}

// Set is to set data by key
func Set(table, key string, value []byte) error {
	//defer times.Track(time.Now(), "boltdb.Set()") //405.695µs

	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(table))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})
}

// Get is to get data by key
func Get(table, key string) ([]byte, error) {
	//defer times.Track(time.Now(), "boltdb.Get()") //7.202µs

	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(table))
		if bucket == nil {
			return errors.New("bucket not found")
		}
		value = bucket.Get([]byte(key))
		return nil
	})
	if err != nil {
		return []byte{}, nil
	}
	return value, nil
}
