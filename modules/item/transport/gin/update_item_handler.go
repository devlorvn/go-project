package ginitem

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/modules/item/business"
	"github.com/devlorvn/go-project/modules/item/model"
	"github.com/devlorvn/go-project/modules/item/storage"
)

func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := storage.NewSQLStore(db)
		biz := business.NewUpdateItemBusiness(storage)

		if err := biz.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
