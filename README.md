# 通过cgo静态链接libpcre

## golang的`regexp`的"问题"

只能将`string`或者`[]byte`看作是UTF8编码的字符串

这就导致一些字节数组匹配的需求无法满足，如`main.go`中`\xFF`的匹配

## pcre的编译

可以在github搜索很多cgo绑定的库，但基本都只支持linux和darwin,并且看起来都是动态链接的

经过测试可按照以下步骤完成windows的静态链接编译

## 依赖项

1. docker

1. [xgo](https://github.com/techknowlogick/xgo)

    ```bash
    docker pull techknowlogick/xgo:latest
    go get src.techknowlogick.com/xgo
    ```

## go绑定代码

参考 [https://github.com/jptosso/coraza-waf/blob/master/pkg/utils/pcre/pcre.go](https://github.com/jptosso/coraza-waf/blob/master/pkg/utils/pcre/pcre.go)

修改如下

```go
#cgo LDFLAGS: -L/usr/local/lib/ -lpcre
#define PCRE_STATIC 1
```

* 第一行保证linux amd64 下强制使用我们静态`libpcre.a`而不是系统的动态库

* 定义`PCRE_STATIC`宏,在pcre源码中的`pcredemo.c`中有这样一段话

    > If you want to statically link this program against a non-dll .a file, you must define PCRE_STATIC before including pcre.h

## 修复xgo的一处bug

在linux/386下应该指定`-m32`参数，才能正常链接静态库

```bash
cd docker
docker build . -t techknowlogick/xgo:fix
```

参考 [https://github.com/techknowlogick/xgo/pull/119](https://github.com/techknowlogick/xgo/pull/119)

## 交叉编译

```bash
xgo --deps https://ftp.pcre.org/pub/pcre/pcre-8.44.tar.gz \
    -ldflags="-w -s" -trimpath \
    --targets linux/amd64,linux/386,windows/386,windows/amd64  \
    -image=techknowlogick/xgo:fix \
    -out dist/pcre .
```
