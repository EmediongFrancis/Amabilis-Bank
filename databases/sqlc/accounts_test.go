package databases

import (
	"context"
	"testing"
	"time"

	"github.com/EmediongFrancis/Amabilis-Bank/util"
	"github.com/stretchr/testify/require"
)

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

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

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
