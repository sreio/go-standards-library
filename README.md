# go-standard

![Go](https://raw.githubusercontent.com/zc2638/material/master/go/go.png)

Go常用规范定义，标准库方法使用示例，请注意这不是Go的中文版标准库

## [Golang 泛型](https://github.com/akutz/go-generics-the-hard-way)
## [Uber编码规范](./doc/style.md)
## <a href="https://github.com/opentracing-contrib/opentracing-specification-zh/blob/master/specification.md" target="_blank">OpenTracing链路追踪规范</a>

## Go Module代理
设置环境变量
```
GOPROXY=https://proxy.golang.org    # 官方推荐，国内还无法正常使用
GOPROXY=https://goproxy.cn          # 国内相对友好
GOPROXY=https://goproxy.io          # 通用
```

## TODO
Change to uint tests

## 简介

- [**archive**](./src/archive) &emsp;&emsp;&emsp;&emsp; tar/zip压缩操作
- [**bufio**](./src/bufio/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 有缓冲的I/O
- [**bytes**](./src/bytes/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 操作[]byte字节片
- [**compress**](./src/compress) &emsp;&emsp;&emsp; bzip2/flate/gzip/lzw/zlib压缩操作
- [**container**](./src/container) &emsp;&emsp;&emsp;&ensp;堆操作/双向链表/环形链表
- [**context**](./src/context/example.go) &emsp;&emsp;&emsp;&emsp;&nbsp;上下文类型
- [**crypto**](./src/crypto) &emsp;&emsp;&emsp;&emsp;&emsp;常用的密码（算法）
- [**database/sql**](./src/database/sql/example.go) &emsp;&emsp;数据库接口
- [**encoding**](./src/encoding) &emsp;&emsp;&emsp;&emsp;数据编码
- [**errors**](./src/errors/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 创建错误函数
- [**expvar**](./src/expvar/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;公共变量
- [**flag**](./src/flag/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;命令行参数解析
- [**fmt**](./src/fmt/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp; 格式化I/O
- [**go**](./src/go/example.md) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;GO常用命令
- [**hash**](./src/hash) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;提供hash函数的接口
- [**html**](./src/html/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 转义和解转义HTML文本
- [**image**](./src/image/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&ensp;实现了基本的2D图片库
- [**index/suffixarray**](./src/index/suffixarray/example.go) &ensp;使用内存后缀数组以对数时间实现子字符串搜索
- [**io**](./src/io/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp; I/O操作
- [**log**](./src/log) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;简单的日志服务
- [**math**](./src/math) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;基本的数学常数和数学函数
- [**mime**](./src/mime/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;实现了MIME的部分规定
- [**net**](./src/net) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket
- [**os**](./src/os) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 操作系统函数
- [**path**](./src/path/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 对斜杠分隔的路径的实用操作
- [**plugin**](./src/plugin/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 插件生成和加载
- [**reflect**](./src/reflect/example.go) &emsp;&emsp;&emsp;&emsp;&emsp; 反射操作任意类型对象
- [**regexp**](./src/regexp/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;正则表达式
- [**rutime**](./src/runtime) &emsp;&emsp;&emsp;&emsp;&emsp;&nbsp;提供和go运行时环境的互操作，如控制go程的函数
- [**sort**](./src/sort/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;排序切片和用户自定义数据集
- [**strconv**](./src/strconv/example.go) &emsp;&emsp;&emsp;&emsp;&ensp; 基本数据类型和其字符串类型的相互转换
- [**strings**](./src/strings/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;操作字符串
- [**sync**](./src/sync) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;提供了基本的同步基元，如互斥锁
- [**text**](./src/text) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&ensp;文本操作
- [**time**](./src/time/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 时间操作
- [**unicode**](./src/unicode) &emsp;&emsp;&emsp;&emsp;&ensp;unicode操作
- [**unsafe**](./src/unsafe/example.go) &emsp;&emsp;&emsp;&emsp;&emsp;提供了一些跳过go语言类型安全限制的操作

### Tip
channel作为参数
```go
chan int   // 可读可写
<-chan int // 只读
chan<- int // 只写
```

### 参考
- [https://github.com/zc2638/go-standard](https://github.com/zc2638/go-standard)
- [**中文版标准库文档**](https://studygolang.com/pkgdoc) | [**中文版标准库文档2**](http://www.php.cn/manual/view/35126.html)
- [**中文版标准库文档(Dash版)**](https://github.com/taigacute/GoDoc-CN)
- [**《Go入门指南》**](https://github.com/unknwon/the-way-to-go_ZH_CN)
- [**Mastering Go(玩转Go中文译本)**](https://github.com/hantmac/Mastering_Go_ZH_CN)

