// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chomchaby/shopchaby-backend-go/db/sqlc (interfaces: StoreTx)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStoreTx is a mock of StoreTx interface.
type MockStoreTx struct {
	ctrl     *gomock.Controller
	recorder *MockStoreTxMockRecorder
}

// MockStoreTxMockRecorder is the mock recorder for MockStoreTx.
type MockStoreTxMockRecorder struct {
	mock *MockStoreTx
}

// NewMockStoreTx creates a new mock instance.
func NewMockStoreTx(ctrl *gomock.Controller) *MockStoreTx {
	mock := &MockStoreTx{ctrl: ctrl}
	mock.recorder = &MockStoreTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreTx) EXPECT() *MockStoreTxMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockStoreTx) CreateCart(arg0 context.Context, arg1 db.CreateCartParams) (db.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", arg0, arg1)
	ret0, _ := ret[0].(db.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockStoreTxMockRecorder) CreateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockStoreTx)(nil).CreateCart), arg0, arg1)
}

// CreateFavorite mocks base method.
func (m *MockStoreTx) CreateFavorite(arg0 context.Context, arg1 db.CreateFavoriteParams) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFavorite indicates an expected call of CreateFavorite.
func (mr *MockStoreTxMockRecorder) CreateFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFavorite", reflect.TypeOf((*MockStoreTx)(nil).CreateFavorite), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStoreTx) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreTxMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStoreTx)(nil).CreateProduct), arg0, arg1)
}

// CreateProductCategory mocks base method.
func (m *MockStoreTx) CreateProductCategory(arg0 context.Context, arg1 db.CreateProductCategoryParams) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductCategory indicates an expected call of CreateProductCategory.
func (mr *MockStoreTxMockRecorder) CreateProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductCategory", reflect.TypeOf((*MockStoreTx)(nil).CreateProductCategory), arg0, arg1)
}

// CreateProductImage mocks base method.
func (m *MockStoreTx) CreateProductImage(arg0 context.Context, arg1 db.CreateProductImageParams) (db.ProductImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductImage", arg0, arg1)
	ret0, _ := ret[0].(db.ProductImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductImage indicates an expected call of CreateProductImage.
func (mr *MockStoreTxMockRecorder) CreateProductImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductImage", reflect.TypeOf((*MockStoreTx)(nil).CreateProductImage), arg0, arg1)
}

// CreateProductListingTx mocks base method.
func (m *MockStoreTx) CreateProductListingTx(arg0 context.Context, arg1 db.CreateProductListingTxParams) (db.ProductListingTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductListingTx", arg0, arg1)
	ret0, _ := ret[0].(db.ProductListingTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductListingTx indicates an expected call of CreateProductListingTx.
func (mr *MockStoreTxMockRecorder) CreateProductListingTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductListingTx", reflect.TypeOf((*MockStoreTx)(nil).CreateProductListingTx), arg0, arg1)
}

// CreatePurchase mocks base method.
func (m *MockStoreTx) CreatePurchase(arg0 context.Context, arg1 db.CreatePurchaseParams) (db.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePurchase", arg0, arg1)
	ret0, _ := ret[0].(db.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePurchase indicates an expected call of CreatePurchase.
func (mr *MockStoreTxMockRecorder) CreatePurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePurchase", reflect.TypeOf((*MockStoreTx)(nil).CreatePurchase), arg0, arg1)
}

// CreateStore mocks base method.
func (m *MockStoreTx) CreateStore(arg0 context.Context, arg1 db.CreateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStore indicates an expected call of CreateStore.
func (mr *MockStoreTxMockRecorder) CreateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStore", reflect.TypeOf((*MockStoreTx)(nil).CreateStore), arg0, arg1)
}

// CreateSubproduct mocks base method.
func (m *MockStoreTx) CreateSubproduct(arg0 context.Context, arg1 db.CreateSubproductParams) (db.Subproduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubproduct", arg0, arg1)
	ret0, _ := ret[0].(db.Subproduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubproduct indicates an expected call of CreateSubproduct.
func (mr *MockStoreTxMockRecorder) CreateSubproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubproduct", reflect.TypeOf((*MockStoreTx)(nil).CreateSubproduct), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStoreTx) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreTxMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStoreTx)(nil).CreateUser), arg0, arg1)
}

// DeleteCart mocks base method.
func (m *MockStoreTx) DeleteCart(arg0 context.Context, arg1 db.DeleteCartParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockStoreTxMockRecorder) DeleteCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockStoreTx)(nil).DeleteCart), arg0, arg1)
}

