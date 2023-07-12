# rfc7807gen

## Generate

```sh
swagger generate server
swagger generate client
```

## Test

Server:

```sh
>go run cmd\swagger-hello-example-server\main.go --port 8081
2023/07/11 16:39:07 Serving swagger hello example at http://127.0.0.1:8081
```

Client

```sh
\rfc7807\cmd\call>go run main.go
test ok
```

Curl

```sh
>curl -v localhost:8081/v1/hello
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081 (#0)
> GET /v1/hello HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.0.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/problem+json
< Date: Wed, 12 Jul 2023 03:03:34 GMT
< Content-Length: 22
<
{"message":"test ok"}
* Connection #0 to host localhost left intact
```