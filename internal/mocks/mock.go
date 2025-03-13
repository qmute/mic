package mocks

//go:generate mockgen -destination=mock_service.go -package=mocks go-micro.dev/v5 Service
//go:generate mockgen -destination=mock_client.go -package=mocks go-micro.dev/v5/client Client,Message
//go:generate mockgen -destination=./mserver/mock_server.go -package=mserver go-micro.dev/v5/server Server,Request,Message
