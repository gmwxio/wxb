
protoc-descriptor:
	protoc -I=vendor --go_out=vendor/ vendor/google/protobuf/descriptor.proto
