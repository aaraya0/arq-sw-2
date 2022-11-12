package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	// Products Mapping

	router.GET("/items/:id", dependencies.ItemController.Get)
	router.POST("/items", dependencies.ItemController.Insert)
	router.GET("/search/:searchQuery", dependencies.ItemController.GetQuery)

	fmt.Println("Finishing mappings configurations")
}
