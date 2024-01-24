# 切片删除

## 基本实现

```
PS E:\code\GolandProjects\slicesDeletIdx> go test -bench='.*CoreV.*' .
goos: windows
goarch: amd64
pkg: github.com/AA1HSHH/slicesDeletIdx
cpu: AMD Ryzen 7 5800H with Radeon Graphics
BenchmarkCoreV1-16        932458              1252 ns/op
BenchmarkCoreV2-16       4480432               268.7 ns/op
PASS
ok      github.com/AA1HSHH/slicesDeletIdx       3.192s

```


V2的append拼接比起V1的遍历有6倍的性能提升

## 缩容
切片的容量是长度的1.5倍则进行缩容， 缩容缩到长度的1.2倍