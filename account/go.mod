module github.com/itswcg/micro/account

go 1.12

require (
	github.com/bilibili/kratos v0.2.4-0.20191025092737-e14170de04ba
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/itswcg/micro/leaf v0.0.0-00010101000000-000000000000
	github.com/itswcg/micro/middleware v0.0.0-00010101000000-000000000000
	github.com/prometheus/common v0.6.0
	golang.org/x/net v0.0.0-20191011234655-491137f69257
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03
	google.golang.org/grpc v1.24.0
)

replace (
	github.com/itswcg/micro/auth => ../auth
	github.com/itswcg/micro/leaf => ../leaf
	github.com/itswcg/micro/middleware => ../middleware
)
