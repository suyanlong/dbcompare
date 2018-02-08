package main

import (
	"log"

	"github.com/dgraph-io/badger"
	"fmt"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
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
