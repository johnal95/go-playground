package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnal95/go-playground/model"
)

func GetExampleV1(ctx *gin.Context) {
	response := model.NewExample("Hello World!")

	ctx.JSON(http.StatusOK, response)
}
