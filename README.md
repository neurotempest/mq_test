A playground to test various message queueing frameworks

## Setup notes

### protoc & proto-gen-go

* `protoc`, the Protobuf compiler, can be installed with `brew install protoc`

* `proto-gen-go` and `proto-gen-go-grpc`, to generate golang GRPC protobuf bindings can be installed with `go get` and `go install`

* To setup proto imports set the `PROTOC_INCLUDE_DIR` enviroment variable and copy the approriate proto files into the correct subpaths into that include dir:

  * Under the hood; `--proto_path=$PROTOC_INCLUDE_DIR` is set in the `generate.go` files for the protobuf go bindings (as well `--proto_path=.`; for some reason it also wants the path to the proto file to be generated to also be in the `proto_path`... apparently `.` is the default.)

  * Inside the path in `PROTOC_INCLUDE_DIR`, you need to copy the proto files to be imported with the correct sub paths - the same path as it appears in the import path (e.g. `PROTOC_INCLUDE_DIR/github.com/luno/reflex/reflexpb/reflex.proto`)

  * Unfortunately you can't just set `PROTOC_INCLUDE_DIR` to be equal to the go mod dir (e.g. `~/go/pkg/mod`); while it contains all the right files, it also has version info in the folder names... I guess the go compiler knows how to interpret that extra version info in the path... protoc obviously does not.

  * **Alternative option** the `protoc` executable also seems to look for proto includes in the path `../include` relative from the executable path. You can see that (when installed with `brew` anyway) the path `../include` has all the _standard_ google proto imports (e.g. `google/protobuf/timestamp.proto`)
