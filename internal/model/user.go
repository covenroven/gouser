package model

type User struct {
    Id int `json: "id"`
    Name string `json: "name"`
    Email string `json: "email"`
    Addresses []Address `json: "addresses"`
}

type Address struct {
    Id int `json: "id"`
    Street string `json: "street"`
    City string `json: "city"`
    Province string `json: "province"`
    PostalCode string `json: "postal_code"`
    Country string `json: "country"`
}