// DeleteFavorite mocks base method.
func (m *MockStoreTx) DeleteFavorite(arg0 context.Context, arg1 db.DeleteFavoriteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFavorite", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFavorite indicates an expected call of DeleteFavorite.
func (mr *MockStoreTxMockRecorder) DeleteFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFavorite", reflect.TypeOf((*MockStoreTx)(nil).DeleteFavorite), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStoreTx) DeleteProduct(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreTxMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStoreTx)(nil).DeleteProduct), arg0, arg1)
}

// DeleteProductCategory mocks base method.
func (m *MockStoreTx) DeleteProductCategory(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductCategory indicates an expected call of DeleteProductCategory.
func (mr *MockStoreTxMockRecorder) DeleteProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductCategory", reflect.TypeOf((*MockStoreTx)(nil).DeleteProductCategory), arg0, arg1)
}

// DeleteProductImage mocks base method.
func (m *MockStoreTx) DeleteProductImage(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductImage indicates an expected call of DeleteProductImage.
func (mr *MockStoreTxMockRecorder) DeleteProductImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductImage", reflect.TypeOf((*MockStoreTx)(nil).DeleteProductImage), arg0, arg1)
}

// DeleteProductImagesByProductID mocks base method.
func (m *MockStoreTx) DeleteProductImagesByProductID(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductImagesByProductID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductImagesByProductID indicates an expected call of DeleteProductImagesByProductID.
func (mr *MockStoreTxMockRecorder) DeleteProductImagesByProductID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductImagesByProductID", reflect.TypeOf((*MockStoreTx)(nil).DeleteProductImagesByProductID), arg0, arg1)
}

// DeletePurchase mocks base method.
func (m *MockStoreTx) DeletePurchase(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePurchase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePurchase indicates an expected call of DeletePurchase.
func (mr *MockStoreTxMockRecorder) DeletePurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePurchase", reflect.TypeOf((*MockStoreTx)(nil).DeletePurchase), arg0, arg1)
}

// DeleteStore mocks base method.
func (m *MockStoreTx) DeleteStore(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStore indicates an expected call of DeleteStore.
func (mr *MockStoreTxMockRecorder) DeleteStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStore", reflect.TypeOf((*MockStoreTx)(nil).DeleteStore), arg0, arg1)
}

// DeleteSubproduct mocks base method.
func (m *MockStoreTx) DeleteSubproduct(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubproduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubproduct indicates an expected call of DeleteSubproduct.
func (mr *MockStoreTxMockRecorder) DeleteSubproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubproduct", reflect.TypeOf((*MockStoreTx)(nil).DeleteSubproduct), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStoreTx) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreTxMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStoreTx)(nil).DeleteUser), arg0, arg1)
}

// GetCart mocks base method.
func (m *MockStoreTx) GetCart(arg0 context.Context, arg1 db.GetCartParams) (db.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0, arg1)
	ret0, _ := ret[0].(db.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockStoreTxMockRecorder) GetCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockStoreTx)(nil).GetCart), arg0, arg1)
}

// GetFavorite mocks base method.
func (m *MockStoreTx) GetFavorite(arg0 context.Context, arg1 db.GetFavoriteParams) (db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorite", arg0, arg1)
	ret0, _ := ret[0].(db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorite indicates an expected call of GetFavorite.
func (mr *MockStoreTxMockRecorder) GetFavorite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorite", reflect.TypeOf((*MockStoreTx)(nil).GetFavorite), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStoreTx) GetProduct(arg0 context.Context, arg1 uuid.UUID) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreTxMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStoreTx)(nil).GetProduct), arg0, arg1)
}

// GetProductCategory mocks base method.
func (m *MockStoreTx) GetProductCategory(arg0 context.Context, arg1 string) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCategory indicates an expected call of GetProductCategory.
func (mr *MockStoreTxMockRecorder) GetProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCategory", reflect.TypeOf((*MockStoreTx)(nil).GetProductCategory), arg0, arg1)
}

