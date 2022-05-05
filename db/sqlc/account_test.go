package db

import (
	"context"
	"database/sql"
	"github.com/rifqiahmad/crispy-engine/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createAccountSample(arg CreateAccountParams) (Account, error) {
	account, err := testQueries.CreateAccount(context.Background(), arg)

	return account, err
}

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
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
}

func TestGetAccount(t *testing.T) {
	createArg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	accountCreated, err := createAccountSample(createArg)

	accountRetrieve, err := testQueries.GetAccount(context.Background(), accountCreated.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accountRetrieve)

	require.Equal(t, accountCreated.ID, accountRetrieve.ID)
	require.Equal(t, accountCreated.Owner, accountRetrieve.Owner)
	require.Equal(t, accountCreated.Balance, accountRetrieve.Balance)
	require.Equal(t, accountCreated.Currency, accountRetrieve.Currency)
	require.WithinDuration(t, accountCreated.CreatedAt, accountRetrieve.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createArg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	accountCreated, err := createAccountSample(createArg)

	accountRetrieve, err := testQueries.GetAccount(context.Background(), accountCreated.ID)

	money := util.RandomMoney()
	updateArg := UpdateAccountParams{
		ID:      accountRetrieve.ID,
		Balance: money,
	}
	accountUpdated, err := testQueries.UpdateAccount(context.Background(), updateArg)

	require.NoError(t, err)
	require.NotEmpty(t, accountUpdated)

	require.Equal(t, accountRetrieve.ID, accountUpdated.ID)
	require.Equal(t, accountRetrieve.Owner, accountUpdated.Owner)
	require.Equal(t, updateArg.Balance, accountUpdated.Balance)
	require.Equal(t, accountRetrieve.Currency, accountUpdated.Currency)
	require.WithinDuration(t, accountRetrieve.CreatedAt, accountUpdated.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createArg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	accountCreated, _ := createAccountSample(createArg)

	err := testQueries.DeleteAccount(context.Background(), accountCreated.ID)

	accountRetrieve, errorRetrieve := testQueries.GetAccount(context.Background(), accountCreated.ID)

	require.NoError(t, err)
	require.EqualError(t, errorRetrieve, sql.ErrNoRows.Error())
	require.Empty(t, accountRetrieve)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createArg := CreateAccountParams{
			Owner:    util.RandomOwner(),
			Balance:  util.RandomMoney(),
			Currency: util.RandomCurrency(),
		}

		createAccountSample(createArg)
	}

	var limit int32 = 5
	var offset int32 = 5
	arg := ListAccountsParams{
		Limit:  limit,
		Offset: offset,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, int(limit))

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
