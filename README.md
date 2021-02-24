# go语言学习demo

> "21世纪的c语言",go语言起源于2007年，并于2009年正式对外发布，
> Go 语言的主要目标是将静态语言的安全性和高效性与动态语言的易开发性进行有机结合，
> 达到完美平衡，从而使编程变得更加有乐趣，而不是在艰难抉择中痛苦前行。

## go基础

### 安装配置go环境 

window系统使用scoop安装

```shell
scoop install golang
```

参考[GOPROXY](https://goproxy.io/zh/) 更改go module代理

~~~shell
go mod init  # 初始化 go.mod
go mod tidy  # 更新依赖文件
go mod download  # 下载依赖文件

go mod vendor  # 将依赖转移至本地的 vendor 文件
go mod edit  # 手动修改依赖文件
go mod graph  # 打印依赖图
go mod verify  # 校验依赖
~~~

_安装go的ide工具，我这里使用的是jetbrains的goland工具,或者使用idea，安装go的开发插件_

GOPATH路径在用户目录的go文件夹，也可以自行更改，window修改环境变量即可

- **src**：存放源代码
- **bin**：编译后生成的可执行文件
- **pak**：编译后生成的文件

go module常用命令


### 基本结构与基本数据类型

### 空控制结构

### 函数

### 数据、切片
 
### Map

### 包
 
### 结构与方法

### 接口与反射

## go高级

## 参考

[go入门学习笔记](https://github.com/xinliangnote/Go)

[the-wat-to-go中文文档](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)