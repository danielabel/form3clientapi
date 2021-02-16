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

type attributes struct {
    Country string `json:"country"`
}

type data struct {
    Type string `json:"type"`
    Id uuid.UUID `json:"id"`
    OrganisationId uuid.UUID `json:"organisation_id"`
    Attributes attributes `json:"attributes"`
}

type payload struct {
    Data data `json:"data"`
}

func createAccount() (uuid.UUID, error) {

    p:= payload{Data: data{
        Type:           "accounts",
        Id:             uuid.New(),
        OrganisationId: uuid.New(),
        Attributes:     attributes{
            Country: "UK",
        }}}

    requestBody, err := json.Marshal(p); if err != nil {
        fmt.Printf("err %v\n", err)
        return uuid.Nil, errors.New("fail Marshal")
    }

    resp, err := http.Post("http://localhost:8080/v1/organisation/accounts",
        "application/vnd.api+json",
        bytes.NewBuffer(requestBody))
    if err != nil {
        fmt.Printf("err %v\n", err)
        return uuid.Nil, errors.New("fail Marshal")
    }
    defer resp.Body.Close()

    if resp.StatusCode != 201 {
        fmt.Printf("non 201 response %v %v", resp.Status, resp.Body)
        return uuid.Nil, errors.New("non 201 response")
    }

    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        fmt.Printf("err %v\n", err)
        return uuid.Nil, errors.New("failed to read body")
    }

    fmt.Printf("SUCESSS %v\n", string(body))

    return p.Data.Id, nil
}

func getAccount(id uuid.UUID) (data, error) {
    resp, err := http.Get("http://localhost:8080/v1/organisation/accounts/" + id.String()); if err != nil {
        fmt.Printf("err %v\n", err)
        return data{}, errors.New("empty name")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body); if err != nil {
        return data{}, errors.New("fail read")
    }

    var s payload
    err = json.Unmarshal(body, &s); if err != nil {
        return data{}, errors.New("fail Unmarshal")
    }

    return s.Data, nil
}