# Notes

## Implement gRPC APIs

- [evans: CI with gRPC testing](https://github.com/ktr0731/evans)
- `evans --host localhost --port 9090 -r repl`

## Implement gRPC gateway: serve both HTTP & gRPC requests

- [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- follow this instructions to generate reverse-proxy using `protoc-gen-grpc-gateway` [with custom annotations](https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file#2-with-custom-annotations)
- Send request to test HTTP gateway server
  - in Postman send `http://localhost:8080/v1/create_user`
  - in Postman send `http://localhost:8080/v1/login_user`
- Send request to evans to test gRPC server
- Parse response format to snake_case or [using proto names in JSON](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/#using-proto-names-in-json)
