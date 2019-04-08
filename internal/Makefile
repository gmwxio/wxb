

goprotoc:
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/test/simple.proto 
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/build/build.proto 
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/workspace/workspace.proto 
	protoc -I=schemas --go_out=paths=source_relative:schemas/ schemas/semantics/semantics.proto 

test:
	protoc -I=schemas -oschemas/test/emptish.ds schemas/test/emptish.proto
	protoc -I=schemas -I../vendor -oschemas/test/simple.ds schemas/test/simple.proto
	protoc -I=schemas -oschemas/test/ds_rpc.ds schemas/test/rpc.proto
	protoc -I=schemas -oschemas/test/ds_fieldopt.ds schemas/test/fieldopt.proto
	protoc -I=schemas -oschemas/test/ds_extends.ds schemas/test/extends.proto
