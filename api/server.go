package api

import (
	"fmt"

	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/token"
	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for shopchaby service.
type Server struct {
	config     util.Config
	storeTx    db.StoreTx
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, storeTx db.StoreTx) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		storeTx:    storeTx,
		tokenMaker: tokenMaker}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//user.go
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	// router.GET("/api/logout")

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/:email", server.getUser)
	// router.GET("api/users", server.listUsers)
	authRoutes.PUT("/users/:email", server.updateUser)

	// store.go
	authRoutes.POST("/users/stores", server.createStore)
	authRoutes.GET("/users/stores/:id", server.getStore)
	authRoutes.PUT("/users/stores/:id", server.updateStore)
	authRoutes.POST("/users/stores/products", server.createProductListing)
	authRoutes.PUT("/users/stores/products", server.updateProductListing)

	// home
	// router.GET("/api/home/onsale", server.getProductOnSale)

	// filter

	// search

	// favorite
	// router.POST("/users/:email/favorites", server.createFavorite)
	// router.GET("/users/:email/favorites", server.listFavoriteByUser)
	// router.DELETE("/users/:email/favorites", server.deleteFavorite)

	//cart
	router.POST("/users/:email/carts", server.createCart)
	router.GET("/users/:email/carts", server.listCartByUser)
	router.GET("/users/:email/carts/product", server.getCartProduct)
	router.DELETE("/users/:email/carts", server.deleteCart)

	// purchase

	server.router = router

}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
