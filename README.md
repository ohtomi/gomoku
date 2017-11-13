# gomoku

## Description

`gomoku` can help you make a mock server written in your favorite languages.

## Usage

```bash
$ gomoku run config.yml
```

### Configuration

```yml
# config.yml
- route: /foo
  method: get|post
  handler:
    path: handlers/foo.sh
    input:
      type: list
      params:
        - "${form.hoge}"
        - ${form.fuga}
    output:
      type: ltsv
  response:
    status: ${output.status}
    headers:
      - 'content/type': application/json
      - x-foo: foo-foo-foo
    template: templates/foo.json
```

```bash
#!/bin/bash
# handlers/foo.sh
echo -e "status:200\tresult:'hoge is $1 and fuga is $2'"
```

```json
{
  "file": "templates/foo.json",
  "result": "${output.result}"
}
```

## Contributing

1. Fork it!
1. Create your feature branch: `git checkout -b my-new-feature`
1. Commit your changes: `git commit -am 'Add some feature'`
1. Push to the branch: `git push origin my-new-feature`
1. Submit a pull request :D

## License

MIT
