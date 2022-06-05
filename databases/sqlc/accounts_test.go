package databases

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/EmediongFrancis/Amabilis-Bank/util"
	"github.com/stretchr/testify/require"
)

// Create a random account.
func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// Test account creation.
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

// Test account querying.
func TestReadAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	accountRead, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accountRead)
	require.Equal(t, account.ID, accountRead.ID)
	require.Equal(t, account.Owner, accountRead.Owner)
	require.Equal(t, account.Balance, accountRead.Balance)
	require.Equal(t, account.Currency, accountRead.Currency)
	require.WithinDuration(t, account.CreatedAt, accountRead.CreatedAt, time.Second)
}

// Test account updating.
func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomBalance(),
	}

	accountUpdated, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accountUpdated)

	require.Equal(t, account.ID, accountUpdated.ID)
	require.Equal(t, account.Owner, accountUpdated.Owner)
	require.Equal(t, arg.Balance, accountUpdated.Balance)
	require.Equal(t, account.Currency, accountUpdated.Currency)
	require.WithinDuration(t, account.CreatedAt, accountUpdated.CreatedAt, time.Second)
}

// Test account deletion.
func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	accountRead, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountRead)
}

// Test accounts querying.
func TestReadAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accountsRead, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accountsRead, 5)

	for _, account := range accountsRead {
		require.NotEmpty(t, account)
	}
}
