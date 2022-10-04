# SimpleHttpPostServer
## 简介
这是一个用go语言写的简单http服务，它只接收post请求。每次接收到请求他会尝试把请求体的内容解析为字符串并打印出来。
## 编译
```go
go build
```
## 使用
编译后在项目目录下会出现一个可执行文件`simpleHttpPostServer`
- 可以使用--help查看参数设置
```go
./simpleHttpPostServer --help
// 输出
//flag needs an argument: -h
//Usage of ./simpleHttpPostServer:
//  -h string
//        host (default "0.0.0.0")
//  -p string
//        端口 (default "8080")
```
你可以自己指定要绑定的host和port。

指定host用-h传参(默认0.0.0.0)，指定port用-p传参(默认8080)。

比如
```go
./simpleHttpPostServer -h 0.0.0.0 -p 8000
```
或者
```go
./simpleHttpPostServer -p 8000
```
也可以什么都不传，直接使用默认的0.0.0.0:8080
```go
./simpleHttpPostServer
```

