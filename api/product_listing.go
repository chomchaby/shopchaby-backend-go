package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createProductListingRequest struct {
	Category    string `json:"category" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	StoreID     string `json:"store_id" binding:"required"`
	Suspend     *bool  `json:"suspend" binding:"required"`
	Subproducts []struct {
		Variation   string `json:"variation" binding:"required"`
		StockAmount int32  `json:"stock_amount" binding:"required"`
		Price       int32  `json:"price" binding:"required"`
		SalePrice   int32  `json:"sale_price"`
	} `json:"subproducts" binding:"required"`
	ProductImages []struct {
		Title     string `json:"title" binding:"required"`
		PicUrl    string `json:"pic_url" binding:"required"`
		IsDefault bool   `json:"is_default" binding:"required"`
	} `json:"product_images" binding:"required"`
}

func (server *Server) createProductListing(ctx *gin.Context) {
	var req createProductListingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id, err := uuid.Parse(req.StoreID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	store, err := server.storeTx.GetStore(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if store.UserEmail != authPayload.Email {
		err := errors.New("store doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreateProductListingTxParams{
		Category:      req.Category,
		Name:          req.Name,
		Description:   req.Description,
		StoreID:       store.ID,
		Suspend:       *req.Suspend,
		Subproducts:   req.Subproducts,
		ProductImages: req.ProductImages,
	}

	result, err := server.storeTx.CreateProductListingTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)

}

type updateProductListingRequest struct {
	ProductID   string `json:"product_id" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Suspend     *bool  `json:"suspend" binding:"required"`
	Subproducts []struct {
		Variation   string `json:"variation" binding:"required"`
		StockAmount int32  `json:"stock_amount" binding:"required"`
		Price       int32  `json:"price" binding:"required"`
		SalePrice   int32  `json:"sale_price"`
	} `json:"subproducts" binding:"required"`
	ProductImages []struct {
		Title     string `json:"title" binding:"required"`
		PicUrl    string `json:"pic_url" binding:"required"`
		IsDefault bool   `json:"is_default" binding:"required"`
	} `json:"product_images" binding:"required"`
}

func (server *Server) updateProductListing(ctx *gin.Context) {
	var req updateProductListingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := uuid.Parse(req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.storeTx.GetProduct(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	store, err := server.storeTx.GetStore(ctx, product.StoreID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if store.UserEmail != authPayload.Email {
		err := errors.New("product doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdateProductListingTxParams{
		ProductID:     product.ID,
		Category:      req.Category,
		Name:          req.Name,
		Description:   req.Description,
		Suspend:       *req.Suspend,
		Subproducts:   req.Subproducts,
		ProductImages: req.ProductImages,
	}

	result, err := server.storeTx.UpdateProductListingTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)

}
