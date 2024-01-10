package api

// import (
// 	"database/sql"
// 	"errors"
// 	"net/http"

// 	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
// 	"github.com/chomchaby/shopchaby-backend-go/token"
// 	"github.com/google/uuid"
// 	"github.com/lib/pq"

// 	"github.com/gin-gonic/gin"
// )

// type createFavoriteRequest struct {
// 	ProductID string `uri:"id" binding:"required"`
// }

// func (server *Server) createFavorite(ctx *gin.Context) {
// 	var req createFavoriteRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	id, err := uuid.Parse(req.ProductID)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	if store.UserEmail != authPayload.Email {
// 		err := errors.New("store doesn't belong to the authenticated user")
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateFavoriteParams{
// 		UserEmail: req.UserEmail,
// 		ProductID: req.ProductID,
// 	}

// 	favorite, err := server.storeTx.CreateFavorite(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, favorite)

// }

// type listFavoriteByUserRequest struct {
// 	UserEmail string `uri:"email" binding:"required"`
// }

// func (server *Server) listFavoriteByUser(ctx *gin.Context) {
// 	var req listFavoriteByUserRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	favorite, err := server.storeTx.ListFavoritesByUser(ctx, req.UserEmail)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, favorite)

// }

// type deleteFavoriteRequest struct {
// 	UserEmail string    `uri:"email" binding:"required"`
// 	ProductID uuid.UUID `json:"product_id" binding:"required"`
// }

// func (server *Server) deleteFavorite(ctx *gin.Context) {
// 	var req deleteFavoriteRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.GetFavoriteParams{
// 		UserEmail: req.UserEmail,
// 		ProductID: req.ProductID,
// 	}

// 	_, err := server.storeTx.GetFavorite(ctx, arg)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	arg2 := db.DeleteFavoriteParams{
// 		UserEmail: req.UserEmail,
// 		ProductID: req.ProductID,
// 	}
// 	err = server.storeTx.DeleteFavorite(ctx, arg2)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})

// }
