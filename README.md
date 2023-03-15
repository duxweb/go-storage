<h1 align="center">Welcome to go-storage 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/duxweb/go-storage/blob/main/LICENSE\" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev/github.com/elliotxx/gulu"><img src="https://pkg.go.dev/badge/github.com/elliotxx/gulu.svg"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/elliotxx/gulu?branch=master"><img src="https://img.shields.io/coveralls/github/elliotxx/gulu/master"></a>
</p>

> This is a local, qiniu, cos, oss file storage integration library

### 🏠 [Homepage](https://github.com/duxweb/go-storage)

## Install

```sh
go get -u https://github.com/duxweb/go-storage/v1
```

## Usage

```go
import "github.com/duxweb/go-storage/v1"
```

## Example


```go
func main() {
    example := storage.New('local', map[string]string{
		"path": "./uploads",
	    "domain": "http://0.0.0.0:8080/uploads",
    })
    example.Write(context.Background(), "example.txt", "hello world!", map[string]any)
}
```

## Implemented methods

- New()
- Write()
- WriteStream()
- Read()
- ReadStream()
- Delete()
- PublicUrl()
- PrivateUrl()


### Creation method

New returns new storage with handlers.

The supported types are as follows:

- local
- qiniu
- cos
- oss

Please configure according to the driver type.

```go
example := storage.New('local', map[string]string{ ... })
```

Some drivers have optional write configurations that can be passed according to your needs.

```go
example.Write(ctx context.Context, path string, contents string, config map[string]any) error
```

```go
example.WriteStream(ctx context.Context, path string, stream io.Reader, config map[string]any) error
```

```go
example.Read(ctx context.Context, path string) (string, error)
```

```go
example.ReadStream(ctx context.Context, path string) (io.Reader, error)
```

```go
example.Delete(ctx context.Context, path string) error
```

```go
example.PublicUrl(ctx context.Context, path string) (string, error)
```

Some drivers do not support private links, so they will return public links instead.

```go
example.PrivateUrl(ctx context.Context, path string) (string, error)
```

## Run tests

You need to modify and configure the driver data

```sh
go test
```

## Test Coverage Report

The following table shows the test coverage results for this project:

| Package | Statements | Coverage |
| --- | --- |----------|
| `github.com/duxweb/go-storage/v1/main.go` | 9.61,11.14 | 100%     |
| `github.com/duxweb/go-storage/v1/main.go` | 12.15,14.8 | 100%      |
| `github.com/duxweb/go-storage/v1/main.go` | 15.15,17.8 | 100%      |
| `github.com/duxweb/go-storage/v1/main.go` | 18.13,20.8 | 100%      |
| `github.com/duxweb/go-storage/v1/main.go` | 21.13,23.8 | 100%      |
| `github.com/duxweb/go-storage/v1/main.go` | 25.2,25.15 | 100%      |

The overall test coverage for this project is 100%.

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