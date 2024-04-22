package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/raphaeldiscky/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	user1 := createRandomUser(t)

	newFullName := util.RandomOwner()
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: user1.Username,
		FullName: pgtype.Text{
			String: newFullName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user1.FullName, updatedUser.FullName)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, user1.Email, updatedUser.Email)
	require.Equal(t, user1.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	user1 := createRandomUser(t)

	newEmail := util.RandomEmail()
	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: user1.Username,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user1.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, user1.FullName, updatedUser.FullName)
	require.Equal(t, user1.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	user1 := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: user1.Username,
		HashedPassword: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user1.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, user1.FullName, updatedUser.FullName)
	require.Equal(t, user1.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	user1 := createRandomUser(t)

	newPassword := util.RandomString(6)
	newEmail := util.RandomEmail()
	newFullName := util.RandomOwner()
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
		Username: user1.Username,
		HashedPassword: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
		FullName: pgtype.Text{
			String: newFullName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, user1.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.NotEqual(t, user1.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, user1.FullName, updatedUser.FullName)
	require.Equal(t, newFullName, updatedUser.FullName)
}
