package storage

import (
	"context"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/modules/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
