package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/go-kit/kit/log"
	"os"
	"time"
)

const (
	dbName     = "test1.db"
	bucketName = "posts"
)

var logger = log.NewLogfmtLogger(os.Stdout)

func worker() {

	db, e := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if e != nil {
		logger.Log("open", e)
		return
	}
	db.Close()
}

func openTimeout() {
	db, e := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if e != nil {
		logger.Log("open", e)
		return
	}
	defer db.Close()

	go worker()
	time.Sleep(10 * time.Second)
}

func simpleUpdate() error {
	db, e := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if e != nil {
		logger.Log("open", e)
		return e
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return b.Put([]byte("2018-09-29"), []byte("Hello Bolt!"))
	})
	return nil
}

func simpleRead() error {
	db, e := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if e != nil {
		logger.Log("open", e)
		return e
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte("2018-09-29"))
		fmt.Printf("2018-09-29: %s\n", v)
		return nil
	})
	return nil
}

type boltGreetings struct {
	Ts  time.Time
	Msg string
}

func structData() {
	db, e := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if e != nil {
		logger.Log("open", e)
	}
	defer db.Close()
	greetings := &boltGreetings{
		Ts:  time.Now(),
		Msg: "Greetings from shawn.",
	}
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		msg, err := json.Marshal(greetings)
		if err != nil {
			return err
		}
		return b.Put([]byte(greetings.Ts.Format(time.RFC3339)), msg)
	})
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucketName)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}
		return nil
	})
}

func main() {
	//openTimeout();
	//simpleUpdate()
	//simpleRead()
	structData()
}
