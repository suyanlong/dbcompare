package test

import (
	"testing"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
	"github.com/tidwall/buntdb"
	goleveldb "github.com/golang/leveldb"
)

func NewLeveLDb() (db *leveldb.DB, err error) {
	return leveldb.OpenFile(" /tmp/level.db", nil)
}

func BenchmarkLeveDbPut(b *testing.B) {
	value := []byte(RandStringBytesMaskImprSrc(5120))
	db, _ := NewLeveLDb()
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Put([]byte(strconv.Itoa(i)), value, nil)
	}
	b.StopTimer()
}

func BenchmarkBoltUpdate(b *testing.B) {
	db := MustOpenDB()
	defer db.MustClose()
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("bench"))
		return err
	}); err != nil {
		b.Fatal(err)
	}
	value := []byte(RandStringBytesMaskImprSrc(5120))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("bench"))
			b.Put([]byte(strconv.Itoa(i)), value)
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkBadgerUpdate(b *testing.B) {
	value := []byte(RandStringBytesMaskImprSrc(5120))
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()
	// Your code here…
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(strconv.Itoa(i)), value)
		})
	}
	b.StopTimer()
}

func BenchmarkGolangLeveDbBindPut(b *testing.B) {
	value := []byte(RandStringBytesMaskImprSrc(5120))
	db, _ := goleveldb.Open("./goleveldb", nil)
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Set([]byte(strconv.Itoa(i)), value, nil)
	}
	b.StopTimer()
}

func BenchmarkBuntDbPut(b *testing.B) {
	value := RandStringBytesMaskImprSrc(5120)
	db, _ := buntdb.Open("bunt.db")
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *buntdb.Tx) error {
			tx.Set(strconv.Itoa(i), value, nil)
			return nil
		})
	}
	b.StopTimer()
}
