package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dgraph-io/badger"
	goleveldb "github.com/golang/leveldb"
	goopt "github.com/golang/leveldb/db"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	//levelDb()
	//boltDb()
	//badgerDb()
	fmt.Println("benchmark begin")
	benchmarkBadgerSet()
	benchmarkBoltSet()

	benchmarkGoLevelDbSet()
	benchmarkLevelDbSet()

	benchmarkBadgerGet()
	benchmarkBoltGet()

	benchmarkGoLevelDbGet()
	benchmarkLevelDbGet()
	fmt.Println("benchmark end!")
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
	db, err := bolt.Open("/tmp/bolt.db", 0777, nil)
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
	db, err := leveldb.OpenFile("/tmp/level.db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	keyValue := []byte("suyanlong")
	db.Put(keyValue, keyValue, nil)
	value, err := db.Get(keyValue, nil)
	fmt.Printf("key = %v, err = %v", string(value), err)
}

var (
	number    = *flag.Int("number", 100, "package")
	N         = *flag.Int("N", 10000, "for limit N")
	value     = RandStringBytesMaskImprSrc(512)
	valueByte = []byte(value)
	src       = rand.NewSource(time.Now().UnixNano())
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

/// RandStringBytesMaskImprSrc
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func benchmarkLevelDbSet() {
	op := opt.Options{WriteBuffer: 128 * 1024 * 1024}
	db, _ := leveldb.OpenFile("level.db", &op)
	defer db.Close()

	wo := opt.WriteOptions{
		Sync: true,
	}

	now := time.Now()
	for i := 0; i < N; i++ {
		batch := new(leveldb.Batch)
		for i := 0; i < number; i++ {
			batch.Put([]byte(fmt.Sprintf("{}{}{}{}", rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63())), valueByte)
		}
		//b.SetBytes(int64(cap(batch)))
		db.Write(batch, &wo)
	}
	fmt.Println("benchmarkLevelDbSet cast time: ", time.Now().Sub(now))
}

func benchmarkBoltSet() {
	op := bolt.Options{
		Timeout:         0,
		InitialMmapSize: 1024 * 1024 * 1024,
		NoGrowSync:      false,
	}
	db, _ := bolt.Open("bolt.db", 0666, &op)
	db.AllocSize = 128 * 1024 * 1024
	defer db.Close()

	now := time.Now()
	for i := 0; i < N; i++ {
		db.Update(func(tx *bolt.Tx) error {
			bk, _ := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(i)))
			for i := 0; i < number; i++ {
				bk.Put([]byte(fmt.Sprintf("{}{}{}{}", rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63())), valueByte)
			}
			return nil
		})
	}
	fmt.Println("benchmarkBoltSet cast time: ", time.Now().Sub(now))
}

func benchmarkBadgerSet() {
	opts := badger.DefaultOptions
	opts.Dir = "badger"
	opts.ValueDir = "badger"
	db, _ := badger.Open(opts)
	defer db.Close()

	now := time.Now()
	for i := 0; i < N; i++ {
		db.Update(func(txn *badger.Txn) error {
			for i := 0; i < number; i++ {
				txn.Set([]byte(fmt.Sprintf("{}{}{}{}", rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63())), valueByte)
			}
			return nil
		})
	}
	fmt.Println("benchmarkBadgerSet cast time: ", time.Now().Sub(now))
}

func benchmarkGoLevelDbSet() {
	op := goopt.Options{WriteBufferSize: 128 * 1024 * 1024}
	db, _ := goleveldb.Open("goleveldb", &op)
	defer db.Close()

	wo := goopt.WriteOptions{Sync: true}
	now := time.Now()
	for i := 0; i < N; i++ {
		batch := goleveldb.Batch{}
		for i := 0; i < number; i++ {
			batch.Set([]byte(fmt.Sprintf("{}{}{}{}", rand.Int63(), rand.Int63(), rand.Int63(), rand.Int63())), valueByte)
		}
		db.Apply(batch, &wo)
	}
	fmt.Println("benchmarkGoLevelDbSet cast time: ", time.Now().Sub(now))
}

// Benchmark read data
func benchmarkLevelDbGet() {
	op := opt.Options{WriteBuffer: 128 * 1024 * 1024}
	db, _ := leveldb.OpenFile("level.db", &op)
	defer db.Close()

	batch := new(leveldb.Batch)
	for i := 0; i < number; i++ {
		batch.Put([]byte(strconv.Itoa(i)), valueByte)
	}
	wo := opt.WriteOptions{
		Sync: true,
	}
	db.Write(batch, &wo)

	now := time.Now()
	for i := 0; i < N; i++ {
		val, _ := db.Get([]byte(strconv.Itoa(rand.Intn(number))), nil)
		if !bytes.Equal(val, valueByte) {
			panic("BenchmarkLevelDbGet error")
		}
	}
	fmt.Println("benchmarkLevelDbGet cast time: ", time.Now().Sub(now))
}

func benchmarkBoltGet() {
	op := bolt.Options{
		Timeout:         0,
		InitialMmapSize: 1024 * 1024 * 1024,
		NoGrowSync:      false,
	}
	db, _ := bolt.Open("bolt.db", 0666, &op)
	db.AllocSize = 128 * 1024 * 1024
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bk, _ := tx.CreateBucketIfNotExists([]byte("bench"))
		for i := 0; i < 10000; i++ {
			bk.Put([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	now := time.Now()
	for i := 0; i < N; i++ {
		db.View(func(tx *bolt.Tx) error {
			bk := tx.Bucket([]byte("bench"))
			if !bytes.Equal(valueByte, bk.Get([]byte(strconv.Itoa(rand.Intn(number))))) {
				panic("BenchmarkBoltGet error")
			}
			return nil
		})
	}
	fmt.Println("benchmarkBoltGet cast time: ", time.Now().Sub(now))
}

func benchmarkBadgerGet() {
	opts := badger.DefaultOptions
	opts.Dir = "badger"
	opts.ValueDir = "badger"
	db, _ := badger.Open(opts)
	defer db.Close()

	db.Update(func(txn *badger.Txn) error {
		for i := 0; i < number; i++ {
			txn.Set([]byte(strconv.Itoa(i)), valueByte)
		}
		return nil
	})

	now := time.Now()
	for i := 0; i < N; i++ {
		db.View(func(txn *badger.Txn) error {
			item, _err := txn.Get([]byte(strconv.Itoa(rand.Intn(number))))
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
	fmt.Println("benchmarkBadgerGet cast time: ", time.Now().Sub(now))
}

func benchmarkGoLevelDbGet() {
	op := goopt.Options{WriteBufferSize: 128 * 1024 * 1024}
	db, _ := goleveldb.Open("goleveldb", &op)
	defer db.Close()

	batch := goleveldb.Batch{}
	for i := 0; i < number; i++ {
		batch.Set([]byte(strconv.Itoa(i)), valueByte)
	}
	wo := goopt.WriteOptions{Sync: true}
	db.Apply(batch, &wo)

	now := time.Now()
	for i := 0; i < N; i++ {
		val, _ := db.Get([]byte(strconv.Itoa(rand.Intn(number))), nil)
		if !bytes.Equal(val, valueByte) {
			panic("BenchmarkGoLevelDbGet error")
		}
	}
	fmt.Println("benchmarkGoLevelDbGet cast time: ", time.Now().Sub(now))
}
