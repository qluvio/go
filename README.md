[![Sourcegraph](https://sourcegraph.com/github.com/qluvio/json-iterator/-/badge.svg)](https://sourcegraph.com/github.com/qluvio/json-iterator?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/qluvio/json-iterator)
[![Build Status](https://travis-ci.org/json-iterator/go.svg?branch=master)](https://travis-ci.org/json-iterator/go)
[![codecov](https://codecov.io/gh/json-iterator/go/branch/master/graph/badge.svg)](https://codecov.io/gh/json-iterator/go)
[![rcard](https://goreportcard.com/badge/github.com/qluvio/json-iterator)](https://goreportcard.com/report/github.com/qluvio/json-iterator)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/json-iterator/go/master/LICENSE)
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/json-iterator/Lobby)

A high-performance 100% compatible drop-in replacement of "encoding/json"

You can also use thrift like JSON using [thrift-iterator](https://github.com/thrift-iterator/go)

```
Go开发者们请加入我们，滴滴出行平台技术部 taowen@didichuxing.com
```

# Benchmark

![benchmark](http://jsoniter.com/benchmarks/go-benchmark.png)

Source code: https://github.com/qluvio/json-iterator-benchmark/blob/master/src/github.com/qluvio/json-iterator-benchmark/benchmark_medium_payload_test.go

Raw Result (easyjson requires static code generation)

| | ns/op | allocation bytes | allocation times |
| --- | --- | --- | --- |
| std decode | 35510 ns/op | 1960 B/op | 99 allocs/op |
| easyjson decode | 8499 ns/op | 160 B/op | 4 allocs/op |
| jsoniter decode | 5623 ns/op | 160 B/op | 3 allocs/op |
| std encode | 2213 ns/op | 712 B/op | 5 allocs/op |
| easyjson encode | 883 ns/op | 576 B/op | 3 allocs/op |
| jsoniter encode | 837 ns/op | 384 B/op | 4 allocs/op |

Always benchmark with your own workload. 
The result depends heavily on the data input.

# Usage

100% compatibility with standard lib

Replace

```go
import "encoding/json"
json.Marshal(&data)
```

with 

```go
import "github.com/qluvio/json-iterator"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Marshal(&data)
```

Replace

```go
import "encoding/json"
json.Unmarshal(input, &data)
```

with

```go
import "github.com/qluvio/json-iterator"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Unmarshal(input, &data)
```

[More documentation](http://jsoniter.com/migrate-from-go-std.html)

# How to get

```
go get github.com/qluvio/json-iterator
```

# Contribution Welcomed !

Contributors

* [thockin](https://github.com/thockin) 
* [mattn](https://github.com/mattn)
* [cch123](https://github.com/cch123)
* [Oleg Shaldybin](https://github.com/olegshaldybin)
* [Jason Toffaletti](https://github.com/toffaletti)

Report issue or pull request, or email taowen@gmail.com, or [![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/json-iterator/Lobby)
