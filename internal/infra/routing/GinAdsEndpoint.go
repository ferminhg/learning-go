package routing

import (
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/ferminhg/learning-go/internal/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdsController struct {
	service application.AdService
}

func New() AdsController {
	return AdsController{
		service: application.NewAdService(infra.NewInMemoryAdRepository()),
	}
}

func (c AdsController) GetHealthEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "👍",
	})
}

func (c AdsController) PostNewAdsEndpoint(ctx *gin.Context) {
	var ad domain.Ad
	err := ctx.ShouldBind(&ad)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ad, err = c.service.Post(ad.Title, ad.Description, ad.Price)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ad is valid 🎊 " + ad.Id.String(),
	})
}

func (c AdsController) GetAdsEndpoint(ctx *gin.Context) {
	ads, err := c.service.FindRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ads": ads,
	})
}

func (c AdsController) GetAdByIdEndpoint(ctx *gin.Context) {
	adId := ctx.Param("id")
	ad, ok := c.service.Find(adId)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Ad not found: " + adId})
	}

	ctx.JSON(http.StatusOK, ad)
}
