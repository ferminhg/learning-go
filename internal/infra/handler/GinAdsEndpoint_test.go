package handler

import (
	"bytes"
	"encoding/json"
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/ferminhg/learning-go/internal/infra/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_PostNewAd(t *testing.T) {
	adRepository := new(storagemocks.AdServiceRepository)
	service := application.NewAdService(adRepository)

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
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it return 201", func(t *testing.T) {

	})
}
