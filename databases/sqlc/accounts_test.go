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
