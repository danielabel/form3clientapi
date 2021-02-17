package form3clientapi

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "io/ioutil"
    "log"
    "net/http"
)

type attributes struct {
    Country string `json:"country"`
}

type account struct {
    Type string `json:"type"`
    Id uuid.UUID `json:"id"`
    OrganisationId uuid.UUID `json:"organisation_id"`
    Attributes attributes `json:"attributes"`
    Version int32 `json:"version"`
}

type payload struct {
    Data account `json:"data"`
}

type accountsPayload struct {
    Data []account `json:"data"`
}

// should it be unit tested?
func extractErrorMessage(body []byte) interface{} {
    var v map[string]interface{}
    err := json.Unmarshal(body, &v); if err != nil {
        return ""
    }
    return v["error_message"]
}

var ErrAccountDoesNotExist = errors.New("account does not exist")
var ErrOperationFailed = errors.New("operation failed")

func createAccount(orgId uuid.UUID, country string) (account, error) {

    p := payload{Data: account{
        Type:           "accounts",
        Id:             uuid.New(),
        OrganisationId: orgId,
        Attributes:     attributes{
            Country: country,
        }}}

    requestBody, err := json.Marshal(p); if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return account{}, ErrOperationFailed
    }

    resp, err := http.Post(getBaseUrl()+ "/organisation/accounts",
        "application/vnd.api+json",
        bytes.NewBuffer(requestBody))
    if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return account{}, ErrOperationFailed
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return account{}, ErrOperationFailed
    }

    if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
        message := extractErrorMessage(body)
        return account{}, fmt.Errorf("validation failure: %s", message)
    }

    var s payload
    err = json.Unmarshal(body, &s); if err != nil {
        return account{}, ErrOperationFailed
    }

    return s.Data, nil
}

func fetchAccount(id uuid.UUID) (account, error) {
    resp, err := http.Get(getBaseUrl() + "/organisation/accounts/" + id.String()); if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return account{}, ErrOperationFailed

    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return account{}, ErrAccountDoesNotExist
    }

    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return account{}, ErrOperationFailed
    }

    var s payload
    err = json.Unmarshal(body, &s); if err != nil {
        return account{}, errors.New("fail Unmarshal")
    }

    return s.Data, nil
}

func countAccounts(pageSize int) (int, error) {
    url := fmt.Sprintf("%s/organisation/accounts?page[size]=%d", getBaseUrl(), pageSize)
    resp, err := http.Get(url)
    if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return 0, ErrOperationFailed
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return 0, ErrOperationFailed
    }

    var s accountsPayload
    err = json.Unmarshal(body, &s)
    if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return 0, ErrOperationFailed
    }

    return len(s.Data), nil
}

func deleteAccount(id uuid.UUID, version int32) error {
    url := fmt.Sprintf("%s/organisation/accounts/%s?version=%d", getBaseUrl(), id.String(), version)
    client := &http.Client{}
    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return ErrOperationFailed
    }
    req.Close = true

    resp, err := client.Do(req); if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return ErrOperationFailed
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        log.Printf("Operation failed. err: %v\n", err)
        return ErrOperationFailed
    }

    // note that the behaviour of the test api differs from the documented API behaviour
    // here we can extract the error from the body and report to calling code using that than
    // having expectations around response codes
    if resp.StatusCode != http.StatusNoContent {
        message := extractErrorMessage(body)
        return fmt.Errorf("delete failed: %s", message)
    }

    return nil
}
