<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

# go语言学习demo

- [go语言学习demo](#go%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0demo)
    - [go基础](#go%E5%9F%BA%E7%A1%80)
        - [安装配置go环境](#%E5%AE%89%E8%A3%85%E9%85%8D%E7%BD%AEgo%E7%8E%AF%E5%A2%83)
        - [基本结构与基本数据类型](#%E5%9F%BA%E6%9C%AC%E7%BB%93%E6%9E%84%E4%B8%8E%E5%9F%BA%E6%9C%AC%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B)
        - [控制结构](#%E6%8E%A7%E5%88%B6%E7%BB%93%E6%9E%84)
        - [函数](#%E5%87%BD%E6%95%B0)
        - [数组、切片](#%E6%95%B0%E7%BB%84%E5%88%87%E7%89%87)
        - [Map](#map)
        - [结构与方法](#%E7%BB%93%E6%9E%84%E4%B8%8E%E6%96%B9%E6%B3%95)
        - [接口与反射](#%E6%8E%A5%E5%8F%A3%E4%B8%8E%E5%8F%8D%E5%B0%84)
    - [go高级](#go%E9%AB%98%E7%BA%A7)
    - [参考](#%E5%8F%82%E8%80%83)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## go基础

> "21世纪的c语言",go语言起源于2007年，并于2009年正式对外发布，
> Go 语言的主要目标是将静态语言的安全性和高效性与动态语言的易开发性进行有机结合，
> 达到完美平衡，从而使编程变得更加有乐趣，而不是在艰难抉择中痛苦前行。

### 安装配置go环境

window系统使用scoop安装

```shell
scoop install golang
```

参考[GOPROXY](https://goproxy.io/zh/) 更改go module代理

go module常用命令

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

### 基本结构与基本数据类型

- [变量申明](Demo02.go)
- [slice切片](Demo03.go)
- [strings和strconv包](Demo04.go)
- [取值符&和反引用*](Demo05.go)
- [条件及控制语句](Demo06.go)
- [函数申明](Demo07.go)
- [map和struct结构](Demo08.go)
- [new和make使用](Demo09.go)
- [接口](Demo10.go)
- [空接口](Demo11.go)
- [反射](Demo12.go)

## go高级

## 参考

[go入门学习笔记](https://github.com/xinliangnote/Go)

[the-wat-to-go中文文档](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)