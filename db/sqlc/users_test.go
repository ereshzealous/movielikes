package db

import (
	"context"
	"testing"

	"example.com/movielikes/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		UserName: util.RandomUsername(),
		HashedPassword: hashedPassword,
		FullName: util.RandomString(10),
		Email: util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t);
	actualUser, err := testQueries.GetUser(context.Background(), user.UserName)
	require.NoError(t, err)
	require.NotEmpty(t, actualUser)
	require.Equal(t, user.UserName, actualUser.UserName)
	require.Equal(t, user.HashedPassword, actualUser.HashedPassword)
	require.Equal(t, user.Email, actualUser.Email)
	require.Equal(t, user.FullName, actualUser.FullName)
	require.Equal(t, user.PasswordChangedAt, actualUser.PasswordChangedAt)
	require.Equal(t, user.CreatedAt, actualUser.CreatedAt)
}