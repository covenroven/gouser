package api

import (
    "net/http"
    "log"
    // "database/sql"
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
	w.Write([]byte("post user"))
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

