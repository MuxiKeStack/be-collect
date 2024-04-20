package grpc

import (
	"context"
	collectv1 "github.com/MuxiKeStack/be-api/gen/proto/collect/v1"
	"github.com/MuxiKeStack/be-collect/domain"
	"github.com/MuxiKeStack/be-collect/service"
	"github.com/ecodeclub/ekit/slice"
	"google.golang.org/grpc"
	"math"
)

type CollectServiceServer struct {
	collectv1.UnimplementedCollectServiceServer
	svc service.CollectService
}

func NewCollectServiceServer(svc service.CollectService) *CollectServiceServer {
	return &CollectServiceServer{svc: svc}
}

func (c *CollectServiceServer) Register(server grpc.ServiceRegistrar) {
	collectv1.RegisterCollectServiceServer(server, c)
}

func (c *CollectServiceServer) AddCollection(ctx context.Context, request *collectv1.AddCollectionRequest) (*collectv1.AddCollectionResponse, error) {
	err := c.svc.AddCollection(ctx, convertToDomain(request.GetCollection()))
	return &collectv1.AddCollectionResponse{}, err
}

func (c *CollectServiceServer) RemoveCollection(ctx context.Context, request *collectv1.RemoveCollectionRequest) (*collectv1.RemoveCollectionResponse, error) {
	err := c.svc.RemoveCollection(ctx, request.GetUid(), request.GetBiz(), request.GetBizId())
	return &collectv1.RemoveCollectionResponse{}, err
}

func (c *CollectServiceServer) ListCollections(ctx context.Context, request *collectv1.ListCollectionsRequest) (*collectv1.ListCollectionsResponse, error) {
	curCollectionId := request.GetCurCollectionId()
	if curCollectionId <= 0 {
		curCollectionId = math.MaxInt64
	}
	collections, err := c.svc.ListCollections(ctx, request.GetUid(), request.GetBiz(), curCollectionId, request.GetLimit())
	return &collectv1.ListCollectionsResponse{
		Collections: slice.Map(collections, func(idx int, src domain.Collection) *collectv1.Collection {
			return convertToV(src)
		}),
	}, err
}

func (c *CollectServiceServer) CountCollections(ctx context.Context, request *collectv1.CountCollectionsRequest) (*collectv1.CountCollectionsResponse, error) {
	count, err := c.svc.CountCollections(ctx, request.GetUid(), request.GetBiz())
	return &collectv1.CountCollectionsResponse{TotalCount: count}, err
}

func (c *CollectServiceServer) CheckCollection(ctx context.Context, request *collectv1.CheckCollectionRequest) (*collectv1.CheckCollectionResponse, error) {
	collected, err := c.svc.Collected(ctx, request.GetUid(), request.GetBiz(), request.GetBizId())
	return &collectv1.CheckCollectionResponse{
		IsCollected: collected,
	}, err
}

func convertToDomain(c *collectv1.Collection) domain.Collection {
	return domain.Collection{
		Id:    c.Id,
		Uid:   c.Uid,
		Biz:   c.Biz,
		BizId: c.BizId,
	}
}

func convertToV(c domain.Collection) *collectv1.Collection {
	return &collectv1.Collection{
		Id:    c.Id,
		Uid:   c.Uid,
		Biz:   c.Biz,
		BizId: c.BizId,
	}
}
