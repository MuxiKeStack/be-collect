package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type CollectDAO interface {
	Insert(ctx context.Context, c Collection) error
	Delete(ctx context.Context, uid int64, biz int32, bizId int64) error
	FindCollections(ctx context.Context, uid int64, biz int32, curCollectionId int64, limit int64) ([]Collection, error)
	GetCountCollections(ctx context.Context, uid int64, biz int32) (int64, error)
	FindCollection(ctx context.Context, uid int64, biz int32, bizId int64) (Collection, error)
}

type GORMCollectDAO struct {
	db *gorm.DB
}

func NewGORMCollectDAO(db *gorm.DB) CollectDAO {
	return &GORMCollectDAO{db: db}
}

// insert没有返回id，因为用不到
func (dao *GORMCollectDAO) Insert(ctx context.Context, c Collection) error {
	now := time.Now().UnixMilli()
	c.Utime = now
	c.Ctime = now
	return dao.db.WithContext(ctx).Create(&c).Error
}

func (dao *GORMCollectDAO) Delete(ctx context.Context, uid int64, biz int32, bizId int64) error {
	res := dao.db.WithContext(ctx).
		Where("uid = ? and biz = ? and biz_id = ?", uid, biz, bizId).
		Delete(&Collection{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
}

func (dao *GORMCollectDAO) FindCollections(ctx context.Context, uid int64, biz int32, curCollectionId int64, limit int64) ([]Collection, error) {
	var collections []Collection
	err := dao.db.WithContext(ctx).
		Where("uid = ? and biz = ? and id < ?", uid, biz, curCollectionId).
		Order("id desc").
		Limit(int(limit)).
		Find(&collections).Error
	return collections, err
}

// 这是直接聚合查询得到的count，没有额外像评论数量那样维护一个count，因为我认为评论数是查多写少，这里并不符合
func (dao *GORMCollectDAO) GetCountCollections(ctx context.Context, uid int64, biz int32) (int64, error) {
	var count int64
	err := dao.db.WithContext(ctx).
		Model(&Collection{}).
		Where("uid = ? and biz = ?", uid, biz).
		Count(&count).Error
	return count, err
}

func (dao *GORMCollectDAO) FindCollection(ctx context.Context, uid int64, biz int32, bizId int64) (Collection, error) {
	var c Collection
	err := dao.db.WithContext(ctx).
		Where("uid = ? and biz = ? and biz_id = ?", uid, biz, bizId).
		First(&c).Error
	return c, err
}

type Collection struct {
	Id    int64 `gorm:"primaryKey,autoIncrement"`
	Uid   int64 `gorm:"uniqueIndex:id_biz_bizId"`
	Biz   int32 `gorm:"uniqueIndex:id_biz_bizId"`
	BizId int64 `gorm:"uniqueIndex:id_biz_bizId"`
	Utime int64
	Ctime int64
}
