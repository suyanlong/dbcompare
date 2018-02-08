package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

func main() {
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
