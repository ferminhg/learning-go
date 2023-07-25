package handler

import (
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealthEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "👍",
		})
	}
}

type PostNewAdsRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
}

func PostNewAdsEndpoint(service application.AdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req PostNewAdsRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ad, err := service.Post(req.Title, req.Description, req.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Ad is valid 🎊 " + ad.Id.String(),
		})
	}
}

func GetAdsEndpoint(service application.AdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ads, err := service.FindRandom()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ads": ads,
		})
	}
}

func GetAdByIdEndpoint(service application.AdService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adId := ctx.Param("id")
		ad, ok := service.Find(adId)
		if !ok {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Ad not found: " + adId})
		}

		ctx.JSON(http.StatusOK, ad)
	}
}
