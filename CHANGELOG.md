# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased
### Added

- Add `websocket` support

<!--
- from dep to vgo
- `HTTP streaming` support
- `gRPC` support
- `Consumer Driven Contract` support
  - mock server
  - mock client
 -->

## 0.6.2 (2018-11-04)

### Fixed

- HTTP headers matcher returns true if all given regular expressions match 

## 0.6.1 (2018-10-30)

### Fixed

- Fix `access-control-allow-origin` header bug

## 0.6.0 (2018-10-21)

### Added

- Support `NO_COLOR` environment variable

### Changed

- Print details of HTTP request when verbose mode
- Remove `--verbose` flag, use `DEBUG` environment variable instead

## 0.5.0 (2018-10-12)

### Added

- Support TLS mode

## 0.4.0 (2018-01-16)

### Added

- Support binary file as HTTP response
- Support `cookie` in HTTP response

### Changed

- Reload config file when it changed
- Rename template functions to parse request body or command's output
- Test whether every HTTP header matches given value in `request` block

## 0.3.0 (2017-12-31)

### Added

- Add header matcher to `request` block
- Add `post body` support
- Add `multi-part form` support

- Add functions to template directive
  - to make `JSON` object from request body
  - to make `YAML` object from request body
  - to read content from uploaded file
  - to read contents from uploaded files

- Respond 500 internal server error if `--error-no-match` enabled

### Fixed

- Fix empty form in conversation bug

## 0.2.0 (2017-12-22)

### Added

- Add functions to template directive
  - to get value from specified `map` object by key
  - to get value from specified `list` object by index
  - to join string array with specified separator

- Add functions to command object
  - to make `JSON` object from command's stdout or stderr
  - to make `YAML` object from command's stdout or stderr

- Add `CORS` support

## 0.1.0 (2017-12-06)

Initial release

### Added

- Add Fundamental features
