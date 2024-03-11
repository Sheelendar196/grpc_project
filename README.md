//Intalll kafka and zooker 

// Run zookeeper
zookeeper-server-start  /opt/homebrew/etc/kafka/zookeeper.properties

//Run kafka
 kafka-server-start   /opt/homebrew/etc/kafka/server.properties

// create topic if not there
  kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic employee_data

// Start consumer on terminla
 kafka-console-consumer --bootstrap-server localhost:9092 --topic employee_data --from-beginning

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

// install these binary files
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 

//run this proto command
protoc ./proto/employee.proto \
            --go_out=. \
            --go_opt=paths=source_relative \
            --go-grpc_out=. \
            --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out . \
            --grpc-gateway_opt logtostderr=true \
            --grpc-gateway_opt paths=source_relative \
            --grpc-gateway_opt generate_unbound_methods=true

// now you will see one more file with (pb.gw.go)
// create new main file to register method of (pb.gw.go) -> test_grpc_proxy/main.go

func main() {
	var grpcServerEndpoint = "localhost:8010"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := proto.RegisterEmployeeServiceHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Listening on port 8011")
	port := ":8011"
	http.ListenAndServe(port, mux)
}

// start grpc service to accept connection
go run cmd/grpc/main.go

// start test gateway service to run rest api
go run test_grpc_proxy/main.go


// Now you can call API from terninal or postman:

//GetEmployee request using post method OR you can also call from postman
curl -X POST -s localhost:8011/grpc_project.EmployeeService/GetEmployee -d '{ "empID": "232" }' 
curl -X POST -s -d '{ "empID": "44543" }' localhost:8011/grpc_project.EmployeeService/GetEmployee 

// CreateEmployee request you can run form Post main as well
curl -X POST -s localhost:8011/grpc_project.EmployeeService/CreateEmployee -d '{"employee": {"empID": "44543","name": "ram","email": "ram@gmail.com","mobile": "8733438792", "department": "IT","address": "bangalore"}}'


grpcurl -insecure -d '{ "empID": "44543" }' localhost:8011 grpc_project.EmployeeService/GetEmployee 
grpcurl -insecure -d '{ "id": 1 }' localhost:8080 api.v1.Activity_Log/Retrieve 


// Generate swagger file run this protoc command then it will generate swagger file as well.
protoc ./proto/employee.proto \
            --go_out=. \
            --go_opt=paths=source_relative \
            --go-grpc_out=. \
            --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out . \
            --grpc-gateway_opt logtostderr=true \
            --grpc-gateway_opt paths=source_relative \
            --grpc-gateway_opt generate_unbound_methods=true \
            --openapiv2_out . \
            --openapiv2_opt logtostderr=true \
            --openapiv2_opt generate_unbound_methods=true

// Copy paste this swagger file into url https://editor.swagger.io/ and can you can see request and resposne.

// for more detail check this web link
https://earthly.dev/blog/golang-grpc-gateway/