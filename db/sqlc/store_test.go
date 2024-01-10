package db

import (
	"context"
	"testing"

	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomStore(t *testing.T) Store {
	randomUser := createRandomUser(t)
	randomStore := util.RandomStoreDetail()
	arg := CreateStoreParams{
		UserEmail:   randomUser.Email,
		Name:        randomStore.Name,
		Description: randomStore.Description,
		Address:     randomStore.Address,
		Phone:       randomStore.Phone,
		PicUrl:      randomStore.PicUrl,
	}

	store, err := testQueries.CreateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store)

	require.Equal(t, arg.UserEmail, store.UserEmail)
	require.Equal(t, arg.Name, store.Name)
	require.Equal(t, arg.Description, store.Description)
	require.Equal(t, arg.Address, store.Address)
	require.Equal(t, arg.Phone, store.Phone)
	require.Equal(t, arg.PicUrl, store.PicUrl)
	require.Equal(t, int32(0), store.Balance)

	require.NotZero(t, store.ID)
	require.NotZero(t, store.CreatedAt)
	require.NotZero(t, store.UpdatedAt)

	return store
}

func TestCreateStore(t *testing.T) {
	createRandomStore(t)
}

// To Do

// TestGetStore
// GetStoreByUserEmail
// TestUpdateStore
// TestDeleteStore
