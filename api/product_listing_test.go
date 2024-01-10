package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/chomchaby/shopchaby-backend-go/db/mock"
	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/token"
	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchProduct(t *testing.T, responseBody []byte, product db.Product) {

	var gotListing db.ProductListingTxResult
	err := json.Unmarshal(responseBody, &gotListing)
	require.NoError(t, err)
	require.Equal(t, product.Category, gotListing.Product.Category)
	require.Equal(t, product.Name, gotListing.Product.Name)
	require.Equal(t, product.Description, gotListing.Product.Description)
	require.Equal(t, product.StoreID, gotListing.Product.StoreID)
	require.Equal(t, product.Suspend, gotListing.Product.Suspend)
	require.Equal(t, product.Vendible, gotListing.Product.Vendible)

}

func requireBodyMatchSubproduct(t *testing.T, responseBody []byte, subproduct db.Subproduct, ind int) {

	var gotListing db.ProductListingTxResult
	err := json.Unmarshal(responseBody, &gotListing)
	require.NoError(t, err)
	require.Equal(t, subproduct.Variation, gotListing.Subproducts[ind].Variation)
	require.Equal(t, subproduct.StockAmount, gotListing.Subproducts[ind].StockAmount)
	require.Equal(t, subproduct.Price, gotListing.Subproducts[ind].Price)
	require.Equal(t, subproduct.SalePrice, gotListing.Subproducts[ind].SalePrice)

}

func requireBodyMatchProductImage(t *testing.T, responseBody []byte, productImage db.ProductImage, ind int) {

	var gotListing db.ProductListingTxResult
	err := json.Unmarshal(responseBody, &gotListing)
	require.NoError(t, err)
	require.Equal(t, productImage.Title, gotListing.ProductImages[ind].Title)
	require.Equal(t, productImage.PicUrl, gotListing.ProductImages[ind].PicUrl)
	require.Equal(t, productImage.IsDefault, gotListing.ProductImages[ind].IsDefault)
}

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////
// Test CreateProductListingAPI //
///////////////////////////////////////////////////////////////

func TestCreateProductListingAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user.Email)

	product := randomProduct(store.ID)
	subproduct1 := randomSubproductWithSale(product.ID)
	subproduct2 := randomSubproductWithNoSale(product.ID)
	productImage1 := randomProductImage(product.ID)
	productImage2 := randomProductImage(product.ID)

	randomUUID := util.RandomUUID()
	user2, _ := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{

		{
			name: "OK",
			body: gin.H{
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"store_id":    product.StoreID.String(),
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.CreateProductListingTxParams{
					Category:    product.Category,
					Name:        product.Name,
					Description: product.Description,
					StoreID:     product.StoreID,
					Suspend:     product.Suspend,
					Subproducts: []struct {
						Variation   string `json:"variation" binding:"required"`
						StockAmount int32  `json:"stock_amount" binding:"required"`
						Price       int32  `json:"price" binding:"required"`
						SalePrice   int32  `json:"sale_price"`
					}{
						{
							Variation:   subproduct1.Variation,
							StockAmount: subproduct1.StockAmount,
							Price:       subproduct1.Price,
							SalePrice:   subproduct1.SalePrice.Int32,
						},
						{
							Variation:   subproduct2.Variation,
							StockAmount: subproduct2.StockAmount,
							Price:       subproduct2.Price,
							SalePrice:   subproduct2.SalePrice.Int32,
						},
					},
					ProductImages: []struct {
						Title     string `json:"title" binding:"required"`
						PicUrl    string `json:"pic_url" binding:"required"`
						IsDefault bool   `json:"is_default" binding:"required"`
					}{
						{
							Title:     productImage1.Title,
							PicUrl:    productImage1.PicUrl,
							IsDefault: productImage1.IsDefault,
						},
						{
							Title:     productImage2.Title,
							PicUrl:    productImage2.PicUrl,
							IsDefault: productImage2.IsDefault,
						},
					},
				}
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(
						store, nil)
				storeTx.EXPECT().
					CreateProductListingTx(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(
						db.ProductListingTxResult{
							Product:       product,
							Subproducts:   []db.Subproduct{subproduct1, subproduct2},
							ProductImages: []db.ProductImage{productImage1, productImage2},
						}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				responseBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)
				requireBodyMatchProduct(t, responseBody, product)
				requireBodyMatchSubproduct(t, responseBody, subproduct1, 0)
				requireBodyMatchSubproduct(t, responseBody, subproduct2, 1)
				requireBodyMatchProductImage(t, responseBody, productImage1, 0)
				requireBodyMatchProductImage(t, responseBody, productImage2, 1)
			},
		},
		{
			name: "MissingRequestBody",
			body: gin.H{
				"category": product.Category,
				// "name":        product.Name,
				"description": product.Description,
				"store_id":    "Invalid Store ID",
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidStoreID",
			body: gin.H{
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"store_id":    "InvalidStoreID",
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "StoreNotFound",
			body: gin.H{
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"store_id":    randomUUID.String(),
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(randomUUID)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileGettingStore",
			body: gin.H{
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"store_id":    product.StoreID.String(),
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UnauthorizedToManageStore",
			body: gin.H{
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"store_id":    product.StoreID.String(),
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user2.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileCreatingProductListing",
			body: gin.H{
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"store_id":    product.StoreID.String(),
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
						"sale_price":   subproduct1.SalePrice.Int32,
					},
					// Not onsell
					{
						"variation":    subproduct2.Variation,
						"stock_amount": subproduct2.StockAmount,
						"price":        subproduct2.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
					{
						"title":      productImage2.Title,
						"pic_url":    productImage2.PicUrl,
						"is_default": productImage2.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.CreateProductListingTxParams{
					Category:    product.Category,
					Name:        product.Name,
					Description: product.Description,
					StoreID:     product.StoreID,
					Suspend:     product.Suspend,
					Subproducts: []struct {
						Variation   string `json:"variation" binding:"required"`
						StockAmount int32  `json:"stock_amount" binding:"required"`
						Price       int32  `json:"price" binding:"required"`
						SalePrice   int32  `json:"sale_price"`
					}{
						{
							Variation:   subproduct1.Variation,
							StockAmount: subproduct1.StockAmount,
							Price:       subproduct1.Price,
							SalePrice:   subproduct1.SalePrice.Int32,
						},
						{
							Variation:   subproduct2.Variation,
							StockAmount: subproduct2.StockAmount,
							Price:       subproduct2.Price,
							SalePrice:   subproduct2.SalePrice.Int32,
						},
					},
					ProductImages: []struct {
						Title     string `json:"title" binding:"required"`
						PicUrl    string `json:"pic_url" binding:"required"`
						IsDefault bool   `json:"is_default" binding:"required"`
					}{
						{
							Title:     productImage1.Title,
							PicUrl:    productImage1.PicUrl,
							IsDefault: productImage1.IsDefault,
						},
						{
							Title:     productImage2.Title,
							PicUrl:    productImage2.PicUrl,
							IsDefault: productImage2.IsDefault,
						},
					},
				}
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(store, nil)
				storeTx.EXPECT().
					CreateProductListingTx(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.ProductListingTxResult{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {

		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeTx := mockdb.NewMockStoreTx(ctrl)
			tc.buildStubs(storeTx)

			// start test server and send request
			server := newTestServer(t, storeTx)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/stores/products"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}
}

///////////////////////////////////////////////////////////////
// Test UpdateProductListingAPI //
///////////////////////////////////////////////////////////////

func TestUpdateProductListingAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user.Email)

	product := randomProduct(store.ID)
	subproduct1 := randomSubproductWithNoSale(product.ID)
	productImage1 := randomProductImage(product.ID)

	randomUUID := util.RandomUUID()
	user2, _ := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"Product_id":  product.ID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.UpdateProductListingTxParams{
					ProductID:   product.ID,
					Category:    product.Category,
					Name:        product.Name,
					Description: product.Description,
					Suspend:     product.Suspend,
					Subproducts: []struct {
						Variation   string `json:"variation" binding:"required"`
						StockAmount int32  `json:"stock_amount" binding:"required"`
						Price       int32  `json:"price" binding:"required"`
						SalePrice   int32  `json:"sale_price"`
					}{
						{
							Variation:   subproduct1.Variation,
							StockAmount: subproduct1.StockAmount,
							Price:       subproduct1.Price,
							SalePrice:   subproduct1.SalePrice.Int32,
						},
					},
					ProductImages: []struct {
						Title     string `json:"title" binding:"required"`
						PicUrl    string `json:"pic_url" binding:"required"`
						IsDefault bool   `json:"is_default" binding:"required"`
					}{
						{
							Title:     productImage1.Title,
							PicUrl:    productImage1.PicUrl,
							IsDefault: productImage1.IsDefault,
						},
					},
				}
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(store, nil)
				storeTx.EXPECT().
					UpdateProductListingTx(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(
						db.ProductListingTxResult{
							Product:       product,
							Subproducts:   []db.Subproduct{subproduct1},
							ProductImages: []db.ProductImage{productImage1},
						}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				responseBody, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)
				requireBodyMatchProduct(t, responseBody, product)
				requireBodyMatchSubproduct(t, responseBody, subproduct1, 0)
				requireBodyMatchProductImage(t, responseBody, productImage1, 0)
			},
		},
		{
			name: "MissingRequestBody",
			body: gin.H{
				// "Product_id":  product.ID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidProductID",
			body: gin.H{
				"Product_id":  "InvalidProductID",
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ProductNotFound",
			body: gin.H{
				"Product_id":  randomUUID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(randomUUID)).
					Times(1).
					Return(db.Product{}, sql.ErrNoRows)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileGettingProduct",
			body: gin.H{
				"Product_id":  product.ID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.Product{}, sql.ErrConnDone)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileGettingStore",
			body: gin.H{
				"Product_id":  product.ID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UnauthorizedToManageStore",
			body: gin.H{
				"Product_id":  product.ID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user2.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(store, nil)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileUpdatingProductListing",
			body: gin.H{
				"Product_id":  product.ID.String(),
				"category":    product.Category,
				"name":        product.Name,
				"description": product.Description,
				"suspend":     product.Suspend,
				"subproducts": []gin.H{
					// Not Onsell
					{
						"variation":    subproduct1.Variation,
						"stock_amount": subproduct1.StockAmount,
						"price":        subproduct1.Price,
					},
				},
				"product_images": []gin.H{
					{
						"title":      productImage1.Title,
						"pic_url":    productImage1.PicUrl,
						"is_default": productImage1.IsDefault,
					},
				},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.UpdateProductListingTxParams{
					ProductID:   product.ID,
					Category:    product.Category,
					Name:        product.Name,
					Description: product.Description,
					Suspend:     product.Suspend,
					Subproducts: []struct {
						Variation   string `json:"variation" binding:"required"`
						StockAmount int32  `json:"stock_amount" binding:"required"`
						Price       int32  `json:"price" binding:"required"`
						SalePrice   int32  `json:"sale_price"`
					}{
						{
							Variation:   subproduct1.Variation,
							StockAmount: subproduct1.StockAmount,
							Price:       subproduct1.Price,
							SalePrice:   subproduct1.SalePrice.Int32,
						},
					},
					ProductImages: []struct {
						Title     string `json:"title" binding:"required"`
						PicUrl    string `json:"pic_url" binding:"required"`
						IsDefault bool   `json:"is_default" binding:"required"`
					}{
						{
							Title:     productImage1.Title,
							PicUrl:    productImage1.PicUrl,
							IsDefault: productImage1.IsDefault,
						},
					},
				}
				storeTx.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(product.StoreID)).
					Times(1).
					Return(store, nil)
				storeTx.EXPECT().
					UpdateProductListingTx(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.ProductListingTxResult{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {

		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeTx := mockdb.NewMockStoreTx(ctrl)
			tc.buildStubs(storeTx)

			// start test server and send request
			server := newTestServer(t, storeTx)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/stores/products"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
