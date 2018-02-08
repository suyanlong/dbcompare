# dbcompare
level db, bolt db, badger db and bunt db benchmark compare
```go
go test -bench .

*** b2d5aad commit benchmark result****
goos: linux
goarch: amd64
pkg: github.com/suyanlong/dbcompare/benchmark
BenchmarkLevelDbPut-4             200000             12916 ns/op
BenchmarkBoltUpdate-4                500           2663169 ns/op
BenchmarkBadgerUpdate-4             1000           2531718 ns/op
BenchmarkGoLevelDbBindPut-4       300000             30954 ns/op
BenchmarkBuntDbPut-4              200000             12824 ns/op
BenchmarkLevelDbGet-4             300000              7776 ns/op
BenchmarkBoltUpdateGet-4            1000           2048873 ns/op
BenchmarkBadgerUpdateGet-4      *** Test killed: ran too long (10m0s).  ==> fuck!
FAIL    github.com/suyanlong/dbcompare/benchmark        605.038s



*** current commit benchmark result****
goos: linux
goarch: amd64
pkg: github.com/suyanlong/dbcompare/benchmark
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
