# gomoku

## Description

`gomoku` can help you make a mock server written in your favorite languages.

## Usage

```bash
$ gomoku run --port 8080 --file path/to/gomoku.yml
```

### Configuration

#### `gomoku.yml`

```yaml
- request:
    route: /foo
    method: get|post
  command:
    path: python3
    args:
      - -m
      - foo
  response:
    status: 200
    headers:
      content/type: application/json; charset=utf-8
      x-custom-header: foo; bar; baz
    body: "stdout: {{ .CommandResult.Stdout }}"
```

#### `foo.py`

```python
#!/usr/bin/env python3
import sys
print('hello, gomoku! argv: {}'.format(sys.argv))
```

### Request & Response

```bash
$ curl -v http://localhost:8080/fuga
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /fuga HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< X-Custom-Header: foo; bar; baz
< Date: Thu, 23 Nov 2017 13:41:25 GMT
< Content-Length: 84
< Content-Type: text/plain; charset=utf-8
<
stdout: hello, gomoku: ['/Users/ohtomi/src/github.com/ohtomi/gomoku/sample/foo.py']
* Connection #0 to host localhost left intact
```

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

MIT
