<h1 align="center"> go-storage</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="https://github.com/duxweb/go-storage/blob/main/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev"><img src="https://img.shields.io/github/go-mod/go-version/duxweb/go-storage"></a>
<a title="Coverage Status" target="_blank" href="https://github.com/duxweb/go-storage"><img src="https://img.shields.io/badge/coverage-100%25-green"></a>
<a title="Coverage Status" target="_blank" href="https://github.com/duxweb/go-storage"><img src="https://img.shields.io/github/downloads/duxweb/go-storage/total"></a>
</p>


> This is a local, qiniu, cos, oss file storage integration library


## Install

```sh
go get -u github.com/duxweb/go-storage
```

## Usage

```go
import "github.com/duxweb/go-storage"
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
| `github.com/duxweb/go-storage/main.go` | 9.61,11.14 | 100%     |
| `github.com/duxweb/go-storage/main.go` | 12.15,14.8 | 100%      |
| `github.com/duxweb/go-storage/main.go` | 15.15,17.8 | 100%      |
| `github.com/duxweb/go-storage/main.go` | 18.13,20.8 | 100%      |
| `github.com/duxweb/go-storage/main.go` | 21.13,23.8 | 100%      |
| `github.com/duxweb/go-storage/main.go` | 25.2,25.15 | 100%      |

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