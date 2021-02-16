package form3clientapi

import (
    "errors"
    "fmt"
    "github.com/google/uuid"
    "testing"
)

func TestCreate(t *testing.T) {
    initialCount, _ := countAccounts(10000)
    fmt.Printf("initialCount %v\n", initialCount)

    _, err := createAccount()
    if err != nil {
        t.Errorf("failed to create %v", err)
        return
    }

    count, err := countAccounts(10000)
    if err != nil {
        t.Errorf("failed to count accounts %v", err)
        return
    }

    if count != initialCount+ 1 {
        t.Errorf("failed to create account")
    }
}

func TestFetchAccount(t *testing.T) {
    createdAccount, _ := createAccount()

    account, err := fetchAccount(createdAccount.Id)
    if err != nil {
        t.Errorf("failed to create and fetch account %v",  createdAccount.Id)
        return
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