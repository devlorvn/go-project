package business

import (
	"context"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/modules/item/model"
)

type ListItemStorage interface {
	ListItem(ctx context.Context,
		filter *model.Filter,
		paging *common.Pagination,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBusiness struct {
	store ListItemStorage
}

func NewListItemBusiness(store ListItemStorage) *listItemBusiness {
	return &listItemBusiness{store: store}
}

func (biz listItemBusiness) ListItem(ctx context.Context,
	filter *model.Filter,
	paging *common.Pagination,
) ([]model.TodoItem, error) {
	items, err := biz.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return items, nil
}
