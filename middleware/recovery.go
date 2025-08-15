package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devlorvn/go-project/common"
)

func Recovery() func(*gin.Context) {
	return func(ctx *gin.Context) {
		log.Println("Recovery")

		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))

				}

				panic(r)
			}
		}()

		ctx.Next()
	}
}
