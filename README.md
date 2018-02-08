# dbcompare
level db, bolt db, badger db and bunt db benchmark compare
```go
go test -bench .

goos: linux
goarch: amd64
pkg: github.com/×××/dbcompare/benchmark
BenchmarkLevelDbPut-4              10000            326518 ns/op
BenchmarkBoltUpdate-4              20000             84175 ns/op
BenchmarkBadgerUpdate-4            50000             27792 ns/op
BenchmarkGoLevelDbBindPut-4        10000            113293 ns/op
BenchmarkBuntDbPut-4              100000             31235 ns/op
BenchmarkLevelDbGet-4             100000             83952 ns/op
BenchmarkBoltUpdateGet-4          100000            137974 ns/op
BenchmarkBadgerUpdateGet-4       1000000              2903 ns/op
BenchmarkGoLevelDbBindGet-4       100000            119436 ns/op
BenchmarkBuntDbGet-4              500000              5171 ns/op
PASS
ok      github.com/suyanlong/dbcompare/benchmark        353.750s


```
