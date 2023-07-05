package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleResponseOk(c *gin.Context, json any) {
	c.JSON(http.StatusOK, json)
}

func HandleResponseBadRequest(c *gin.Context, json any) {
	c.JSON(http.StatusBadRequest, json)
}

func HandleResponseNotFound(c *gin.Context, json any) {
	c.JSON(http.StatusNotFound, json)
}