// GetProductImage mocks base method.
func (m *MockStoreTx) GetProductImage(arg0 context.Context, arg1 uuid.UUID) (db.ProductImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductImage", arg0, arg1)
	ret0, _ := ret[0].(db.ProductImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductImage indicates an expected call of GetProductImage.
func (mr *MockStoreTxMockRecorder) GetProductImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductImage", reflect.TypeOf((*MockStoreTx)(nil).GetProductImage), arg0, arg1)
}

// GetPurchase mocks base method.
func (m *MockStoreTx) GetPurchase(arg0 context.Context, arg1 uuid.UUID) (db.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPurchase", arg0, arg1)
	ret0, _ := ret[0].(db.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPurchase indicates an expected call of GetPurchase.
func (mr *MockStoreTxMockRecorder) GetPurchase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPurchase", reflect.TypeOf((*MockStoreTx)(nil).GetPurchase), arg0, arg1)
}

// GetStore mocks base method.
func (m *MockStoreTx) GetStore(arg0 context.Context, arg1 uuid.UUID) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore.
func (mr *MockStoreTxMockRecorder) GetStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockStoreTx)(nil).GetStore), arg0, arg1)
}

// GetStoreByUserEmail mocks base method.
func (m *MockStoreTx) GetStoreByUserEmail(arg0 context.Context, arg1 string) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreByUserEmail", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreByUserEmail indicates an expected call of GetStoreByUserEmail.
func (mr *MockStoreTxMockRecorder) GetStoreByUserEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreByUserEmail", reflect.TypeOf((*MockStoreTx)(nil).GetStoreByUserEmail), arg0, arg1)
}

