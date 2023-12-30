module github.com/Bubotka/Microservices/auth

go 1.19

require (
	github.com/Bubotka/Microservices/proxy v0.0.0
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace (
	github.com/Bubotka/Microservices/proxy v0.0.0 => ../proxy
	github.com/Bubotka/Microservices/user v0.0.0 => ../user
    github.com/Bubotka/Microservices/geo v0.0.0 => ../geo
)