package db

import (
	"context"
	"testing"
	"time"

	"github.com/raphaeldiscky/simple-bank.git/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: utils.RandomOwner(),
		HashedPassword: "secret",
		FullName: utils.RandomOwner(),
		Email: utils.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
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
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

// func TestUpdateuser(t *testing.T) {
// 	user1 := createRandomUser(t)

// 	arg := UpdateUserCreateUserParams{
// 		Username: user1.Username,
// 		Balance: utils.RandomMoney(),
// 	}

// 	user2, err := testQueries.Updateuser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.Username, user2.Username)
// 	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
// 	require.Equal(t, user1.FullName, user2.FullName)
// 	require.Equal(t, user1.Email, user2.Email)
// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }

// func TestDeleteuser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	err := testQueries.Deleteuser(context.Background(), user1.Username)
// 	require.NoError(t, err)

// 	user2, err := testQueries.Getuser(context.Background(), user1.Username)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, user2)
// }

// func TestListusers(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomUser(t)
// 	}

// 	arg := ListusersParams{
// 		Limit: 5, // skip the first 5 records
// 		Offset: 5, // return the next 5
// 	}

// 	users, err := testQueries.Listusers(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, users, 5)

// 	for _, user := range users {
// 		require.NotEmpty(t, user)
// 	}
// }