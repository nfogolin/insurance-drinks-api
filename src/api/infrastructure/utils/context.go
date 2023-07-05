package utils

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

func GetMockedContext(method, url string, requestBody io.Reader, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)

	c.Request, _ = http.NewRequest(method, url, requestBody)

	return c
}

func GetTargetResponse() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
