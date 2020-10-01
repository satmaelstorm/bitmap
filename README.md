# Go Simple Bitmap
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