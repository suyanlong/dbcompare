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


```
固态SSD

go run main.go

benchmark begin

benchmarkBadgerSet cast time:  37.021794095s
benchmarkBoltSet cast time:  1m20.12233871s
benchmarkGoLevelDbSet cast time:  1m26.119174521s
benchmarkLevelDbSet cast time:  1m24.257447006s
benchmarkBadgerGet cast time:  17.894939ms
benchmarkBoltGet cast time:  17.918811ms
benchmarkGoLevelDbGet cast time:  38.801512ms
benchmarkLevelDbGet cast time:  22.903695ms

benchmark end!

------------------------------------------------------------------------
普通机械硬盘

benchmark begin
benchmarkBadgerSet cast time:  8m18.04950085s
benchmarkBoltSet cast time:  14m45.598502581s
benchmarkGoLevelDbSet cast time:  8m21.642394108s
benchmarkLevelDbSet cast time:  8m26.888921908s
benchmarkBadgerGet cast time:  37.469583ms
benchmarkBoltGet cast time:  24.423392ms
benchmarkGoLevelDbGet cast time:  78.883051ms
benchmarkLevelDbGet cast time:  13.867441ms
benchmark end!


```
