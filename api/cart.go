package api

import (
	"database/sql"
	"net/http"

	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type createCartRequest struct {
	UserEmail    string    `uri:"email" binding:"required"`
	SubproductID uuid.UUID `json:"subproduct_id" binding:"required"`
	Quantity     int32     `json:"quantity" binding:"required"`
}

func (server *Server) createCart(ctx *gin.Context) {
	var req createCartRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCartParams{
		UserEmail:    req.UserEmail,
		SubproductID: req.SubproductID,
	}
	arg2 := db.CreateCartParams{
		UserEmail:    req.UserEmail,
		SubproductID: req.SubproductID,
		Quantity:     req.Quantity,
	}

	_, err := server.storeTx.GetCart(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			cart, err := server.storeTx.CreateCart(ctx, arg2)
			if err != nil {
				if pqErr, ok := err.(*pq.Error); ok {
					switch pqErr.Code.Name() {
					case "foreign_key_violation", "unique_violation":
						ctx.JSON(http.StatusForbidden, errorResponse(err))
						return
					}
				}
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, cart)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg3 := db.UpdateCartParams{
		UserEmail:    req.UserEmail,
		SubproductID: req.SubproductID,
		Quantity:     req.Quantity,
	}

	cart, err := server.storeTx.UpdateCart(ctx, arg3)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

type listCartByUserRequest struct {
	UserEmail string `uri:"email" binding:"required"`
}

func (server *Server) listCartByUser(ctx *gin.Context) {
	var req listCartByUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	cart, err := server.storeTx.ListCartsByUser(ctx, req.UserEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cart)

}

type getCartProductRequest struct {
	UserEmail    string    `uri:"email" binding:"required"`
	SubproductID uuid.UUID `json:"subproduct_id" binding:"required"`
}

func (server *Server) getCartProduct(ctx *gin.Context) {
	var req getCartProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//////////// continue here

}

type deleteCartRequest struct {
	UserEmail    string    `uri:"email" binding:"required"`
	SubproductID uuid.UUID `json:"product_id" binding:"required"`
}

func (server *Server) deleteCart(ctx *gin.Context) {
	var req deleteCartRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCartParams{
		UserEmail:    req.UserEmail,
		SubproductID: req.SubproductID,
	}

	_, err := server.storeTx.GetCart(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg2 := db.DeleteCartParams{
		UserEmail:    req.UserEmail,
		SubproductID: req.SubproductID,
	}
	err = server.storeTx.DeleteCart(ctx, arg2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})

}
