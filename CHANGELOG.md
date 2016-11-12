Change Log
==========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/).

<!--
Added      new features.
Changed    changes in existing functionality.
Deprecated once-stable features removed in upcoming releases.
Removed    deprecated features removed in this release.
Fixed      bug fixes.
Security   invites users to upgrade in case of vulnerabilities.
-->

[Unreleased]
------------

### Added
- `func (f File) Split(sep string) []string`
  - Split the file into a string slice using separator sep.


[0.2.0] - 2016-11-07
--------------------

### Added
- Test status and coverage thanks to [travis-ci.org](https://travis-ci.org/nochso/golden)
  and [coveralls.io](https://coveralls.io/github/nochso/golden).
- `Case.Diff(string)` to compare with `Case.Out.String()` and print diff on failure.
- `TestDir()` to run named sub tests for each golden in a directory.
- Colourful diff output.

### Changed
- In absence of `testing.T` an error will cause `log.Println` to be called instead of a panic.
- Errors now cause `t.Error` instead of `t.Fatal`.


0.1.0 - 2016-10-29
------------------

### Added
- Initial public release under the MIT license.


[Unreleased]: https://github.com/nochso/golden/compare/0.2.0...HEAD
[0.2.0]: https://github.com/nochso/golden/compare/0.1.0...0.2.0