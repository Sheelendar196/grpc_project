//make dir: 
mkdir graphql-project
cd graphql-project
//to generate project: 
 go mod init github.com/sheelendar196/go-projects/grpc_project  

// to install proto gen if not installed
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

// generate go files for .proto
 protoc --go_out=. --go-grpc_out=. proto/employee.proto
