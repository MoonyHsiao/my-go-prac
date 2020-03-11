package errno

import (
	"github.com/MoonyHsiao/my-go-prac/viewmodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, DefaultSuccess(data))
}

func DefaultSuccess(data interface{}) viewmodels.APIResult {
	return viewmodels.APIResult{
		Success: true,
		Code:    0,
		Data:    data,
	}
}
