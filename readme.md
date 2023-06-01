# 说明
使用gocc来分析gops stack 输出. 查找相同的grountine堆栈. 


``` shell
  -block_limit int
        阻塞位置相同的协程数量限制 (default 5)
  -same_limit int
        阻塞位置和起始位置相同的协程数量限制 (default 5)
  -same_top int
        阻塞位置和起始位置相同的协程,打印最多的前几个堆栈 (default 1)
  -start_limit int
        起始位置相同的协程数量限制 (default 5)
```
