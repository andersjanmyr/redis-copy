# redis-copy

Copies one Redis database to anther for Redis versions that don't support
`MIGRATE` and `SAVE`.


## Usage

```
# redis-copy from to

# Example
$ redis-copy localhost:6379 my-redis-in-the-cloud:6379

```
