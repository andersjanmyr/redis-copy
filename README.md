# redis-copy

Copies one Redis database to another, for Redis versions that don't support
`MIGRATE` and `SAVE`. It uses DUMP and RESTORE.

## Installation

### Homebrew Installation on OS X

```
$ brew tap andersjanmyr/tap
$ brew install redis-copy
```
`redis-copy` is a single binary. Install it by right-clicking and `Save as...`
or with `curl`.

### Links

* [OS X](https://github.com/andersjanmyr/redis-copy/releases/download/v1.0.2/redis-copy-osx)
* [Linux](https://github.com/andersjanmyr/redis-copy/releases/download/v1.0.2/redis-copy-linux)
* [Windows](https://github.com/andersjanmyr/redis-copy/releases/download/v1.0.2/redis-copy.exe)

### Curl

```
# OS X
$ curl -L https://github.com/andersjanmyr/redis-copy/releases/download/v1.0.2/redis-copy-osx \
  > /usr/local/bin/redis-copy

# Linux
$ curl -L https://github.com/andersjanmyr/redis-copy/releases/download/v1.0.2/redis-copy-linux \
  > /usr/local/bin/redis-copy

# Make executable
$ chmod a+x /usr/local/bin/redis-copy

```



## Usage

```
$ ./redis-copy --help
Usage: redis-copy [options] <from> <to>
  -force=false: Overwrite existing keys
  -help=false: Show help text
  -verbose=false: Verbose output

# Example
$ redis-copy localhost:6379 my-redis-in-the-cloud:6379

```
