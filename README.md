# Filehound

[![Build Status](https://travis-ci.org/nspragg/go-filehound.svg)](https://travis-ci.org/nspragg/go-filehound) [![Coverage Status](https://coveralls.io/repos/github/nspragg/go-filehound/badge.svg?branch=master)](https://coveralls.io/github/nspragg/go-filehound?branch=master)

> [WIP] Flexible and fluent interface for searching the file system

## Installation

```
go get https://github.com/nspragg/go-filehound
```

<!-- ## Demo

<img src="https://cloud.githubusercontent.com/assets/917111/13683231/7e915c2c-e6fd-11e5-9d58-e7228cf76ccf.gif" width="600"/> -->

## Usage

The example below prints all of the files in a directory that have the `.json` file extension:

```go
import filehound

files := filehound.Create().
  Paths("/tmp").
  Ext("json").
  Find()

fmt.Println(files)
```

## Documentation
For more examples and API details, see [API documentation](https://nspragg.github.io/go-filehound/)

## Build

```
go build github.com/nspragg/go-filehound/filehound
```

## Test

```
go test github.com/nspragg/go-filehound/filehound
```

To generate a test coverage report:

```
go test -coverprofile=coverage.out github.com/nspragg/go-filehound/filehound
go tool cover -html=coverage.out
```
## Contributing

* If you're unsure if a feature would make a good addition, you can always [create an issue](https://github.com/nspragg/go-filehound/issues/new) first.
* We aim for 100% test coverage. Please write tests for any new functionality or changes.
* Any API changes should be fully documented.
* Make sure your code meets our linting standards. Run `golint github.com/nspragg/go-filehound/filehound` to check your code.
* Maintain the existing coding style. Always run `gofmt`.
* Be mindful of others when making suggestions and/or code reviewing.
