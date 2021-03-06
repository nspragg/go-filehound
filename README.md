# Filehound

[![GoDoc](https://godoc.org/github.com/nspragg/go-filehound/filehound?status.svg)](https://godoc.org/github.com/nspragg/go-filehound) 
[![Build Status](https://travis-ci.org/nspragg/go-filehound.svg)](https://travis-ci.org/nspragg/go-filehound) 
[![Coverage Status](https://coveralls.io/repos/github/nspragg/go-filehound/badge.svg?branch=master)](https://coveralls.io/github/nspragg/go-filehound?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nspragg/go-filehound)](https://goreportcard.com/report/github.com/nspragg/go-filehound)

> Flexible and fluent interface for searching the file system. Implementation of [FileHound](https://github.com/nspragg/filehound) in Go. 

## Installation

```
go get github.com/nspragg/go-filehound/filehound
```

<!-- ## Demo

<img src="https://cloud.githubusercontent.com/assets/917111/13683231/7e915c2c-e6fd-11e5-9d58-e7228cf76ccf.gif" width="600"/> -->

## Usage

The example below prints all of the files in a directory that have the `.json` file extension:

```go
import filehound

fh := filehound.New()
fh.Query(Paths("/tmp"))
fh.Query(Ext("json"))
files := fh.Find()

fmt.Println(files)
```

## Documentation
For more examples and API details, see [API documentation](https://godoc.org/github.com/nspragg/go-filehound/filehound)

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
