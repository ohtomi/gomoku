## 0.3.0 (2018-12-31)

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