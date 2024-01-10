package api

import (
	"testing"

	db "github.com/chomchaby/shopchaby-backend-go/db/sqlc"
	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user db.User, password string) {
	userDetails := util.RandomUserDetail()
	password = util.RandomString(10)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Email:    userDetails.Email,
		Username: userDetails.Username,
		PwdHash:  hashedPassword,
		Phone:    userDetails.Phone,
		Address:  userDetails.Address,
	}
	return
}

func randomStore(userEmail string) db.Store {
	storeDetails := util.RandomStoreDetail()
	return db.Store{
		ID:          storeDetails.ID,
		UserEmail:   userEmail,
		Name:        storeDetails.Name,
		Description: storeDetails.Description,
		Address:     storeDetails.Address,
		Phone:       storeDetails.Phone,
		PicUrl:      storeDetails.PicUrl,
	}
}

func randomProduct(store_id uuid.UUID) db.Product {
	productDetails := util.RandomProductDetail()
	return db.Product{
		ID:          productDetails.ID,
		StoreID:     store_id,
		Category:    productDetails.Category,
		Name:        productDetails.Name,
		Description: productDetails.Description,
		Suspend:     productDetails.Suspend,
	}
}

func randomSubproductWithNoSale(product_id uuid.UUID) db.Subproduct {
	subproductDetails := util.RandomSubproductDetail(false)
	return db.Subproduct{
		ID:          subproductDetails.ID,
		ProductID:   product_id,
		Variation:   subproductDetails.Variation,
		StockAmount: subproductDetails.StockAmount,
		Price:       subproductDetails.Price,
		SalePrice:   subproductDetails.SalePrice,
	}
}

func randomSubproductWithSale(product_id uuid.UUID) db.Subproduct {
	subproductDetails := util.RandomSubproductDetail(true)
	return db.Subproduct{
		ID:          subproductDetails.ID,
		ProductID:   product_id,
		Variation:   subproductDetails.Variation,
		StockAmount: subproductDetails.StockAmount,
		Price:       subproductDetails.Price,
		SalePrice:   subproductDetails.SalePrice,
	}
}

func randomProductImage(product_id uuid.UUID) db.ProductImage {
	productImageDetails := util.RandomProductImageDetail()
	return db.ProductImage{
		ID:        productImageDetails.ID,
		ProductID: product_id,
		Title:     productImageDetails.Title,
		PicUrl:    productImageDetails.PicUrl,
		IsDefault: productImageDetails.IsDefault,
	}
}
