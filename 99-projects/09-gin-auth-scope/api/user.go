package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginUserRequest struct {
	Username string   `json:"username" binding:"required,alphanum"`
	Password string   `json:"password" binding:"required,min=6"`
	Scopes   []string `json:"scopes"` // taking scopes from user just for simplicity
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}

func (s *Server) login(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accessToken, err := s.tokenMaker.CreateToken(
		req.Username,
		s.config.AccessTokenDuration,
		req.Scopes,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
	}
	ctx.JSON(http.StatusOK, rsp)
}
