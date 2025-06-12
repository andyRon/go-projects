清空文件三种方式
---
ch030_clear_file

### 为什么需要清空TXT文件？
在实际应用中，清空TXT文件的需求多种多样：

- 日志轮转：为了避免日志文件无限增长，定期清空或截断日志文件是常见做法。
- 数据重置：在某些场景下，需要重置数据文件以开始新一轮的数据写入。
- 内存优化：清空不再需要的文件内容，释放磁盘空间，优化系统性能。

### Golang文件操作基础
在Golang中，文件操作主要通过os包实现。以下是一些常用的文件操作函数：

- os.Open：打开文件
- os.Create：创建文件
- os.Truncate：截断文件
- os.WriteFile：写入文件
- os.Remove：删除文件

### ref

https://www.oryoy.com/news/gao-xiao-qing-kong-txt-wen-jian-golang-shi-xian-zui-jia-shi-jian-yu-dai-ma-shi-li-jie-xi.html