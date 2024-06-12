package controllers

import (
	"context"
	"find-a-friend/db"
	"find-a-friend/services"
	"find-a-friend/services/factories"
	"find-a-friend/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PetController struct {
	petService *services.PetService
}

func NewPetController(router *gin.Engine) *PetController {
	s := factories.MakePetService(db.DB)
	c := &PetController{petService: s}

	router.POST("/pet", c.CreatePet)
	router.GET("/pet/:city", c.GetFromCity)

	return c
}

func (c *PetController) CreatePet(ctx *gin.Context) {
	var pet types.CreatePet
	if err := ctx.ShouldBindJSON(&pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing data: " + err.Error()})
		return
	}

	createdPet, err := c.petService.Create(ctx, pet)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating pet: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "pet created", "pet": createdPet})
}

func (c *PetController) GetFromCity(ctx *gin.Context) {
	city := ctx.Param("city")
	pet, err := c.petService.GetFromCity(context.Background(), city)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error getting pet: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "pet found", "pet": pet})
}