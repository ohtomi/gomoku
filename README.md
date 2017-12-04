# gomoku

## Description

`gomoku` can help you make a HTTP server written in your favorite languages.

## Usage

```bash
$ gomoku init sample
$ cd sample
$ gomoku run --port 8080 --file ./gomoku.yml

$ curl -v http://localhost:8080/foo
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /foo HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sat, 25 Nov 2017 15:05:39 GMT
< Content-Length: 42
<
* Connection #0 to host localhost left intact
{"greeting": "hello, gomoku. url is /foo"}
```

## Configuration

### request

```yaml
- request:
    route: /foo
    method: get|post
```

#### route (type: `regular expression`)
`route` is a regular expression of an URL of an HTTP request.

#### method (type: `regular expression`)
`method` is a regular expression of a method of an HTTP request.

### command
In `command` block, users can use `.Request` object at the inside of a template literal.

```yaml
- command:
    env:
      - HOGE=hoge
    path: python3
    args:
      - -m
      - foo
      - '{{ .Request.URL.Path }}'
```

#### env (type: `string=string`)
`env` will be exported before executing command.

#### path (type: `string`)
`path` is a path to executable command.

#### args (type: `any array`)
`args` will be passed as command line arguments to executable command.

### response
In `response` block, users can use `.Request` object and `.Command` object at the inside of a template literal.

```yaml
- response:
    status: 200
    headers:
      content-type: application/json; charset=utf-8 
    body: '{"greeting": "{{ .Command.Stdout }}"}'
```

#### status (type: `integer`)
`status` is an HTTP response status code.

To redirect, set `status` to `301`, `302`, `303`, `307`, `308`.

#### headers (type: `map of string to string`)
`headers` is an HTTP response headers.

#### body (type: `string`)
`body` is an HTTP response body.

#### template (type: `string`)
`template` is a path to a template file representing an HTTP response body.

If `body` is defined in a same `response` block, the value of `body` will be used as an HTTP response body.

#### file (type: `string`)
`file` is a path to a content file representing an HTTP response body.

If `body` (or `template`) is defined in a same `response` block, the value of `body` (or `template`) will be used as an HTTP response body.

### .Request object
- `.Method`: HTTP request method (`string`)
- `.URL`: HTTP request URL (`net/url`'s `URL`)
- `.Headers`: HTTP request headers (`map of string to string array`)
- `.Form`: HTTP request form (`map of string to string array`)
- `.RemoteAddr`: a remote address of an HTTP request (`string`)

### .Command object
- `.Env`: exported variables (`string=string array`)
- `.Path`: path to executable command (`string`)
- `.Args`: command line arguments (`any array`)
- `.Dir`: working directory (`string`)
- `.Stdout`: standard output stream of executable command (`string`)
- `.Stderr`: standard error stream of executable command (`string`)

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

MIT
