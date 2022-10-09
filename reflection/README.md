# Reflection

grpc-go server reflection is implemented in package [grpc-go/reflection at master · grpc/grpc-go (github.com)](https://github.com/grpc/grpc-go/tree/master/reflection).
To enable it, we need to install this package and register server reflection service on gRPC server.

After registering reflection service, we can also use [grpcurl](https://github.com/fullstorydev/grpcurl) to talk to server.

## grpcurl Examples

```bash
grpcurl -plaintext localhost:8080 list
grpcurl -plaintext localhost:8080 describe reflection.HelloService.SayHello
grpcurl -plaintext -d '{"name": "Tony Huang"}' localhost:8080 reflection.HelloService.SayHello | jq .
```
