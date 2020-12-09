package service

import (
    "log"
    "fmt"
    "bytes"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"
    "encoding/json"
    "github.com/covenroven/gouser/internal/model"
)

type AddressResponse struct {
    Status int
    Message string
    Data []model.Address
}

type AddressRequest struct {
    UserId int `json: "user_id"`
    Addresses []model.Address `json: "addresses"`
}

var AddressUrl = "http://localhost:3100"

func GetAddressOfUser(userID int) ([]model.Address, error) {
    res, err := http.Get(AddressUrl + "/addresses?user_id=" + strconv.Itoa(userID))
    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    return ConvertResponseToAddress(res.Body)
}

func storeAddressOfUser(userID int, addresses []model.Address) ([]model.Address, error) {
    body, err := json.Marshal(AddressRequest{
        UserId: userID,
        Addresses: addresses,
    })
    if err != nil {
        return nil, err
    }

    res, err := http.Post(AddressUrl + "/addresses", "application/json", bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }

    defer res.Body.Close()

    return ConvertResponseToAddress(res.Body)
}

func ConvertResponseToAddress(body io.ReadCloser) ([]model.Address, error) {
    data, err := ioutil.ReadAll(body)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var addressRes AddressResponse
    json.Unmarshal(data, &addressRes)

    fmt.Println("%v", addressRes)
    fmt.Println("%v", addressRes.Data)
    return addressRes.Data, nil
}
