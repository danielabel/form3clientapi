package form3clientapi

import (
    "errors"
    "fmt"
    "github.com/google/uuid"
    "testing"
)

func TestCreate(t *testing.T) {
    initialCount, err := countAccounts(10000)
    fmt.Printf("initialCount %v\n", initialCount)
    createdAccount, _ := createAccount()
    count, err := countAccounts(10000)
    if count != initialCount+ 1 {
        t.Errorf("failed to create account")
    }
    account, err := fetchAccount(createdAccount.Id)
    if err != nil {
        t.Errorf("failed to create and fetch account %v",  createdAccount.Id)
    }
    if account.Id != createdAccount.Id {
        t.Errorf("did not get account ID = %v, want %v", account.Id, createdAccount.Id)
    }
}

func TestFetchNonExistingAccount(t *testing.T) {
    _, err := fetchAccount(uuid.New())
    if !errors.Is(err, ErrAccountDoesNotExist) {
        t.Errorf("should have failed to get non-existant account")
    }
}

func TestDelete(t *testing.T) {
    createdAccount, _ := createAccount()
   err := deleteAccount(createdAccount.Id)
    if err != nil {
        t.Errorf("failed to delete %v", err)
    }
   _, err = fetchAccount(createdAccount.Id)
   if !errors.Is(err, ErrAccountDoesNotExist) {
       t.Errorf("account %v should not exist after being deleted", createdAccount.Id)
   }
}