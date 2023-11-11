genauth:
	protoc --go_out=. --go-grpc_out=. --proto_path=./pkg/services/auth-svc/pb auth.proto