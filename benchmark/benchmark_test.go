package benchmark

import (
	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
	goleveldb "github.com/golang/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tidwall/buntdb"
	"strconv"
	"testing"
)

var value = RandStringBytesMaskImprSrc(5120)
var valueByte = []byte(value)

func NewLevelDb() (db *leveldb.DB, err error) {
	return leveldb.OpenFile(" /tmp/level.db", nil)
}

func BenchmarkLeveDbPut(b *testing.B) {
	db, _ := NewLevelDb()
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Put([]byte(strconv.Itoa(i)), valueByte, nil)
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("bench"))
			b.Put([]byte(strconv.Itoa(i)), valueByte)
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkBadgerUpdate(b *testing.B) {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()
	// Your code here…
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(strconv.Itoa(i)), valueByte)
		})
	}
	b.StopTimer()
}

func BenchmarkGoLevelDbBindPut(b *testing.B) {
	db, _ := goleveldb.Open("./goleveldb", nil)
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Set([]byte(strconv.Itoa(i)), valueByte, nil)
	}
	b.StopTimer()
}

func BenchmarkBuntDbPut(b *testing.B) {
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

// Benchmark read data
func BenchmarkLeveDbGet(b *testing.B) {
	db, _ := NewLevelDb()
	defer db.Close()

	for i := 0; i < b.N; i++ {
		db.Put([]byte(strconv.Itoa(i)), valueByte, nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Get([]byte(strconv.Itoa(i)), nil)
	}

	b.StopTimer()
}

func BenchmarkBoltUpdateGet(b *testing.B) {
	db := MustOpenDB()
	defer db.MustClose()
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("bench"))
		return err
	}); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("bench"))
			b.Put([]byte(strconv.Itoa(i)), valueByte)
			return nil
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("bench"))
			b.Get([]byte(strconv.Itoa(i)))
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkBadgerUpdateGet(b *testing.B) {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()
	// Your code here…
	for i := 0; i < b.N; i++ {
		db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(strconv.Itoa(i)), valueByte)
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(txn *badger.Txn) error {
			_, err := txn.Get([]byte(strconv.Itoa(i)))
			return err
		})
	}
	b.StopTimer()
}

func BenchmarkGoLeveDbBindGet(b *testing.B) {
	db, _ := goleveldb.Open("./goleveldb", nil)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		db.Set([]byte(strconv.Itoa(i)), valueByte, nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Get([]byte(strconv.Itoa(i)), nil)
	}
	b.StopTimer()
}

func BenchmarkBuntDbGet(b *testing.B) {
	db, _ := buntdb.Open("bunt.db")
	defer db.Close()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *buntdb.Tx) error {
			tx.Set(strconv.Itoa(i), value, nil)
			return nil
		})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.View(func(tx *buntdb.Tx) error {
			tx.Get(strconv.Itoa(i))
			return nil
		})
	}
	b.StopTimer()
}
