package controllers

import (
	"fmt"
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
		fmt.Println(apiErr)
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) Insert(c *gin.Context) {
	var item dtos.ItemDTO
	if err := c.BindJSON(&item); err != nil {
		fmt.Println(err)
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := ctrl.service.Insert(item)
	if apiErr != nil {
		fmt.Println(apiErr)
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
		fmt.Println(err)
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, itemsDto)
	return
}

func (ctrl *Controller) QueueItems(c *gin.Context) {
	var itemsDto dtos.ItemsDTO
	err := c.BindJSON(&itemsDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := ctrl.service.QueueItems(itemsDto)

	// Error Queueing
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemsDto)
	return
}
