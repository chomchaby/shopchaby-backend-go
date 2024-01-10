package db

import (
	"context"
	"testing"

	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomCart(t *testing.T) Cart {
	user := createRandomUser(t)

	listing := createRandomProductListingTx(t)
	var arg CreateCartParams
	if listing[0].Subproducts[0].StockAmount != 0 {
		arg = CreateCartParams{
			UserEmail:    user.Email,
			SubproductID: listing[0].Subproducts[0].ID,
			Quantity:     int32(util.RandomInt(1, int64(listing[0].Subproducts[0].StockAmount))),
		}
	} else if listing[1].Subproducts[0].StockAmount != 0 {
		arg = CreateCartParams{
			UserEmail:    user.Email,
			SubproductID: listing[0].Subproducts[0].ID,
			Quantity:     int32(util.RandomInt(1, int64(listing[1].Subproducts[0].StockAmount))),
		}
	} else {
		arg = CreateCartParams{
			UserEmail:    user.Email,
			SubproductID: listing[0].Subproducts[0].ID,
			Quantity:     int32(util.RandomInt(1, int64(listing[2].Subproducts[0].StockAmount))),
		}
	}

	cart, err := testQueries.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.UserEmail, cart.UserEmail)
	require.Equal(t, arg.SubproductID, cart.SubproductID)
	require.Equal(t, arg.Quantity, cart.Quantity)

	require.NotZero(t, cart.CreatedAt)

	return cart
}

func TestCreateCart(t *testing.T) {
	createRandomCart(t)
}
