package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lets-learn-it/gin-auth-scope/token"
	"github.com/lets-learn-it/gin-auth-scope/util"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// add routes to router
	router.POST("/login", server.login)

	router.GET("/resource",
		authMiddleware(server.tokenMaker, []string{"resource.get"}),
		server.getResource)

	router.POST("/resource",
		authMiddleware(server.tokenMaker, []string{"resource.post"}),
		server.postResource)

	router.PUT("/resource",
		authMiddleware(server.tokenMaker, []string{"resource.get", "resource.post"}),
		server.putResource)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
