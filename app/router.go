package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	// Products Mapping

	router.GET("/items/:id", dependencies.ItemController.Get)
	router.POST("/item", dependencies.ItemController.Insert)
	router.POST("/items", dependencies.ItemController.QueueItems)
	router.GET("/search/:searchQuery", dependencies.ItemController.GetQuery)

	fmt.Println("Finishing mappings configurations")
}
