Redundant import alias
======================

Checks redundant import aliases of Go code.


## Installation

```
$ go install github.com/nobeans/redundantimportalias@latest
```


## Usage

```sh
redundantimportalias: checks redundant import alias

Usage: redundantimportalias [-flag] [package]

Flags:
  -V    print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -source
        no effect (deprecated)
  -tags string
        no effect (deprecated)
  -trace string
        write trace log to this file
  -v    no effect (deprecated)
```


## Examples

```go
// [Not redundant] Avoiding conflicted package names
import (
        aliased_foo "lib-a/foo"
        "lib-b/foo"
)
```

```go
// [Redundant] Forgotten to remove import alias (or using alias just for a taste)
import (
        aliased_foo "lib-a/foo"
        "lib-b/bar"
)
```
