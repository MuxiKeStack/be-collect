package repository

import (
	"context"
	collectv1 "github.com/MuxiKeStack/be-api/gen/proto/collect/v1"
	"github.com/MuxiKeStack/be-collect/domain"
	"github.com/MuxiKeStack/be-collect/repository/dao"
	"github.com/ecodeclub/ekit/slice"
)

type CollectRepository interface {
	CreateCollection(ctx context.Context, c domain.Collection) error
	DeleteCollection(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) error
	GetListCollections(ctx context.Context, uid int64, biz collectv1.Biz, curCollectionId int64, limit int64) ([]domain.Collection, error)
	GetCountCollections(ctx context.Context, uid int64, biz collectv1.Biz) (int64, error)
	Collected(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) (bool, error)
}

type collectRepository struct {
	dao dao.CollectDAO
}

func NewCollectRepository(dao dao.CollectDAO) CollectRepository {
	return &collectRepository{dao: dao}
}

func (repo *collectRepository) CreateCollection(ctx context.Context, c domain.Collection) error {
	return repo.dao.Insert(ctx, repo.toEntity(c))
}

func (repo *collectRepository) DeleteCollection(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) error {
	return repo.dao.Delete(ctx, uid, int32(biz), bizId)
}

func (repo *collectRepository) GetListCollections(ctx context.Context, uid int64, biz collectv1.Biz, curCollectionId int64, limit int64) ([]domain.Collection, error) {
	collections, err := repo.dao.FindCollections(ctx, uid, int32(biz), curCollectionId, limit)
	return slice.Map(collections, func(idx int, src dao.Collection) domain.Collection {
		return repo.toDomain(src)
	}), err
}

func (repo *collectRepository) GetCountCollections(ctx context.Context, uid int64, biz collectv1.Biz) (int64, error) {
	return repo.dao.GetCountCollections(ctx, uid, int32(biz))
}

func (repo *collectRepository) Collected(ctx context.Context, uid int64, biz collectv1.Biz, bizId int64) (bool, error) {
	_, err := repo.dao.FindCollection(ctx, uid, int32(biz), bizId)
	switch {
	case err == nil:
		return true, nil
	case err == dao.ErrRecordNotFound:
		return false, nil
	default:
		return false, err
	}
}

func (repo *collectRepository) toEntity(c domain.Collection) dao.Collection {
	return dao.Collection{
		Id:    c.Id,
		Uid:   c.Uid,
		Biz:   int32(c.Biz),
		BizId: c.BizId,
	}
}

func (repo *collectRepository) toDomain(c dao.Collection) domain.Collection {
	return domain.Collection{
		Id:    c.Id,
		Uid:   c.Uid,
		Biz:   collectv1.Biz(c.Biz),
		BizId: c.BizId,
	}
}
