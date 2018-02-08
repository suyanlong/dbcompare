package main

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	levelDb()
	boltDb()
	badgerDb()
}

func badgerDb() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()
	// Your code here…
	db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("10"), []byte("100"))
	})

	db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte("1"))
		fmt.Println(item.ToString())
		return nil
	})
}

func boltDb() {
	db, err := bolt.Open("bolt.db", 0777, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println(db.Path())
	fmt.Println(db.GoString())

	db.Update(func(tx *bolt.Tx) error {
		val := []byte(fmt.Sprintf("%v", "suyanlong"))
		bucket, _ := tx.CreateBucketIfNotExists(val)
		bucket.Put(val, val)
		fmt.Println(err)
		return nil
	})
	fmt.Println("end")
}

func levelDb() {
	db, err := leveldb.OpenFile("./level.db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	keyvalue := []byte("suyanlong")
	db.Put(keyvalue, keyvalue, nil)
	value, err := db.Get(keyvalue, nil)
	fmt.Printf("key = %v, err = %v", string(value), err)
}
