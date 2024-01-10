package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/chomchaby/shopchaby-backend-go/db/mock"
	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/token"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchStore(t *testing.T, body *bytes.Buffer, store db.Store) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotStore db.Store
	err = json.Unmarshal(data, &gotStore)
	require.NoError(t, err)
	require.Equal(t, store.UserEmail, gotStore.UserEmail)
	require.Equal(t, store.Name, gotStore.Name)
	require.Equal(t, store.Address, gotStore.Address)
	require.Equal(t, store.Phone, gotStore.Phone)
	require.Equal(t, store.PicUrl, gotStore.PicUrl)

}

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////
// Test CreateStoreAPI //
///////////////////////////////////////////////////////////////

func TestCreateStoreAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user.Email)

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
				"user_email":  store.UserEmail,
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.CreateStoreParams{
					UserEmail:   store.UserEmail,
					Name:        store.Name,
					Description: store.Description,
					Address:     store.Address,
					Phone:       store.Phone,
					PicUrl:      store.PicUrl,
				}
				storeTx.EXPECT().
					CreateStore(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStore(t, recorder.Body, store)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"user_email":  store.UserEmail,
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					CreateStore(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateUserEmail",
			body: gin.H{
				"user_email":  store.UserEmail,
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					CreateStore(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Store{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
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

			url := "/users/stores"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}

///////////////////////////////////////////////////////////////
// Test GetStoreAPI //
///////////////////////////////////////////////////////////////

func TestGetStoreAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user.Email)
	testCases := []struct {
		name          string
		id            string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   store.ID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStore(t, recorder.Body, store)
			},
		},
		{
			name: "NotFound",
			id:   store.ID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			id:   store.ID.String(),
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			id:   "InvalidID",
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

			url := fmt.Sprintf("/users/stores/%s", tc.id)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}

///////////////////////////////////////////////////////////////
// Test UpdateStoreAPI //
///////////////////////////////////////////////////////////////

func TestUpdateStoreAPI(t *testing.T) {
	user, _ := randomUser(t)
	store := randomStore(user.Email)

	testCases := []struct {
		name          string
		id            string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   store.ID.String(),
			body: gin.H{
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.UpdateStoreParams{
					ID:          store.ID,
					Name:        store.Name,
					Description: store.Description,
					Address:     store.Address,
					Phone:       store.Phone,
					PicUrl:      store.PicUrl,
				}
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil) // Returns the expected store when GetStore is called with the provided ID

				storeTx.EXPECT().
					UpdateStore(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(store, nil) // Returns the updated store when UpdateStore is called with the provided parameters
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStore(t, recorder.Body, store)
			},
		},
		{
			name: "NotFound",
			id:   store.ID.String(),
			body: gin.H{
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileUpdatingStore",
			id:   store.ID.String(),
			body: gin.H{
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Eq(store.ID)).
					Times(1).
					Return(store, nil)
				storeTx.EXPECT().
					UpdateStore(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalErrorWhileGettingStore",
			id:   store.ID.String(),
			body: gin.H{
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetStore(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			id:   "InvalidID",
			body: gin.H{
				"name":        store.Name,
				"description": store.Description,
				"address":     store.Address,
				"phone":       store.Phone,
				"pic_url":     store.PicUrl,
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

			url := fmt.Sprintf("/users/stores/%s", tc.id)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}
