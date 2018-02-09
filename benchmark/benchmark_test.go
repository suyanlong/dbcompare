package benchmark

import (
	"bytes"
	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
	goleveldb "github.com/golang/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tidwall/buntdb"
	"math/rand"
	"strconv"
	"testing"
)

var value = RandStringBytesMaskImprSrc(512)
var valueByte = []byte(value)

func BenchmarkLevelDbSet(b *testing.B) {
	db, _ := leveldb.OpenFile("level.db", nil)
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		batch := new(leveldb.Batch)
		for i := 0; i < 10000; i++ {
			batch.Put([]byte(strconv.Itoa(int(rand.Int63()))), valueByte)
		}
		//b.SetBytes(int64(cap(batch)))
		db.Write(batch, nil)
	}
	b.StopTimer()
}

func BenchmarkBoltSet(b *testing.B) {
	db, _ := bolt.Open("bolt.db", 0666, nil)
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			bk, _ := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(i)))
			for i := 0; i < 10000; i++ {
				bk.Put([]byte(strconv.Itoa(int(rand.Int63()))), valueByte)
			}
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkBadgerSet(b *testing.B) {
	opts := badger.DefaultOptions
	opts.Dir = "badger"
	opts.ValueDir = "badger"
	db, _ := badger.Open(opts)
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(txn *badger.Txn) error {
			for i := 0; i < 10000; i++ {
				txn.Set([]byte(strconv.Itoa(int(rand.Int63()))), valueByte)
			}
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkGoLevelDbSet(b *testing.B) {
	db, _ := goleveldb.Open("goleveldb", nil)
	defer db.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		batch := goleveldb.Batch{}
		for i := 0; i < 10000; i++ {
			batch.Set([]byte(strconv.Itoa(int(rand.Int63()))), valueByte)
		}
		db.Apply(batch, nil)
	}
	b.StopTimer()
}

func BenchmarkBuntDbSet(b *testing.B) {
	db, _ := buntdb.Open("bunt.db")
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *buntdb.Tx) error {
			for i := 0; i < 10000; i++ {
				tx.Set(strconv.Itoa(int(rand.Int63())), value, nil)
			}
			return nil
		})
	}
	b.StopTimer()
}

// Benchmark read data
func BenchmarkLevelDbGet(b *testing.B) {
	db, _ := leveldb.OpenFile("level.db", nil)
	defer db.Close()

	batch := new(leveldb.Batch)
	for i := 0; i < 10000; i++ {
		batch.Put([]byte(strconv.Itoa(i)), valueByte)
	}
	db.Write(batch, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := db.Get([]byte(strconv.Itoa(rand.Intn(10000))), nil)
		if !bytes.Equal(val, valueByte) {
			panic("BenchmarkLevelDbGet error")
		}
	}
	b.StopTimer()
}

func BenchmarkBoltGet(b *testing.B) {
	db, _ := bolt.Open("bolt.db", 0666, nil)
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bk, _ := tx.CreateBucketIfNotExists([]byte("bench"))
		for i := 0; i < 10000; i++ {
			bk.Put([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			bk := tx.Bucket([]byte("bench"))
			if !bytes.Equal(valueByte, bk.Get([]byte(strconv.Itoa(rand.Intn(10000))))) {
				panic("BenchmarkBoltGet error")
			}
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkBadgerGet(b *testing.B) {
	opts := badger.DefaultOptions
	opts.Dir = "badger"
	opts.ValueDir = "badger"
	db, _ := badger.Open(opts)
	defer db.Close()

	db.Update(func(txn *badger.Txn) error {
		for i := 0; i < 10000; i++ {
			txn.Set([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.View(func(txn *badger.Txn) error {
			item, _err := txn.Get([]byte(strconv.Itoa(rand.Intn(10000))))
			if _err != nil {
				panic("BenchmarkBadgerGet error")
			} else {
				if val, err := item.Value(); err == nil {
					if !bytes.Equal(valueByte, val) {
						panic("BenchmarkBadgerGet error")
					}
				} else {
					panic("BenchmarkBadgerGet error")
				}
			}
			return nil
		})
	}
	b.StopTimer()
}

func BenchmarkGoLevelDbGet(b *testing.B) {
	db, _ := goleveldb.Open("goleveldb", nil)
	defer db.Close()

	batch := goleveldb.Batch{}
	for i := 0; i < 10000; i++ {
		batch.Set([]byte(strconv.Itoa(i)), valueByte)
	}
	db.Apply(batch, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := db.Get([]byte(strconv.Itoa(rand.Intn(10000))), nil)
		if !bytes.Equal(val, valueByte) {
			panic("BenchmarkGoLevelDbGet error")
		}
	}
	b.StopTimer()
}

func BenchmarkBuntDbGet(b *testing.B) {
	db, _ := buntdb.Open("bunt.db")
	defer db.Close()

	db.Update(func(tx *buntdb.Tx) error {
		for i := 0; i < 10000; i++ {
			tx.Set(strconv.Itoa(i), value, nil)
		}
		return nil
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db.View(func(tx *buntdb.Tx) error {
			val, _ := tx.Get(strconv.Itoa(rand.Intn(10000)))
			if val != value {
				panic("BenchmarkBuntDbGet error")
			}
			return nil
		})
	}
	b.StopTimer()
}
