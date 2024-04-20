package service

import (
	"context"
	collectv1 "github.com/MuxiKeStack/be-api/gen/proto/collect/v1"
	"github.com/MuxiKeStack/be-collect/domain"
	"github.com/MuxiKeStack/be-collect/repository"
)

type CollectService interface {
	AddCollection(ctx context.Context, c domain.Collection) error
	RemoveCollection(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) error
	ListCollections(ctx context.Context, uid int64, biz collectv1.Biz, curCollectionId int64, limit int64) ([]domain.Collection, error)
	CountCollections(ctx context.Context, uid int64, biz collectv1.Biz) (int64, error)
	Collected(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) (bool, error)
}

type collectService struct {
	repo repository.CollectRepository
}

func NewCollectService(repo repository.CollectRepository) CollectService {
	return &collectService{repo: repo}
}

func (s *collectService) AddCollection(ctx context.Context, c domain.Collection) error {
	return s.repo.CreateCollection(ctx, c)
}

func (s *collectService) RemoveCollection(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) error {
	return s.repo.DeleteCollection(ctx, uid, biz, bizId)
}

func (s *collectService) ListCollections(ctx context.Context, uid int64, biz collectv1.Biz, curCollectionId int64, limit int64) ([]domain.Collection, error) {
	return s.repo.GetListCollections(ctx, uid, biz, curCollectionId, limit)
}

func (s *collectService) CountCollections(ctx context.Context, uid int64, biz collectv1.Biz) (int64, error) {
	return s.repo.GetCountCollections(ctx, uid, biz)
}

func (s *collectService) Collected(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) (bool, error) {
	return s.repo.Collected(ctx, uid, biz, bizId)
}
