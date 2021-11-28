# counter
使用异步的思想，使用channel来完成并发写入，无需加锁，达到高性能与并发安全的目的。



# 补充内容

## 1 sync_counter

使用sync.Mutex来进行同步的counter

## 2 基准测试

### 2.1 同步Incr

| counter     | sync_counter |
| ----------- | ------------ |
| 346.2 ns/op | 366.2 ns/op  |

### 2.2 同步Get

| counter     | sync_counter |
| ----------- | ------------ |
| 260.7 ns/op | 27.42 ns/op  |

### 2.3 并发Incr

| counter     | sync_counter |
| ----------- | ------------ |
| 348.7 ns/op | 561.1 ns/op  |

### 2.4 并发Get

| counter     | sync_counter |
| ----------- | ------------ |
| 372.8 ns/op | 579.6 ns/op  |

## 3 结论

- 同步锁的方式，在同步的情况下表现优秀。当存放并发时，由于会争夺临界资源，所以速度会明显下降，不如异步channel方式。
- 异步channel方式，最大的优点在于异步，不必等待即可返回。缺点是channel需要较大的缓冲区。且channel通信本身有开销。



