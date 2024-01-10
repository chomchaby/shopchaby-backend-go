package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomFavorite(t *testing.T) Favorite {
	user := createRandomUser(t)

	listing := createRandomProductListingTx(t)

	arg := CreateFavoriteParams{
		UserEmail: user.Email,
		ProductID: listing[0].Product.ID,
	}

	favorite, err := testQueries.CreateFavorite(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, favorite)

	require.Equal(t, arg.UserEmail, favorite.UserEmail)
	require.Equal(t, arg.ProductID, favorite.ProductID)

	require.NotZero(t, favorite.Timestamp)

	return favorite
}

func TestCreateFavorite(t *testing.T) {
	createRandomFavorite(t)
}
