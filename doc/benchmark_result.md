## [bolt vs LevelDB, RocksDB](https://github.com/boltdb/bolt#leveldb-rocksdb)
```
LevelDB and its derivatives (RocksDB, HyperLevelDB) are similar to Bolt in that they are libraries bundled into the application, however, their underlying structure is a log-structured merge-tree (LSM tree). An LSM tree optimizes random writes by using a write ahead log and multi-tiered, sorted files called SSTables.
Bolt uses a B+tree internally and only a single file. Both approaches have trade-offs.
If you require a high random write throughput (>10,000 w/sec) or you need to use spinning disks then LevelDB could be a good choice.
如果是10000w/s建议使用levelDB
If your application is read-heavy or does a lot of range scans then Bolt could be a good choice.
如果是读操作或者检索，boltDB是好的选择。
One other important consideration is that LevelDB does not have transactions.
It supports batch writing of key/values pairs and it supports read snapshots but it will not give you the ability to do a compare-and-swap operation safely. Bolt supports fully serializable ACID transactions.

```
## [bolt vs LMDB](https://github.com/boltdb/bolt#lmdb)
```
Bolt was originally a port of LMDB so it is architecturally similar. Both use a B+tree, have ACID semantics with fully serializable transactions, and support lock-free MVCC using a single writer and multiple readers.
The two projects have somewhat diverged. LMDB heavily focuses on raw performance while Bolt has focused on simplicity and ease of use. For example, LMDB allows several unsafe actions such as direct writes for the sake of performance. Bolt opts to disallow actions which can leave the database in a corrupted state. The only exception to this in Bolt is DB.NoSync.
There are also a few differences in API. LMDB requires a maximum mmap size when opening an mdb_env whereas Bolt will handle incremental mmap resizing automatically. LMDB overloads the getter and setter functions with multiple flags whereas Bolt splits these specialized cases into their own functions.
```
## [bodger vs lmdb boltdb](https://blog.dgraph.io/post/badger-lmdb-boltdb/)
```
Based on our benchmarks, Badger is at least 1.7✕-22.3✕ faster than LMDB and BoltDB when doing random writes. Sorted range iteration is a bit slower when value size is small, but as value sizes increase, Badger is 4✕-111✕ times faster. On the flip side, Badger is currently up to 1.9✕ slower when doing random reads.
```
## [bodger post](https://blog.dgraph.io/post/badger/)
```
Badger is a simple, efficient, and persistent key-value store. Inspired by the simplicity of LevelDB, it provides Get, Set, Delete, and Iterate functions. On top of it, it adds CompareAndSet and CompareAndDelete atomic operations (see GoDoc). It does not aim to be a database and hence does not provide transactions, versioning or snapshots. Those things can be easily built on top of Badger.
Badger separates keys from values. The keys are stored in LSM tree, while the values are stored in a write-ahead log called the value log. Keys tend to be smaller than values. Thus this set up produces much smaller LSM trees. When required, the values are directly read from the log stored on SSD, utilizing its vastly superior random read performance.
Badger stored on SSD
Guiding principles
These are the guiding principles that decide the design, what goes in and what doesn’t in Badger.

1. Write it purely in Go language.
2. Use the latest research to build the fastest key-value store.
3. Keep it simple, stupid.
4. SSD-centric design.
```
### 测试结果
```
go test -bench=. -timeout=100m -benchtime=10s  -args -number=10000
goos: linux
goarch: amd64
pkg: github.com/suyanlong/dbcompare/benchmark

//wirte
BenchmarkLevelDbSet-8                100         342658654 ns/op
BenchmarkBoltSet-8                   100         136600787 ns/op
BenchmarkGoLevelDbSet-8              100         135141703 ns/op
BenchmarkBadgerSet-8                 200          72120621 ns/op
BenchmarkBuntDbSet-8                 300          51254884 ns/op

//read
BenchmarkBoltGet-8                 10000           1791715 ns/op
BenchmarkLevelDbGet-8            2000000              7363 ns/op
BenchmarkBuntDbGet-8             5000000              3937 ns/op
BenchmarkGoLevelDbGet-8          5000000              3031 ns/op
BenchmarkBadgerGet-8            10000000              1586 ns/op

PASS
ok      github.com/suyanlong/dbcompare/benchmark        309.499s
```
## 磁盘占用：
```
2.2G    ./bunt.db
1.8G    ./badger
1.4G    ./bolt.db
119M    ./goleveldb
112M    ./level.db
```

## 占用磁盘空间排名：
```
levelDB < golevelDB < boltDB < badgerDB < buntDB
```
## 写效率排名：
```
BuntDb < Badger < GoLevelDb < BoltDB < LevelDb
```
## 各个数据库的优缺点：

#### badgerDB
优点：1. 速度最快。
     2. 接口简单、易用。
     3. 主要是针对SSD磁盘而设计的数据库。
     4. 通过构建 key的LSM trees

缺点：1. 占用磁盘最大。
     2. 没有快照功能。

#### boltDB
BoltDB设计源于LMDB，具有以下特点：
与LMDB一样，具有ACID功能，使用B+tree,支持事务（transaction）,以及MVCC(单写多读=>读写锁)，嵌入式的数据库。
优点：1. 接口简单、易用。通过使用一个内存映射的磁盘文件来管理数据，逻辑清晰，接口简单易用。
     2. 直接使用API存取数据，没有查询语句；
     3. 支持完全可序列化的ACID事务，这个特性比LevelDB强；

缺点： 1. 数据保存在内存映射的文件里。没有wal、线程压缩和垃圾回收；
      2. 通过COW技术，可实现无锁的读写并发，但是无法实现无锁的写写并发，这就注定了读性能超高，但写性能一般，适合与读多写少的场景。

#### leveldb:
优点：
    1. 速度比go官方实现的leveldb快。
    2. 有快照。
    3. 占用磁盘最小。
    4. 自带压缩功能。
缺点：
    1. 没有transaction,但是有batch功能。
    2. 写效率低。

### Bunt：
内存k-v数据库，不支持事务，不支持回滚操作。并且占用磁盘空间最大，（已经可以去除掉了。）