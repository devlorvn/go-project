package ginitem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/modules/item/business"
	"github.com/devlorvn/go-project/modules/item/model"
	"github.com/devlorvn/go-project/modules/item/storage"
)

func ListItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Pagination
		var filter model.Filter

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		storage := storage.NewSQLStore(db)
		biz := business.NewListItemBusiness(storage)

		data, err := biz.ListItem(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging))
	}
}
