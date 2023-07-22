# 02_partitioned

Creates an HTTP server that can handle `PUT`, `GET`, and `DELETE` requests of key value pairs. The key value pairs are stored across partitioned key value workers.

I am currently assuming a constant number of immortal key value workers.

## Usage

In one terminal session start the server:

```shell
$ cd main
$ go run main.go
```

In other terminal sessions issue requests:

`PUT`

```shell
$ curl 'localhost:3000/put?key=test_key&value=test_value'
```

`GET`

```shell
$ curl 'localhost:3000/get?key=test_key'
```

`DELETE`

```shell
$ curl 'localhost:3000/delete?key=test_key'
```
