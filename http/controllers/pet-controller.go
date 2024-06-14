package controllers

import (
	"context"
	"find-a-friend/db"
	auth "find-a-friend/http/middlewares"
	"find-a-friend/services"
	"find-a-friend/services/factories"
	"find-a-friend/types"
	"find-a-friend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PetController struct {
	petService *services.PetService
}

func NewPetController(router *gin.Engine) *PetController {
	s := factories.MakePetService(db.DB)
	c := &PetController{petService: s}

	router.POST("/pet", auth.Authenticate, c.Create) // org needs to be logged in to register a pet
	router.GET("/pet", c.GetFromCity)
	router.GET("/pet/:id", c.GetById)

	return c
}

func (c *PetController) Create(ctx *gin.Context) {
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

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "pet": createdPet})
}

func (c *PetController) GetFromCity(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	filter := types.PetFilter{City: city}

	if species := ctx.Query("species"); species != "" {
		filter.Species = &species
	}
	if breed := ctx.Query("breed"); breed != "" {
		filter.Breed = &breed
	}
	if height := ctx.Query("height"); height != "" {
		heightInt, err := utils.ConvertStringToFloat(height)
		if err == nil {
			filter.Height = &heightInt
		}
	}
	if weight := ctx.Query("weight"); weight != "" {
		weightInt, err := utils.ConvertStringToFloat(weight)
		if err == nil {
			filter.Weight = &weightInt
		}
	}

	pets, err := c.petService.GetFromCity(context.Background(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error getting pets: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "pets": pets})
}


func (c *PetController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := utils.ConvertStringToInt(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing data: " + err.Error()})
		return
	}

	pet, err := c.petService.GetById(context.Background(), intId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error getting pet: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "pet": pet})
}