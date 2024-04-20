//go:build wireinject

package main

import (
	"github.com/MuxiKeStack/be-collect/grpc"
	"github.com/MuxiKeStack/be-collect/ioc"
	"github.com/MuxiKeStack/be-collect/pkg/grpcx"
	"github.com/MuxiKeStack/be-collect/repository"
	"github.com/MuxiKeStack/be-collect/repository/dao"
	"github.com/MuxiKeStack/be-collect/service"
	"github.com/google/wire"
)

func InitGRPCServer() grpcx.Server {
	wire.Build(
		ioc.InitGRPCxKratosServer,
		grpc.NewCollectServiceServer,
		service.NewCollectService,
		repository.NewCollectRepository,
		dao.NewGORMCollectDAO,
		ioc.InitDB,
		ioc.InitEtcdClient,
		ioc.InitLogger,
	)
	return grpcx.Server(nil)
}
