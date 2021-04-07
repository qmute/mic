package mocks

//go:generate mockgen -destination=mock_service.go -package=mocks github.com/micro/go-micro/v2 Service
//go:generate mockgen -destination=mock_client.go -package=mocks github.com/micro/go-micro/v2/client Client,Message
//go:generate mockgen -destination=./mserver/mock_server.go -package=mserver github.com/micro/go-micro/v2/server Server,Request,Message
