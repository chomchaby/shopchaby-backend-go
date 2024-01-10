package db

import (
	"context"
	"testing"
	"time"

	"github.com/chomchaby/shopchaby-backend-go/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	random := util.RandomUserDetail()
	hashedPassword, err := util.HashPassword("secret")
	require.NoError(t, err)

	arg := CreateUserParams{
		Email:    random.Email,
		Username: random.Username,
		PwdHash:  hashedPassword,
		Phone:    random.Phone,
		Address:  random.Address,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PwdHash, user.PwdHash)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Address, user.Address)
	require.Equal(t, int32(0), user.Balance)

	// require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.PwdHash, user2.PwdHash)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Address, user2.Address)
	require.Equal(t, user1.Balance, user2.Balance)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

// To Do
// Test UpdateUser
