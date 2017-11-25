# gomoku

## Description

`gomoku` can help you make a mock server written in your favorite languages.

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

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

MIT
