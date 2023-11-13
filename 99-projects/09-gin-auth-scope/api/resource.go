package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getResource(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get resource")
}

func (server *Server) putResource(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "put resource")
}

func (server *Server) postResource(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "post resource")
}
