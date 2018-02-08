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

func BenchmarkLevelDbPut(b *testing.B) {
	db, _ := leveldb.OpenFile(" /tmp/level.db", nil)
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
	db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte("bench"))
		for i := 0; i < b.N; i++ {
			bk.Put([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})
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
	db.Update(func(txn *badger.Txn) error {
		for i := 0; i < b.N; i++ {
			txn.Set([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})
	b.StopTimer()
}

func BenchmarkGoLevelDbBindPut(b *testing.B) {
	db, _ := goleveldb.Open("/tmp/goleveldb", nil)
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
	db.Update(func(tx *buntdb.Tx) error {
		for i := 0; i < b.N; i++ {
			tx.Set(strconv.Itoa(i), value, nil)
		}
		return nil
	})
	b.StopTimer()
}

// Benchmark read data
func BenchmarkLevelDbGet(b *testing.B) {
	db, _ := leveldb.OpenFile(" /tmp/level.db", nil)
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

	db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte("bench"))
		for i := 0; i < b.N; i++ {
			bk.Put([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	b.ResetTimer()
	db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte("bench"))
		for i := 0; i < b.N; i++ {
			bk.Get([]byte(strconv.Itoa(i)))
		}
		return nil
	})
	b.StopTimer()
}

func BenchmarkBadgerUpdateGet(b *testing.B) {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, _ := badger.Open(opts)
	defer db.Close()
	// Your code here…
	db.Update(func(txn *badger.Txn) error {
		for i := 0; i < b.N; i++ {
			txn.Set([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	b.ResetTimer()
	db.Update(func(txn *badger.Txn) error {
		for i := 0; i < b.N; i++ {
			txn.Get([]byte(strconv.Itoa(i)))
		}
		return nil
	})
	b.StopTimer()
}

func BenchmarkGoLevelDbBindGet(b *testing.B) {
	db, _ := goleveldb.Open("/tmp/goleveldb", nil)
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

	db.Update(func(tx *buntdb.Tx) error {
		for i := 0; i < b.N; i++ {
			tx.Set(strconv.Itoa(i), value, nil)
		}
		return nil
	})

	b.ResetTimer()
	db.View(func(tx *buntdb.Tx) error {
		for i := 0; i < b.N; i++ {
			tx.Get(strconv.Itoa(i))
		}
		return nil
	})
	b.StopTimer()
}
