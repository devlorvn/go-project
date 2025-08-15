package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/devlorvn/go-project/common"
	"github.com/devlorvn/go-project/middleware"
	ginitem "github.com/devlorvn/go-project/modules/item/transport/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(middleware.Recovery())

	r.GET("/ping", func(c *gin.Context) {

		go func() { // ví dụ recovery khi tạo goroutine
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	r.Run(":3000")
}
