package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
)

func main() {
	db, err := leveldb.OpenFile("./level.db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	keyVlue := []byte("suyanlong")
	db.Put(keyVlue, keyVlue, nil)
	value, err := db.Get(keyVlue, nil)
	fmt.Printf("key = %v, err = %v", string(value), err)
}
