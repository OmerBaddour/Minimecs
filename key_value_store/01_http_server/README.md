# 01_http_server

Creates an HTTP server that can handle `GET`, `PUT` and `DELETE` requests of key value pairs on a concurrency-safe map.

## Usage

In one terminal session start the server:

```shell
$ go run main.go
```

In other terminal sessions issue requests:

### `PUT`

```shell
$ curl 'localhost:3000/put?key=test_key&value=test_value'
```

### `GET`

```shell
$ curl 'localhost:3000/get?key=test_key'
```

### `DELETE`

```shell
$ curl 'localhost:3000/delete?key=test_key'
```

## Final remarks

Note: the concurrency-safety of the map [might be irrelevant to my use-case](https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c). I'll revisit this in the next iteration.
