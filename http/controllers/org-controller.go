package controllers

import (
	"context"
	"find-a-friend/db"
	"find-a-friend/services"
	"find-a-friend/services/factories"
	"find-a-friend/types"
	"find-a-friend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrgControler struct {
	orgService *services.OrgService
}

func NewOrgController(router *gin.Engine) *OrgControler {
	s := factories.MakeOrgService(db.DB)
	c := &OrgControler{orgService: s}

	router.POST("/org/create", c.Create)
	// TODO: add login and create token
	router.GET("/org/:id", c.GetById)

	return c
}

func (c *OrgControler) Create(ctx *gin.Context) {
	var org types.CreateOrg
	err := ctx.ShouldBindJSON(&org)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing data: " + err.Error()})
		return
	}

	createdOrg, err := c.orgService.Create(context.Background(), org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating org: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "org created", "org": createdOrg})
}

func (c *OrgControler) GetById(ctx *gin.Context) {
	id, err := utils.ConvertStringToInt(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing id: " + err.Error()})
		return
	}

	org, err := c.orgService.GetById(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating org: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "org found", "org": org})
}