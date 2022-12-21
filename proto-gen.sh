protoc --go_out=./proto-gen \
--go-grpc_out=./proto-gen \
./protos/order.proto

protoc --go_out=./proto-gen \
--go-grpc_out=./proto-gen \
./protos/product.proto

protoc --go_out=./proto-gen \
--go-grpc_out=./proto-gen \
./protos/category.proto

protoc --go_out=./proto-gen \
--go-grpc_out=./proto-gen \
./protos/auth.proto