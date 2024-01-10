package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mockdb "github.com/chomchaby/shopchaby-backend-go/db/mock"
	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/token"
	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user.Email, gotUser.Email)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Phone, gotUser.Phone)
	require.Equal(t, user.Address, gotUser.Address)
	require.Empty(t, gotUser.PwdHash)
}

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams) // convert to db.CreateUserParams object
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.PwdHash)
	if err != nil {
		return false
	}

	e.arg.PwdHash = arg.PwdHash
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

///////////////////////////////////////////////////////////////
// Test CreateUserAPI //
///////////////////////////////////////////////////////////////

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    user.Email,
				"username": user.Username,
				"pwd":      password,
				"phone":    user.Phone,
				"address":  user.Address,
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.CreateUserParams{
					Email:    user.Email,
					Username: user.Username,
					// PwdHash:  user.PwdHash,
					Phone:   user.Phone,
					Address: user.Address,
				}
				storeTx.EXPECT().
					// CreateUser(gomock.Any(), gomock.Eq(arg)).
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"email":    user.Email,
				"username": user.Username,
				"pwd":      password,
				"phone":    user.Phone,
				"address":  user.Address,
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateEmailOrUsername",
			body: gin.H{
				"email":    user.Email,
				"username": user.Username,
				"pwd":      password,
				"phone":    user.Phone,
				"address":  user.Address,
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"email":    "InvalidEmail",
				"username": user.Username,
				"pwd":      password,
				"phone":    user.Phone,
				"address":  user.Address,
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "TooShortPassword",
			body: gin.H{
				"email":    user.Email,
				"username": user.Username,
				"pwd":      "123",
				"phone":    user.Phone,
				"address":  user.Address,
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
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

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}

///////////////////////////////////////////////////////////////
// Test GetUserAPI //
///////////////////////////////////////////////////////////////

func TestGetUserAPI(t *testing.T) {
	user, _ := randomUser(t)

	testCases := []struct {
		name          string
		email         string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			email: user.Email,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:  "NotFound",
			email: user.Email,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:  "InternalError",
			email: user.Email,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "InvalidEmail",
			email: "InvalidEmail",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/users/%s", tc.email)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}

///////////////////////////////////////////////////////////////
// Test UpdateUserAPI //
///////////////////////////////////////////////////////////////

type eqUpdateUserParamsMatcher struct {
	arg      db.UpdateUserParams
	password string
}

func (e eqUpdateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.UpdateUserParams) // convert to db.UpdateUserParams object
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.PwdHash)
	if err != nil {
		return false
	}

	e.arg.PwdHash = arg.PwdHash
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqUpdateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqUpdateUserParams(arg db.UpdateUserParams, password string) gomock.Matcher {
	return eqUpdateUserParamsMatcher{arg, password}
}

func TestUpdateUserAPI(t *testing.T) {
	user, _ := randomUser(t)
	user2, _ := randomUser(t)
	new_username := "new_username"
	new_pwd := "new_pwd"
	new_phone := "new_phone"

	testCases := []struct {
		name          string
		email         string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(storeTx *mockdb.MockStoreTx)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			email: user.Email,
			body: gin.H{
				"username": new_username,
				"pwd":      new_pwd,
				"phone":    new_phone,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				arg := db.UpdateUserParams{
					Email:    user.Email,
					Username: new_username,
					// PwdHash:  user.PwdHash,
					Phone:   new_phone,
					Address: user.Address,
				}
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)

				storeTx.EXPECT().
					UpdateUser(gomock.Any(), EqUpdateUserParams(arg, new_pwd)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:  "NotFound",
			email: "notfound@example.com",
			body: gin.H{
				"username": new_username,
				"pwd":      new_pwd,
				"phone":    new_phone,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq("notfound@example.com")).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:  "InternalErrorWhileUpdatingUser",
			email: user.Email,
			body: gin.H{
				"username": new_username,
				"pwd":      new_pwd,
				"phone":    new_phone,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
				storeTx.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "InternalErrorWhileGettingUser",
			email: user.Email,
			body: gin.H{
				"username": new_username,
				"pwd":      new_pwd,
				"phone":    new_phone,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:  "DuplicateUsername",
			email: user.Email,
			body: gin.H{
				"username": user2.Username,
				"pwd":      new_pwd,
				"phone":    new_phone,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
				storeTx.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name:  "InvalidEmail",
			email: "InvalidEmail",
			body: gin.H{
				"username": new_username,
				"pwd":      new_pwd,
				"phone":    new_phone,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStubs: func(storeTx *mockdb.MockStoreTx) {
				storeTx.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/users/%s", tc.email)
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}

}
