Stack inplementation in Go

1、 Array implementation
2、 Link implementation

---

1、 Array implementation benchmark result:

```
PASS
BenchmarkPush-4	 3000000	       458 ns/op
BenchmarkPop-4 	20000000	       116 ns/op
--- BENCH: BenchmarkPop-4
	stack_test.go:49: 0
	stack_test.go:49: 99
	stack_test.go:49: 9999
	stack_test.go:49: 999999
	stack_test.go:49: 19999999
ok  	github.com/holys/CS/data-structures/stack	17.466s
```

2、 Link implementation
