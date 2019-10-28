module github.com/itswcg/micro/middleware

go 1.12

require (
	github.com/bilibili/kratos v0.2.4-0.20191025092737-e14170de04ba
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/itswcg/micro/account v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.0.0-20191011234655-491137f69257
	google.golang.org/grpc v1.24.0
)

replace github.com/itswcg/micro/account => ../account
