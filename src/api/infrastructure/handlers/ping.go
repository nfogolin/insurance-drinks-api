package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/insurance-drinks-api/src/api/infrastructure/utils"
)

type Ping struct {

}

func (h Ping) Handle(c *gin.Context) {
	utils.HandleResponseOk(c, "pong")
}
