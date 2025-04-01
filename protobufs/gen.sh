protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative content.proto
cp content.pb.go ../src/gateway/content_grpc/content.pb.go
cp content.pb.go ../src/content/content_grpc/content.pb.go
cp content_grpc.pb.go ../src/gateway/content_grpc/content_grpc.pb.go
cp content_grpc.pb.go ../src/content/content_grpc/content_grpc.pb.go
