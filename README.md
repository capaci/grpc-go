

# Generating protobufs

> protoc --go_out=. --go-grpc_out=. proto/course_category.proto


This will generate the files `.pb.go` within `internal/pb` directory (defined as the package in the .proto file)