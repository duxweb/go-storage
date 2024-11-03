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


> Simple repository based on local and s3 protocols, supporting Alibaba Cloud, Tencent Cloud, Qiniuyun, Huawei Cloud, minio and other S3 compatible protocols.

> 基于本地和 s3 协议的简单存储库，支持 阿里云、腾讯云、七牛云、华为云、minio 和其他 S3 兼容协议。

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
    example := storage.New("s3", map[string]string{
        "region":    "cn-south-1", 
        "endpoint":  "s3.cn-south-1.qiniucs.com",
        "bucket":    "dux",
        "access_key": "",
        "secret_key": "",
		
		// public url
        "domain":    domain,
		
		// optional
		"ssl": "true"
		"immutable": "true"
    }, nil)
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
- Size()
- Exists()

### Creation method 创建方法

New returns new storage with handlers.

New返回带有处理程序的新存储。

Support for local and s3 compatible repositories:

支持本地和 s3 兼容存储库：


- qiniu 七牛云存储
- cos 阿里云存储
- oss 腾讯云存储
- obs 华为云存储



```go
// 初始化S3存储库
example := storage.New("s3", map[string]string{
    map[string]string{
        "region":    "cn-south-1",
        "endpoint":  "s3.cn-south-1.qiniucs.com",
        "bucket":    "dux",
        "access_key": "",
        "access_key": "",
        
        // public url 公共地址，用于生成公共URL
        "domain":    "http://storage.test/upload",
        
        // optional
        "ssl": "true"
        "immutable": "true"
    }
}, nil)

// 初始化本地存储库
example := storage.New("local", map[string]string{
	"root": "./upload",
	// public url 公共地址，用于生成公共URL
	"domain": "http://storage.test/upload"
}, func(path string) (string, error) {
	return "Signature result"
})
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
// 文件大小
example.Size(ctx context.Context, path string) (int64, error)
```


```go
// 文件存在
Exists(ctx context.Context, path string) (bool, error)
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
// 获取 POST 上传预签名
example.SignPostUrl(ctx context.Context, path string) (url string, params map[string]string, err error)
```


```go
// 获取 PUT 上传预签名
example.SignPutUrl(ctx context.Context, path string) (string, error)
```

## Presignature Info

使用 PUT 方式上传，获取签名后直接使用返回地址 PUT 文件流，示例如下：

```ts
fetch(url, {
    method: 'PUT',
    headers: {
        'Content-Type': 'application/octet-stream'
    },
    body: file
})
```



使用 POST Form 方式上传，获取签名后使用返回的 `params` 参数代入文件表单，需要注意的是文件key为 `file`，并且文件放置在最后，否则可能无法请求，七牛有 `NotSupportAnonymous` 问题，js 如下：

```ts
const formData = new FormData()

Object.keys(params || {}).forEach((key) => {
	formData.append(key, params?.[key])
})
formData.append('file', file)

fetch(url, {
    method: 'POST',
    headers: {
        'Content-Type': 'multipart/form-data'
    },
    body: formData
})
```



## Local Description 
Local storage instructions, local storage using the local file system, support for all methods, local url signature need to configure their own initialization of the signature function, signature verification, please verify their own.

本地存储说明，本地存储使用本地文件系统，支持所有方法，本地url签名需要自行在初始化时配置签名函数，签名验证请自行进行验证。

## Run tests

You need to modify and configure the driver data

您需要修改并配置驱动数据

```sh
go test
```

## Test Coverage Report

coverage: 60.0% of statements

该项目测试覆盖率达60%，因部分方法为关联方法，故未覆盖。

coverage: 60% of statements

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