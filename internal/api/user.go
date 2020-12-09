package api

import (
    "net/http"
    "log"
    "fmt"
    "encoding/json"
	"github.com/go-chi/chi"
    "github.com/covenroven/gouser/internal/database"
    "github.com/covenroven/gouser/internal/model"
    "github.com/covenroven/gouser/internal/service"
)

func IndexUsers(w http.ResponseWriter, r *http.Request) {
    db, err := database.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var users []model.Model
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var user model.User
        rows.Scan(&user.Id, &user.Name, &user.Email)

        users = append(users, user)
    }

    response := model.Response{
        Status: 200,
        Message: "ok",
        Data: users,
    }

    responseWithJson(w, response)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    if r.Body == nil {
        responseWithJson(w, model.Response{
            Status: 422,
            Message: "No parameter provided",
            Data: []model.Model{},
        })
    }

    var param model.User
    err := json.NewDecoder(r.Body).Decode(&param)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 400,
            Message: err.Error(),
            Data: []model.Model{},
        })
    }

    var user model.User
    err = db.QueryRow(
        "INSERT INTO users(name, email) VALUES($1, $2) RETURNING id, name, email",
        param.Name,
        param.Email,
    ).Scan(&user.Id, &user.Name, &user.Email)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(user)
    responseWithJson(w, model.Response{
        Status: 201,
        Message: "Created",
        Data: []model.Model{user},
    })
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()

    userID := chi.URLParam(r, "userID")

    row := db.QueryRow("SELECT * FROM users WHERE id = $1", userID)

    var user model.User
    err := row.Scan(&user.Id, &user.Name, &user.Email)
    if err != nil {
        responseWithJson(w, model.Response{
            Status: 404,
            Message: "Not found",
            Data: []model.Model{},
        })
        return
    }

    user.Addresses, _ = service.GetAddressOfUser(user.Id)

    responseWithJson(w, model.Response{
        Status: 200,
        Message: "ok",
        Data: []model.Model{user},
    })
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    db, _ := database.Connect()
    defer db.Close()
}
