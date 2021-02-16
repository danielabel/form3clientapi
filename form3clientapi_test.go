package form3clientapi

import "testing"


func TestCreate(t *testing.T) {
    accountId, _ := createAccount();
    account, _ := getAccount(accountId)
    if account.Id != accountId {
        t.Errorf("did not get account ID = %v, want %v", account.Id, accountId)
    }
}