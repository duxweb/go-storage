<h1 align="center"> go-storage V2</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/duxweb/go-storage/blob/main/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev"><img src="https://img.shields.io/github/go-mod/go-version/duxweb/go-storage"></a>
<a title="Coverage Status" target="_blank" href="https://github.com/duxweb/go-storage"><img src="https://img.shields.io/badge/coverage-100%25-green"></a>
<a title="Coverage Status" target="_blank" href="https://github.com/duxweb/go-storage"><img src="https://img.shields.io/github/downloads/duxweb/go-storage/total"></a>
</p>


> A simple repository based on the s3 protocol, supporting cos, oss, qiniu, obs, minio and other S3-compatible protocols.

> 基于 s3 协议的简单存储库，支持 阿里云、腾讯云、七牛云、华为云、minio 和其他 S3 兼容协议。

## Install 安装

```sh
go get -u github.com/duxweb/go-storage/v2
```

## Usage 使用

```go
import "github.com/duxweb/go-storage/v2"
```

## Example 示例


```go
func main() {
    example := storage.New(map[string]string{
        "region":    "cn-south-1", 
        "endpoint":  "s3.cn-south-1.qiniucs.com",
        "bucket":    "dux",
        "accessKey": "",
        "secretKey": "",
		
		// public url
        "domain":    domain,
		
		// optional
		"ssl": "true"
		"immutable": "true"
    })
    example.Write(context.Background(), "example.txt", "hello world!", map[string]any)
}
```

## Implemented methods 支持方法

- New()
- Write()
- WriteStream()
- Read()
- ReadStream()
- Delete()
- PublicUrl()
- PrivateUrl()
- SignPostUrl()
- SignPutUrl()

### Creation method 创建方法

New returns new storage with handlers.

New返回带有处理程序的新存储。

The supported types are as follows:

支持以下类型的S3兼容存储库。

- qiniu 七牛云存储
- cos 阿里云存储
- oss 腾讯云存储
- obs 华为云存储



```go
// 初始化方法
example := storage.New(map[string]string{ ... })
```

```go
// 写入字符串文件
example.Write(ctx context.Context, path string, contents string) error
```

```go
// 写入文件流
example.WriteStream(ctx context.Context, path string, stream io.Reader) error
```

```go
// 读取字符串文件
example.Read(ctx context.Context, path string) (string, error)
```

```go
// 读取文件流
example.ReadStream(ctx context.Context, path string) (io.Reader, error)
```

```go
// 删除文件
example.Delete(ctx context.Context, path string) error
```

```go
// 获取公共链接
example.PublicUrl(ctx context.Context, path string) (string, error)
```


```go
// 获取私有链接
example.PrivateUrl(ctx context.Context, path string) (string, error)
```


```go
// 获取 POST 上传预签名，获取后使用表单参数和表单文件 POST 到 URL
example.SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error)
```


```go
// 获取 PUT 上传预签名，获取后直接使用返回地址 PUT 文件
example.PrivateUrl(ctx context.Context, path string) (string, error)
```

## Run tests

You need to modify and configure the driver data

您需要修改并配置驱动数据

```sh
go test
```

## Test Coverage Report

The overall test coverage for this project is 76.5%.

该项目测试覆盖率达76.5%，因部分方法为关联方法，故未覆盖。

coverage: 76.5% of statements

ok      github.com/duxweb/go-storage/v2 0.603s


## Author

👤 **duxweb**

* Website: https://github.com/duxweb
* Github: [@duxweb](https://github.com/duxweb)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/duxweb/go-storage/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2023 [duxweb](https://github.com/duxweb).<br />
This project is [MIT](https://github.com/duxweb/go-storage/blob/main/LICENSE\) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_