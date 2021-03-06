# gomoku

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg)](http://godoc.org/github.com/ohtomi/gomoku)
[![License](https://img.shields.io/badge/license-mit-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/ohtomi/gomoku.svg?branch=master)](https://travis-ci.org/ohtomi/gomoku)
[![codecov](https://codecov.io/gh/ohtomi/gomoku/branch/master/graph/badge.svg)](https://codecov.io/gh/ohtomi/gomoku)

## Description

`gomoku` can help you make a HTTP server written in your favorite languages.

## Usage

```bash
$ gomoku init sample
$ cd sample
$ gomoku run --port 8080 --file ./gomoku.yml

$ curl -v -H 'x-gomoku:yes' http://localhost:8080/foo
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /foo HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
> x-gomoku:yes
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 28 Dec 2017 14:17:23 GMT
< Content-Length: 61
<
* Connection #0 to host localhost left intact
{"greeting": "hello, gomoku", "method": "GET", "url": "/foo"}
```

To see example configurations, please visit [this page](../master/example). 

## Configuration

### request block

```yaml
- request:
    method: get|post
    route: /foo
    headers:
      x-gomoku: yes
```

#### method (type: `regular expression`)
`method` is a regular expression of an HTTP request method.

#### route (type: `regular expression`)
`route` is a regular expression of an HTTP request URL.

#### headers (type: `map of string to regular expression`)
`headers` is a list of a regular expression of an HTTP request headers.

### command block
In `command` block, users can use `.Request` object at the inside of a template literal.

```yaml
- command:
    env:
    - GOMOKU=gomoku
    - METHOD={{ .Request.Method }}
    path: python3
    args:
    - -m
    - foo
    - '{{ .Request.URL.Path }}'
```

#### env (type: `string=string array`)
`env` will be exported before executing command.

#### path (type: `string`)
`path` is a path to executable command.

#### args (type: `any array`)
`args` will be passed as command line arguments to executable command.

### response block
In `response` block, users can use `.Request` object and `.Command` object at the inside of a template literal.

```yaml
- response:
    status: 200
    headers:
      content-type: application/json; charset=utf-8
    cookies:
    - name: session-id
      value: 12345
      path: /
      domain: localhost
      expires: 2018-01-14 22:47:17 JST
      maxage: 0
      secure: false
      httponly: false
    body: >
      {
      "greeting": "{{ .Command.ParseStdoutAsJson.greet }}",
      "method": "{{ .Command.ParseStdoutAsJson.method }}",
      "url": "{{ .Command.ParseStdoutAsJson.url }}"
      }
```

#### status (type: `integer`)
`status` is an HTTP response status code.

To redirect, set `status` to `301`, `302`, `303`, `307`, `308`.

#### headers (type: `map of string to string`)
`headers` is an HTTP response headers.

#### cookies (type: `map array`)
`cookies` is an HTTP response cookies.

#### body (type: `string`)
`body` is an HTTP response body.

#### template (type: `string`)
`template` is a path to a template file representing an HTTP response body.

If `body` is set in the same `response` block, `body` will be used instead of `template`.

#### file (type: `string`)
`file` is a path to a content file representing an HTTP response body.

If `body` (or `template`) is set in the same `response` block, `body` (or `template`) will be used instead of `file`.

## Template

To generate textual outputs, see an API document of the [Go standard template engine](https://golang.org/pkg/text/template/).

### function
- `.GetByKey(map of string to string array, string)`: returns `string array`
- `.GetByIndex(string array, integer)`: returns `string`
- `.JoinWith(string array, string)`: returns joined `string array`

- `.ReadFile(string)`: returns `string` that is a content of the first of uploaded file
- `.ReadFiles(string)`: returns `string array` that is a list of a content of uploaded files

### variable

#### `.Request` object
- `.Method`: HTTP request method (`string`)
- `.URL`: HTTP request URL (`net/url`'s `URL`)
- `.Headers`: HTTP request headers (`map of string to string array`)
- `.RemoteAddr`: a remote address of an HTTP request (`string`)

- `.Body`: HTTP request body (`string` or `map`)
- `.ParseBodyAsJson()`: returns `JSON` object made from `.Body` (`map`)
- `.ParseBodyAsYaml()`: returns `YAML` object made from `.Body` (`map`)

- `.Form`: HTTP request form (`map of string to string array`)

#### `.Command` object
- `.Env`: exported variables (`string=string array`)
- `.Path`: path to executable command (`string`)
- `.Args`: command line arguments (`any array`)
- `.Dir`: working directory (`string`)

- `.Stdout`: standard output stream of executable command (`string`)
- `.ParseStdoutAsJson()`: returns `JSON` object made from `.Stdout` (`map`)
- `.ParseStdoutAsYaml()`: returns `YAML` object made from `.Stdout` (`map`)

- `.Stderr`: standard error stream of executable command (`string`)
- `.ParseStderrAsJson()`: returns `JSON` object made from `.Stderr` (`map`)
- `.ParseStderrAsYaml()`: returns `YAML` object made from `.Stderr` (`map`)

## Installation

```console
$ go get -u github.com/ohtomi/gomoku/cmd/gomoku
```

Or clone the [repository](https://github.com/ohtomi/gomoku) and run:
```console
$ make install
```

Or get binary from [release page](../../releases/latest).


## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

MIT

## Author

[Kenichi Ohtomi](https://github.com/ohtomi)
