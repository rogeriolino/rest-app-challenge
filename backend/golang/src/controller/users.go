package controller

import (
  "fmt"
  "app"
  "model"
  "net/http"
  "github.com/gorilla/mux"
)

type Users struct {
  App app.App
}

func (c *Users) Index(w http.ResponseWriter, r *http.Request) {
  db        := c.App.DB
  users     := []model.User{}
  keyword   := r.URL.Query().Get("keyword")
  since     := r.URL.Query().Get("since")
  search    := fmt.Sprintf("%v%%", keyword)

  query := `
    select
      BIN_TO_UUID(u.id),
      u.username,
      u.name
    from
      users u
      join (
        select id from users
        order by username
      ) u2 on u2.id = u.id
    where
      (? = '' OR username > ?) AND
      (u.username LIKE ? OR u.name LIKE ?)
    limit 10
  `

  rows, err := db.Query(query, since, since, search, search)
  defer rows.Close()

  if err != nil {
    c.App.Error(w, err.Error(), 500)
	}

  for rows.Next() {
    user := model.User{}
    rows.Scan(&user.Id, &user.Username, &user.Name)
    users = append(users, user)
  }

  c.App.Json(w, users, http.StatusOK)
}


func (c *Users) View(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  userId := params["id"]

  var user model.User
  row := c.App.DB.QueryRow("SELECT BIN_TO_UUID(u.id), u.username, u.name FROM users u WHERE u.id = UUID_TO_BIN(?)", userId)
  user = model.User{}
  err := row.Scan(&user.Id, &user.Username, &user.Name)

  if err != nil {
    c.App.NotFound(w)
    return
  }

  c.App.Json(w, user, http.StatusOK)
}
