
protoc-descriptor:
	protoc -I=vendor --go_out=vendor/ vendor/google/protobuf/descriptor.proto

adlctest:
	adlc verify -I /home/garym/devel/helix/jcms/helix/adl  internal/adl/oneofeach.adl

antlrexec:
	go run main.go adl exec internal/adl/oneofeach.adl

antlrwalk:
	go run main.go adl walk internal/adl/oneofeach.adl