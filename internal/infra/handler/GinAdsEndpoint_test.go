package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/ferminhg/learning-go/internal/domain"
	"github.com/ferminhg/learning-go/internal/infra/eventHandler"
	"github.com/ferminhg/learning-go/internal/infra/generator"
	"github.com/ferminhg/learning-go/internal/infra/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_PostNewAd(t *testing.T) {
	adRepository := new(storagemocks.AdServiceRepository)
	adRepository.On("Save", mock.Anything).Return(nil)
	sp := eventHandler.NewMockEventHandler(t)
	service := application.NewAdService(
		adRepository,
		generator.New(false),
		sp,
	)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/ads", PostNewAdsEndpoint(service))

	t.Run("given a invalid request it returns 400", func(t *testing.T) {
		request := PostNewAdsRequest{
			Title: "t1",
			Price: 1,
		}

		b, err := json.Marshal(request)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/ads", bytes.NewBuffer(b))
		assertStatusCodeAndMessage(t, r, req, http.StatusBadRequest, "")
	})

	t.Run("given a valid request it return 201", func(t *testing.T) {
		sp.MockSP.ExpectSendMessageAndSucceed()
		request := PostNewAdsRequest{
			Title:       "t1",
			Description: "d1",
			Price:       15,
		}

		b, err := json.Marshal(request)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/ads", bytes.NewBuffer(b))
		assertStatusCodeAndMessage(t, r, req, http.StatusCreated, "Ad is valid 🎊")
	})
}

func TestHandler_FindById(t *testing.T) {
	adRepository := new(storagemocks.AdServiceRepository)

	service := application.NewAdService(
		adRepository,
		generator.New(false),
		eventHandler.NewMockEventHandler(t),
	)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ads/:id", GetAdByIdEndpoint(service))

	t.Run("given an invalid id, it returns 404", func(t *testing.T) {
		NotFoundAdId, _ := uuid.NewRandom()
		adRepository.On("Find", mock.Anything).Return(domain.Ad{}, false)
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/ads/%s", NotFoundAdId), nil)
		assertStatusCodeAndMessage(t, r, req, http.StatusNotFound, "")
		adRepository.On("Find").Unset()
	})

	t.Run("given a valid id, it returns the Ad", func(t *testing.T) {
		ad := domain.RandomAdFactory()
		adRepository.On("Find", mock.Anything).Return(ad, true)
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/ads/%s", ad.Id), nil)
		assertStatusCodeAndMessage(t, r, req, http.StatusOK, "")
		adRepository.On("Find").Unset()
	})
}

func TestHandler_GetAds(t *testing.T) {
	adRepository := new(storagemocks.AdServiceRepository)
	adRepository.On("Search", 5).Return([]domain.Ad{domain.RandomAdFactory()}, nil)

	service := application.NewAdService(
		adRepository,
		generator.New(false),
		eventHandler.NewMockEventHandler(t),
	)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/ads", GetAdsEndpoint(service))

	t.Run("it return a Ad list empty", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/ads", nil)
		assertStatusCodeAndMessage(t, r, req, http.StatusOK, "")
	})
}

func TestHandler_DeleteAd(t *testing.T) {
	repository := new(storagemocks.AdServiceRepository)
	service := application.NewAdService(
		repository,
		generator.New(false),
		eventHandler.NewMockEventHandler(t),
	)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.DELETE("/ads/:id", DeleteAdByIdHandler(service))

	t.Run("given a not valid AdId it returns 400", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/ads/notvalidadid", nil)
		assertStatusCodeAndMessage(t, r, req, http.StatusBadRequest, "invalid UUID")
	})

	t.Run("given an Id when does not exist then return 404", func(t *testing.T) {
		randomUUID, _ := uuid.NewRandom()
		repository.On("Delete", mock.Anything).Return(false)
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/ads/%s", randomUUID), nil)
		assertStatusCodeAndMessage(t, r, req, http.StatusNotFound, "")
		repository.On("Delete", mock.Anything).Unset()
	})

	t.Run("given an Id when exits then return 200", func(t *testing.T) {
		randomUUID, _ := uuid.NewRandom()
		repository.On("Delete", mock.Anything).Return(true)
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/ads/%s", randomUUID), nil)
		assertStatusCodeAndMessage(t, r, req, http.StatusOK, "")
	})
}

func assertStatusCodeAndMessage(t *testing.T, r *gin.Engine, req *http.Request, expectedStatusCode int, expectedMessage string) {
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	require.Equal(t, expectedStatusCode, res.StatusCode)

	if len(expectedMessage) > 0 {
		var data map[string]string
		require.NoError(t, json.NewDecoder(res.Body).Decode(&data))
		assert.Contains(t, data["message"], expectedMessage)
	}
}
