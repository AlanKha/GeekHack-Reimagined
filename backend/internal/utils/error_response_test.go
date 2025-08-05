package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the function to test
	RespondWithError(c, http.StatusBadRequest, "Test error message")

	// Assert the HTTP status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert the JSON response body
	var responseBody ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Test error message", responseBody.Message)

	// Assert that the context was aborted
	assert.True(t, c.IsAborted())
}
