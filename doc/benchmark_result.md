## [bolt vs LevelDB, RocksDB](https://github.com/boltdb/bolt#leveldb-rocksdb)
LevelDB and its derivatives (RocksDB, HyperLevelDB) are similar to Bolt in that they are libraries bundled into the application, however, their underlying structure is a log-structured merge-tree (LSM tree). An LSM tree optimizes random writes by using a write ahead log and multi-tiered, sorted files called SSTables.
Bolt uses a B+tree internally and only a single file. Both approaches have trade-offs.
If you require a high random write throughput (>10,000 w/sec) or you need to use spinning disks then LevelDB could be a good choice. If your application is read-heavy or does a lot of range scans then Bolt could be a good choice.
One other important consideration is that LevelDB does not have transactions. It supports batch writing of key/values pairs and it supports read snapshots but it will not give you the ability to do a compare-and-swap operation safely. Bolt supports fully serializable ACID transactions.

## [bolt vs LMDB](https://github.com/boltdb/bolt#lmdb)
Bolt was originally a port of LMDB so it is architecturally similar. Both use a B+tree, have ACID semantics with fully serializable transactions, and support lock-free MVCC using a single writer and multiple readers.
The two projects have somewhat diverged. LMDB heavily focuses on raw performance while Bolt has focused on simplicity and ease of use. For example, LMDB allows several unsafe actions such as direct writes for the sake of performance. Bolt opts to disallow actions which can leave the database in a corrupted state. The only exception to this in Bolt is DB.NoSync.
There are also a few differences in API. LMDB requires a maximum mmap size when opening an mdb_env whereas Bolt will handle incremental mmap resizing automatically. LMDB overloads the getter and setter functions with multiple flags whereas Bolt splits these specialized cases into their own functions.

## [bodger vs lmdb boltdb](https://blog.dgraph.io/post/badger-lmdb-boltdb/)
Based on our benchmarks, Badger is at least 1.7✕-22.3✕ faster than LMDB and BoltDB when doing random writes. Sorted range iteration is a bit slower when value size is small, but as value sizes increase, Badger is 4✕-111✕ times faster. On the flip side, Badger is currently up to 1.9✕ slower when doing random reads.

## [bodger post](https://blog.dgraph.io/post/badger/)