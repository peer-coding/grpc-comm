1. Gerar contrato (.proto)

   - Pode gerar para diferentes linguagens, apenas mexer nas options.

2. Compilar

   - Comando varia para as linguagens que vai gerar.

   ```bash
   	find ./api/proto/contract -name *.proto -print0 | xargs -0 protoc --proto_path=./api/proto/contract --go_out=./api/proto/pb --go_opt=paths=source_relative --go-grpc_out=./api/proto/pb --go-grpc_opt=paths=source_relative
   ```

3. Implementar os arquivos gerados no pb (se usar gRPC), se usar apenas o contrato protobuf, já está pronto.
