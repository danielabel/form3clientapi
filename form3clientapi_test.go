package form3clientapi

import (
    "errors"
    "github.com/google/uuid"
    "strings"
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

    t.Run("Passes on API error so user can fix", func(t *testing.T) {
        createdAccount, err := createAccount(uuid.New(), "USA")

        if err == nil {
            t.Errorf("Failed to pass back error %v", createdAccount)
            return
        }

        // we want to check that it passes on details about bad params
        if !strings.Contains(err.Error(), "country in body should match '^[A-Z]{2}$") {
            t.Errorf("failed to report failure well %s", err.Error())
        }
    })
}

func TestFetchAccount(t *testing.T) {
    t.Run("fetches existing account", func(t *testing.T) {
        createdAccount, _ := createAccountWithDefaults()

        account, err := fetchAccount(createdAccount.Id)
        if err != nil {
            t.Errorf("failed to create and fetch account %v, %s", createdAccount.Id, err.Error())
            return
        }
        if account.Id != createdAccount.Id {
            t.Errorf("did not get account ID = %v, want %v", account.Id, createdAccount.Id)
        }
    })

    t.Run("reports failure for nonexistent account", func(t *testing.T) {
        _, err := fetchAccount(uuid.New())
        if !errors.Is(err, ErrAccountDoesNotExist) {
            t.Errorf("should have failed to get non-existant account")
        }
    })
}

func TestDelete(t *testing.T) {
    t.Run("deletes existing account with upto date version", func(t *testing.T) {
        createdAccount, _ := createAccountWithDefaults()
        err := deleteAccount(createdAccount.Id, createdAccount.Version)
        if err != nil {
            t.Errorf("failed to delete %v", err)
            return
        }
        _, err = fetchAccount(createdAccount.Id)
        if !errors.Is(err, ErrAccountDoesNotExist) {
            t.Errorf("account %v should not exist after being deleted", createdAccount.Id)
        }
    })

    // note that the test api does not return 404 for "Specified resource does not exist"
    //   instead it returns a 204 (as it does for success)
    t.Run("Reports on non-existing account id helpfully", func(t *testing.T) {
       err := deleteAccount(uuid.New(), 0)
       if err != nil {
           t.Errorf("reject delete of non existing account id - unexpected")
           return
       }
    })

    // note that the test api does not return 409 for "Specified version incorrect"
    //   it instead returns a 404
    t.Run("Reports on rejected bad version usefully", func(t *testing.T) {
        createdAccount, _ := createAccountWithDefaults()
        err := deleteAccount(createdAccount.Id, createdAccount.Version-1)
        if err == nil {
            t.Errorf("failed to reject delete")
            return
        }

        // we want to check that it passes on details about the invalid version
        if !strings.Contains(err.Error(), "invalid version") {
            t.Errorf("failed to report failure well")
        }
    })
}


