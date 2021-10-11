
# protoc --go_out=plugins=grpc:. *.proto

protoc  \
    --proto_path=${GOPATH}/pkg/mod \
    --proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
    --proto_path=. \
    --govalidators_out=. --go_out=plugins=grpc:.\
    *.proto