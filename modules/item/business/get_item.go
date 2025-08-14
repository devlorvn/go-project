package business

import (
	"context"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/modules/item/model"

)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store: store}
}

func (biz getItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	item, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}
	return item, nil
}