// GetSubproduct mocks base method.
func (m *MockStoreTx) GetSubproduct(arg0 context.Context, arg1 uuid.UUID) (db.Subproduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubproduct", arg0, arg1)
	ret0, _ := ret[0].(db.Subproduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubproduct indicates an expected call of GetSubproduct.
func (mr *MockStoreTxMockRecorder) GetSubproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubproduct", reflect.TypeOf((*MockStoreTx)(nil).GetSubproduct), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStoreTx) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreTxMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStoreTx)(nil).GetUser), arg0, arg1)
}

// ListCarts mocks base method.
func (m *MockStoreTx) ListCarts(arg0 context.Context) ([]db.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCarts", arg0)
	ret0, _ := ret[0].([]db.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCarts indicates an expected call of ListCarts.
func (mr *MockStoreTxMockRecorder) ListCarts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCarts", reflect.TypeOf((*MockStoreTx)(nil).ListCarts), arg0)
}

// ListCartsByUser mocks base method.
func (m *MockStoreTx) ListCartsByUser(arg0 context.Context, arg1 string) ([]db.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCartsByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCartsByUser indicates an expected call of ListCartsByUser.
func (mr *MockStoreTxMockRecorder) ListCartsByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCartsByUser", reflect.TypeOf((*MockStoreTx)(nil).ListCartsByUser), arg0, arg1)
}

// ListFavorites mocks base method.
func (m *MockStoreTx) ListFavorites(arg0 context.Context) ([]db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavorites", arg0)
	ret0, _ := ret[0].([]db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFavorites indicates an expected call of ListFavorites.
func (mr *MockStoreTxMockRecorder) ListFavorites(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavorites", reflect.TypeOf((*MockStoreTx)(nil).ListFavorites), arg0)
}

// ListFavoritesByUser mocks base method.
func (m *MockStoreTx) ListFavoritesByUser(arg0 context.Context, arg1 string) ([]db.Favorite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFavoritesByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.Favorite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFavoritesByUser indicates an expected call of ListFavoritesByUser.
func (mr *MockStoreTxMockRecorder) ListFavoritesByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFavoritesByUser", reflect.TypeOf((*MockStoreTx)(nil).ListFavoritesByUser), arg0, arg1)
}

// ListProductCategories mocks base method.
func (m *MockStoreTx) ListProductCategories(arg0 context.Context) ([]db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductCategories", arg0)
	ret0, _ := ret[0].([]db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductCategories indicates an expected call of ListProductCategories.
func (mr *MockStoreTxMockRecorder) ListProductCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductCategories", reflect.TypeOf((*MockStoreTx)(nil).ListProductCategories), arg0)
}

// ListProductImages mocks base method.
func (m *MockStoreTx) ListProductImages(arg0 context.Context) ([]db.ProductImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductImages", arg0)
	ret0, _ := ret[0].([]db.ProductImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductImages indicates an expected call of ListProductImages.
func (mr *MockStoreTxMockRecorder) ListProductImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductImages", reflect.TypeOf((*MockStoreTx)(nil).ListProductImages), arg0)
}

// ListProductImagesByProductID mocks base method.
func (m *MockStoreTx) ListProductImagesByProductID(arg0 context.Context, arg1 uuid.UUID) ([]db.ProductImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductImagesByProductID", arg0, arg1)
	ret0, _ := ret[0].([]db.ProductImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductImagesByProductID indicates an expected call of ListProductImagesByProductID.
func (mr *MockStoreTxMockRecorder) ListProductImagesByProductID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductImagesByProductID", reflect.TypeOf((*MockStoreTx)(nil).ListProductImagesByProductID), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStoreTx) ListProducts(arg0 context.Context) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreTxMockRecorder) ListProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStoreTx)(nil).ListProducts), arg0)
}

// ListProductsByCategory mocks base method.
func (m *MockStoreTx) ListProductsByCategory(arg0 context.Context, arg1 string) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductsByCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductsByCategory indicates an expected call of ListProductsByCategory.
func (mr *MockStoreTxMockRecorder) ListProductsByCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductsByCategory", reflect.TypeOf((*MockStoreTx)(nil).ListProductsByCategory), arg0, arg1)
}

// ListProductsByFilter mocks base method.
func (m *MockStoreTx) ListProductsByFilter(arg0 context.Context, arg1 db.ListProductsByFilterParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductsByFilter", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductsByFilter indicates an expected call of ListProductsByFilter.
func (mr *MockStoreTxMockRecorder) ListProductsByFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductsByFilter", reflect.TypeOf((*MockStoreTx)(nil).ListProductsByFilter), arg0, arg1)
}

// ListProductsByName mocks base method.
func (m *MockStoreTx) ListProductsByName(arg0 context.Context, arg1 string) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductsByName", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductsByName indicates an expected call of ListProductsByName.
func (mr *MockStoreTxMockRecorder) ListProductsByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductsByName", reflect.TypeOf((*MockStoreTx)(nil).ListProductsByName), arg0, arg1)
}

// ListProductsByStoreAndCategory mocks base method.
func (m *MockStoreTx) ListProductsByStoreAndCategory(arg0 context.Context, arg1 db.ListProductsByStoreAndCategoryParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductsByStoreAndCategory", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductsByStoreAndCategory indicates an expected call of ListProductsByStoreAndCategory.
func (mr *MockStoreTxMockRecorder) ListProductsByStoreAndCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductsByStoreAndCategory", reflect.TypeOf((*MockStoreTx)(nil).ListProductsByStoreAndCategory), arg0, arg1)
}

// ListPurchases mocks base method.
func (m *MockStoreTx) ListPurchases(arg0 context.Context) ([]db.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPurchases", arg0)
	ret0, _ := ret[0].([]db.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPurchases indicates an expected call of ListPurchases.
func (mr *MockStoreTxMockRecorder) ListPurchases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPurchases", reflect.TypeOf((*MockStoreTx)(nil).ListPurchases), arg0)
}

// ListPurchasesByBuyer mocks base method.
func (m *MockStoreTx) ListPurchasesByBuyer(arg0 context.Context, arg1 string) ([]db.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPurchasesByBuyer", arg0, arg1)
	ret0, _ := ret[0].([]db.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPurchasesByBuyer indicates an expected call of ListPurchasesByBuyer.
func (mr *MockStoreTxMockRecorder) ListPurchasesByBuyer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPurchasesByBuyer", reflect.TypeOf((*MockStoreTx)(nil).ListPurchasesByBuyer), arg0, arg1)
}

// ListPurchasesByBuyerAndStore mocks base method.
func (m *MockStoreTx) ListPurchasesByBuyerAndStore(arg0 context.Context, arg1 db.ListPurchasesByBuyerAndStoreParams) ([]db.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPurchasesByBuyerAndStore", arg0, arg1)
	ret0, _ := ret[0].([]db.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPurchasesByBuyerAndStore indicates an expected call of ListPurchasesByBuyerAndStore.
func (mr *MockStoreTxMockRecorder) ListPurchasesByBuyerAndStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPurchasesByBuyerAndStore", reflect.TypeOf((*MockStoreTx)(nil).ListPurchasesByBuyerAndStore), arg0, arg1)
}

// ListStores mocks base method.
func (m *MockStoreTx) ListStores(arg0 context.Context, arg1 db.ListStoresParams) ([]db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStores", arg0, arg1)
	ret0, _ := ret[0].([]db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStores indicates an expected call of ListStores.
func (mr *MockStoreTxMockRecorder) ListStores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStores", reflect.TypeOf((*MockStoreTx)(nil).ListStores), arg0, arg1)
}

// ListSubproductsByProductID mocks base method.
func (m *MockStoreTx) ListSubproductsByProductID(arg0 context.Context, arg1 uuid.UUID) ([]db.Subproduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubproductsByProductID", arg0, arg1)
	ret0, _ := ret[0].([]db.Subproduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubproductsByProductID indicates an expected call of ListSubproductsByProductID.
func (mr *MockStoreTxMockRecorder) ListSubproductsByProductID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubproductsByProductID", reflect.TypeOf((*MockStoreTx)(nil).ListSubproductsByProductID), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStoreTx) ListUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreTxMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStoreTx)(nil).ListUsers), arg0)
}

// UpdateCart mocks base method.
func (m *MockStoreTx) UpdateCart(arg0 context.Context, arg1 db.UpdateCartParams) (db.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", arg0, arg1)
	ret0, _ := ret[0].(db.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockStoreTxMockRecorder) UpdateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockStoreTx)(nil).UpdateCart), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockStoreTx) UpdateProduct(arg0 context.Context, arg1 db.UpdateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreTxMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStoreTx)(nil).UpdateProduct), arg0, arg1)
}

// UpdateProductCategory mocks base method.
func (m *MockStoreTx) UpdateProductCategory(arg0 context.Context, arg1 db.UpdateProductCategoryParams) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductCategory indicates an expected call of UpdateProductCategory.
func (mr *MockStoreTxMockRecorder) UpdateProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductCategory", reflect.TypeOf((*MockStoreTx)(nil).UpdateProductCategory), arg0, arg1)
}

// UpdateProductImage mocks base method.
func (m *MockStoreTx) UpdateProductImage(arg0 context.Context, arg1 db.UpdateProductImageParams) (db.ProductImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductImage", arg0, arg1)
	ret0, _ := ret[0].(db.ProductImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductImage indicates an expected call of UpdateProductImage.
func (mr *MockStoreTxMockRecorder) UpdateProductImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductImage", reflect.TypeOf((*MockStoreTx)(nil).UpdateProductImage), arg0, arg1)
}

// UpdateProductListingTx mocks base method.
func (m *MockStoreTx) UpdateProductListingTx(arg0 context.Context, arg1 db.UpdateProductListingTxParams) (db.ProductListingTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductListingTx", arg0, arg1)
	ret0, _ := ret[0].(db.ProductListingTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductListingTx indicates an expected call of UpdateProductListingTx.
func (mr *MockStoreTxMockRecorder) UpdateProductListingTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductListingTx", reflect.TypeOf((*MockStoreTx)(nil).UpdateProductListingTx), arg0, arg1)
}

// UpdateStore mocks base method.
func (m *MockStoreTx) UpdateStore(arg0 context.Context, arg1 db.UpdateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStore indicates an expected call of UpdateStore.
func (mr *MockStoreTxMockRecorder) UpdateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStore", reflect.TypeOf((*MockStoreTx)(nil).UpdateStore), arg0, arg1)
}

// UpdateSubproduct mocks base method.
func (m *MockStoreTx) UpdateSubproduct(arg0 context.Context, arg1 db.UpdateSubproductParams) (db.Subproduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubproduct", arg0, arg1)
	ret0, _ := ret[0].(db.Subproduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubproduct indicates an expected call of UpdateSubproduct.
func (mr *MockStoreTxMockRecorder) UpdateSubproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubproduct", reflect.TypeOf((*MockStoreTx)(nil).UpdateSubproduct), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStoreTx) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreTxMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStoreTx)(nil).UpdateUser), arg0, arg1)
}