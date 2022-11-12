package controllers

import (
	"net/http"

	"github.com/aaraya0/arq-software/arq-sw-2/dtos"
	services "github.com/aaraya0/arq-software/arq-sw-2/services"

	e "github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service services.Service
}

func NewController(service services.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	item, apiErr := ctrl.service.Get(c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) Insert(c *gin.Context) {
	var item dtos.ItemDTO
	if err := c.BindJSON(&item); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := ctrl.service.Insert(item)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, item)
	return
}

func (ctrl *Controller) GetQuery(c *gin.Context) {
	var itemsDto dtos.ItemsDTO
	query := c.Param("searchQuery")

	itemsDto, err := ctrl.service.GetQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, itemsDto)
	}

	c.JSON(http.StatusOK, itemsDto)

}
