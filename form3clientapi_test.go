package form3clientapi

import (
    "errors"
    "github.com/google/uuid"
    "testing"
)

func createAccountWithDefaults() (account, error) {
    return createAccount(uuid.New(), "UK")
}

func TestCreate(t *testing.T) {
    t.Run("creates an account", func(t *testing.T) {

        // what's the starting number of accounts?
        initialCount, _ := countAccounts(10000)

        // create a new account
        _, err := createAccountWithDefaults(); if err != nil {
            t.Errorf("failed to create %v", err)
            return
        }

        // do we get one more account
        count, err := countAccounts(10000); if err != nil {
            t.Errorf("failed to count accounts %v", err)
            return
        }
        if count != initialCount+1 {
            t.Errorf("failed to create account")
        }
    })

    t.Run("returns account details including Id", func(t *testing.T) {
        createdAccount, err := createAccountWithDefaults(); if err != nil {
            t.Errorf("failed to create %v", err)
            return
        }

        if createdAccount.Id == uuid.Nil {
            t.Error("failed to set account Id")
        }
        if createdAccount.Type != "accounts" {
            t.Error("failed to set type to account")
        }
    })

    t.Run("sets org Id to what we ask it to be ", func(t *testing.T) {
        // create a new account
        orgId := uuid.New()
        createdAccount, err := createAccount(orgId, "UK"); if err != nil {
            t.Errorf("failed to create %v", err)
            return
        }

        if createdAccount.OrganisationId != orgId {
            t.Errorf("failed to set org Id (%v) to what we asked (%v)", createdAccount.OrganisationId, orgId)
        }
    })

    t.Run("sets country to what we ask it to be ", func(t *testing.T) {
       country := "US"
       createdAccount, err := createAccount(uuid.New(), country)
       if err != nil {
           t.Errorf("failed to create %v", err)
           return
       }

       if createdAccount.Attributes.Country != country {
           t.Errorf("failed to set country (%v) to what we asked (%v)", createdAccount.Attributes.Country, country)
       }
    })
}

func TestFetchAccount(t *testing.T) {
    createdAccount, _ := createAccountWithDefaults()

    account, err := fetchAccount(createdAccount.Id); if err != nil {
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
    createdAccount, _ := createAccountWithDefaults()
   err := deleteAccount(createdAccount.Id); if err != nil {
        t.Errorf("failed to delete %v", err)
    }
   _, err = fetchAccount(createdAccount.Id)
   if !errors.Is(err, ErrAccountDoesNotExist) {
       t.Errorf("account %v should not exist after being deleted", createdAccount.Id)
   }
}