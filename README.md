golden
======

[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/nochso/golden.svg)](https://github.com/nochso/golden/releases)
[![GoDoc](https://godoc.org/github.com/nochso/golden?status.svg)](http://godoc.org/github.com/nochso/golden)  
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/golden)](https://goreportcard.com/report/github.com/nochso/golden)
[![Build Status](https://travis-ci.org/nochso/golden.svg?branch=master)](https://travis-ci.org/nochso/golden)
[![Coverage Status](https://coveralls.io/repos/github/nochso/golden/badge.svg?branch=master)](https://coveralls.io/github/nochso/golden?branch=master)

Package golden helps writing golden tests.

A golden test (a.k.a. gold master or approval test) consists of pairs of
files: given test input and expected output (the gold master).

This package is mainly intended to be used with the standard Go testing
library. You can still use it without. It is not intended to be a full
testing framework.

- [Installation](#installation)
- [Documentation](#documentation)
- [Notes](#notes)
- [Changes / Versioning](#changes--versioning)
- [License](#license)


Installation
------------

    go get github.com/nochso/golden


Documentation
-------------

See [godoc](https://godoc.org/github.com/nochso/golden) for API docs and
examples.

Notes
-----
It's important that golden files are easy to read, edit and diff. Usually it's
enough to add a `String() string` method or more detailed `TestString`,
`debugString`, etc. to dump the data you want to compare in a human readable
way.

Any methods that could fail and have access to T (as in `testing.T`)
will call `T.Error(err)`. If T is nil, `log.Println(err)` is called instead.
This mainly concerns Case and File structs and any methods mentioning
failing.

Changes / Versioning
--------------------

All notable changes to this project will be documented in [CHANGELOG.md](CHANGELOG.md)

This project adheres to [Semantic Versioning](http://semver.org/).


License
-------

This package is released under the [MIT license](LICENSE).