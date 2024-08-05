## Regenerate gRPC code

```sh
make
```

## Runing

First Step: run grpc server

```sh
go run grpc/server.go
```

Second Step: run gin server

```sh
go run gin/main.go
```

## Testing

Send data to gin server:

```sh
curl -v 'http://localhost:8080/rest/n/gin'
```

or using [grpcurl](https://github.com/fullstorydev/grpcurl) command:

```sh
sudo docker run --rm --network="host" fullstorydev/grpcurl -d '{"name": "gin"}' -plaintext localhost:50051 helloworld.v1.Greeter/SayHello
```