package storage

import (
	"context"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/modules/item/model"
)

func (s *sqlStore) ListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Pagination,
	moreKeys ...string,
) ([]model.TodoItem, error) {
	var data []model.TodoItem

	db := s.db

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
