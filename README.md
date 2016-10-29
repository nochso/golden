golden
======

[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md)
[![GitHub release](https://img.shields.io/github/release/nochso/golden.svg)](https://github.com/nochso/golden/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/golden)](https://goreportcard.com/report/github.com/nochso/golden)
[![Build Status](https://travis-ci.org/nochso/golden.svg?branch=master)](https://travis-ci.org/nochso/golden)
[![Coverage Status](https://coveralls.io/repos/github/nochso/golden/badge.svg?branch=master)](https://coveralls.io/github/nochso/golden?branch=master)

Package golden helps writing golden tests.

A golden test (a.k.a. gold master or approval test) consists of pairs of
files: given test input and expected output (the gold master).

This package is mainly intended to be used with the standard Go testing
library. You can still use it without. It is not intended to be a full
testing framework.


Installation
------------

    go get github.com/nochso/golden


Documentation
-------------

See [godoc](https://godoc.org/github.com/nochso/golden) for API docs and
examples.


License
-------

This package is released under the [MIT license](LICENSE.md).