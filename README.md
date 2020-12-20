# Go Simple Bitmap

[![Go Report](https://goreportcard.com/badge/github.com/satmaelstorm/bitmap)](https://goreportcard.com/report/github.com/satmaelstorm/bitmap) 
[![GoDoc](https://godoc.org/github.com/satmaelstorm/bitmap?status.svg)](http://godoc.org/github.com/satmaelstorm/bitmap)
[![Coverage Status](https://coveralls.io/repos/github/satmaelstorm/bitmap/badge.svg?branch=master)](https://coveralls.io/github/satmaelstorm/bitmap?branch=master) 
![Go](https://github.com/satmaelstorm/bitmap/workflows/Go/badge.svg)

Implements bitmap index functions.
## Install

`go get github.com/satmaelstorm/bitmap`

## Types

### Unlimited
Implements an unlimited bitmap without any compression. Immutable and thread-safe.

### Index64
Bitmap, which can contain numbers from 0 to 63.
Immutable. Thread-safe.

### Atomic64
Bitmap, which can contain numbers from 0 to 63.
Thread-safe.