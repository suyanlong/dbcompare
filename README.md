# dbcompare
level db, bolt db, badger db and bunt db benchmark compare
```go
cd benchmark

first run: go test -bench .

goos: linux
goarch: amd64
pkg: github.com/suyanlong/dbcompare/benchmark
BenchmarkLevelDbSet-4                 20         109431183 ns/op
BenchmarkBoltSet-4                    10         172565745 ns/op
BenchmarkBadgerSet-4                  30          73219897 ns/op
BenchmarkGoLevelDbSet-4               30         100224908 ns/op
BenchmarkBuntDbSet-4                  30          37309376 ns/op
BenchmarkLevelDbGet-4              50000             21898 ns/op
BenchmarkBoltGet-4                   500           2474387 ns/op
BenchmarkBadgerGet-4              500000              2981 ns/op
BenchmarkGoLevelDbGet-4           300000              5632 ns/op
BenchmarkBuntDbGet-4             1000000              2944 ns/op
PASS
ok      github.com/suyanlong/dbcompare/benchmark        35.536s

two run: go test -bench . -v
goos: linux
goarch: amd64
pkg: github.com/suyanlong/dbcompare/benchmark
BenchmarkLevelDbSet-4                 20         374127785 ns/op
BenchmarkBoltSet-4                    10         140051018 ns/op
BenchmarkBadgerSet-4                  20          68027663 ns/op
BenchmarkGoLevelDbSet-4               30         170119310 ns/op
BenchmarkBuntDbSet-4                  30          66782111 ns/op
BenchmarkLevelDbGet-4             100000             22287 ns/op
BenchmarkBoltGet-4                   500           2496795 ns/op
BenchmarkBadgerGet-4              500000              2282 ns/op
BenchmarkGoLevelDbGet-4           300000              5323 ns/op
BenchmarkBuntDbGet-4              500000              3649 ns/op
PASS
ok      github.com/suyanlong/dbcompare/benchmark        52.046s

three run: go test -bench . -v

goos: linux
goarch: amd64
pkg: github.com/suyanlong/dbcompare/benchmark
BenchmarkLevelDbSet-4                 20         607825638 ns/op
BenchmarkBoltSet-4                     5         205219111 ns/op
BenchmarkBadgerSet-4                  30          72856209 ns/op
BenchmarkGoLevelDbSet-4               30         300955982 ns/op
BenchmarkBuntDbSet-4                  20          50155096 ns/op
BenchmarkLevelDbGet-4             100000             20660 ns/op
BenchmarkBoltGet-4                   500           2834335 ns/op
BenchmarkBadgerGet-4             1000000              2350 ns/op
BenchmarkGoLevelDbGet-4           300000              5599 ns/op
BenchmarkBuntDbGet-4              500000              3697 ns/op
PASS
ok      github.com/suyanlong/dbcompare/benchmark        69.769s


```
