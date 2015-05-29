# redis-copy

Copies one Redis database to anther for Redis versions that don't support
`MIGRATE` and `SAVE`. It uses DUMP and RESTORE.


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
