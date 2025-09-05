.PHONY: gen-pb
gen-pb:
	find ./api/proto/contract -name *.proto -print0 | xargs -0 protoc --proto_path=./api/proto/contract --go_out=./api/proto/pb --go_opt=paths=source_relative --go-grpc_out=./api/proto/pb --go-grpc_opt=paths=source_relative

.PHONY: clear-pbs
clear-pbs:
	rm -f ./api/proto/pb/*.go