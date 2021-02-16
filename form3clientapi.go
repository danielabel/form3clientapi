package form3clientapi

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "io/ioutil"
    "net/http"
)

type accountsPayload struct {
    Data []account `json:"data"`
}


type attributes struct {
    Country string `json:"country"`
}

type account struct {
    Type string `json:"type"`
    Id uuid.UUID `json:"id"`
    OrganisationId uuid.UUID `json:"organisation_id"`
    Attributes attributes `json:"attributes"`
}

type payload struct {
    Data account `json:"data"`
}



var ErrAccountDoesNotExist = errors.New("account does not exist")
var ErrOperationFailed = errors.New("operation failed")

const baseURL = "http://localhost:8080/v1"

func createAccount() (account, error) {

    p:= payload{Data: account{
        Type:           "accounts",
        Id:             uuid.New(),
        OrganisationId: uuid.New(),
        Attributes:     attributes{
            Country: "UK",
        }}}

    requestBody, err := json.Marshal(p); if err != nil {
        fmt.Printf("err %v\n", err)
        return account{}, errors.New("fail Marshal")
    }

    resp, err := http.Post(baseURL + "/organisation/accounts",
        "application/vnd.api+json",
        bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Printf("err %v\n", err)
        return account{}, errors.New("fail Marshal")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        fmt.Printf("non 201 response %v %v", resp.Status, resp.Body)
        return account{}, errors.New("non 201 response")
    }

    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        fmt.Printf("err %v\n", err)
        return account{}, errors.New("failed to read body")
    }

    var s payload
    err = json.Unmarshal(body, &s); if err != nil {
        return account{}, errors.New("fail Unmarshal")
    }

    return s.Data, nil
}

func fetchAccount(id uuid.UUID) (account, error) {
    resp, err := http.Get(baseURL + "/organisation/accounts/" + id.String()); if err != nil {
        fmt.Printf("err %v\n", err)
        return account{}, errors.New("empty name")
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return account{}, ErrAccountDoesNotExist
    }

    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        return account{}, errors.New("fail read")
    }

    var s payload
    err = json.Unmarshal(body, &s); if err != nil {
        return account{}, errors.New("fail Unmarshal")
    }

    return s.Data, nil
}
func countAccounts(pageSize int) (int, error) {
    resp, err := http.Get(fmt.Sprintf("%s/organisation/accounts?page[size]=%d", baseURL, pageSize)); if err != nil {
        fmt.Printf("err %v\n", err)
        return 0, ErrOperationFailed
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        return 0, errors.New("fail read")
    }

    var s accountsPayload
    err = json.Unmarshal(body, &s); if err != nil {
        return 0, errors.New("fail Unmarshal")
    }

    return len(s.Data), nil
}

func deleteAccount(id uuid.UUID) error {
    client := &http.Client{}
    req, err := http.NewRequest("DELETE", baseURL + "/organisation/accounts/" + id.String() + "?version=0", nil)
    if err != nil {
        fmt.Printf("NewRequest err %v\n", err)
        return errors.New("NewRequest err")
    }

    resp, err := client.Do(req); if err != nil {
        fmt.Printf("err %v\n", err)
        return errors.New("empty name")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent {
        return errors.New("delete failed")
    }

    return nil
}