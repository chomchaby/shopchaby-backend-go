package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/token"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type createStoreRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	PicUrl      string `json:"pic_url" binding:"required"`
}

func (server *Server) createStore(ctx *gin.Context) {
	var req createStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateStoreParams{
		UserEmail:   authPayload.Email,
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Phone:       req.Phone,
		PicUrl:      req.PicUrl,
	}

	store, err := server.storeTx.CreateStore(ctx, arg)
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
	ctx.JSON(http.StatusOK, store)

}

type getStoreRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (server *Server) getStore(ctx *gin.Context) {
	var req getStoreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := uuid.Parse(req.ID)
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

	ctx.JSON(http.StatusOK, store)

}

// type listStoreRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func (server *Server) listStore(ctx *gin.Context) {
// 	var req listStoreRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	arg := db.ListStoresParams{
// 		Limit:  req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}

// 	store, err := server.storeTx.ListStores(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, store)
// }

type updateStoreRequest struct {
	ID          string `uri:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	PicUrl      string `json:"pic_url"`
}

func (server *Server) updateStore(ctx *gin.Context) {
	var req updateStoreRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
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

	arg := db.UpdateStoreParams{}
	arg.ID = store.ID
	if req.Name != "" {
		arg.Name = req.Name
	} else {
		arg.Name = store.Name
	}
	if req.Description != "" {
		arg.Description = req.Description
	} else {
		arg.Description = store.Description
	}
	if req.Address != "" {
		arg.Address = req.Address
	} else {
		arg.Address = store.Address
	}
	if req.Phone != "" {
		arg.Phone = req.Phone
	} else {
		arg.Phone = store.Phone
	}
	if req.PicUrl != "" {
		arg.PicUrl = req.PicUrl
	} else {
		arg.PicUrl = store.PicUrl
	}

	store, err = server.storeTx.UpdateStore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, store)

}
