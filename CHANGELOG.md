## 0.5.0 (2018-10-12)

### Added

- Support TLS mode

## 0.4.0 (2018-01-16)

### Added

- Support binary file as HTTP response
- Support `cookie` in HTTP reponse

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