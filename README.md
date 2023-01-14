

# Generating protobufs

> protoc --go_out=. --go-grpc_out=. proto/course_category.proto --experimental_allow_proto3_optional


This will generate the files `.pb.go` within `internal/pb` directory (defined as the package in the .proto file)



# running evans

> evans -r repl -p 50051
