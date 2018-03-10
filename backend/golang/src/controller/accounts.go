package controller

import (
  "fmt"
  "app"
  "model"
  "net/http"
  "crypto/rand"
  "crypto/sha256"
  "golang.org/x/crypto/bcrypt"
)

type Accounts struct {
  App app.App
}

// Generate an access token from existing account
func (c *Accounts) Token(w http.ResponseWriter, r *http.Request) {
  params := model.Account{}
  c.App.Parse(r, &params)

  account := model.Account{}
  row := c.App.DB.QueryRow("SELECT a.username, a.password FROM `accounts` a WHERE a.username = ?", params.Username)
  err := row.Scan(&account.Username, &account.Password)

  if err != nil {
    c.App.NotFound(w)
    return
  }

  err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(params.Password))

  if err != nil {
    c.App.Error(w, "Password incorrect", http.StatusBadRequest)
    return
  }

  // remove current token (if exists)
  c.App.DB.Exec("DELETE FROM `tokens` WHERE `username` = ?", account.Username)

  // generate token hash
  data := make([]byte, 10)
  rand.Read(data)
  hash := fmt.Sprintf("%x", sha256.Sum256(data))

  // store token on database
  token := model.Token{account.Username, hash}
  stmt, err := c.App.DB.Prepare("INSERT INTO `tokens` (`username`, `token`) VALUES (?, ?)")
  defer stmt.Close()

  if err != nil {
    c.App.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  _, err = stmt.Exec(token.Username, token.Token)

  if err != nil {
    c.App.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // success
  c.App.Json(w, token, http.StatusOK)
}

// Create a new account
func (c *Accounts) New(w http.ResponseWriter, r *http.Request) {
  account := model.Account{}
  c.App.Parse(r, &account)

  if account.Username == "" {
    c.App.Error(w, "Please fill the account username", http.StatusBadRequest)
    return
  }

  if account.Password == "" {
    c.App.Error(w, "Please fill the account password", http.StatusBadRequest)
    return
  }

  stmt, err := c.App.DB.Prepare("INSERT INTO `accounts` (`username`, `password`) VALUES (?, ?)")
  defer stmt.Close()

  if err != nil {
    c.App.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
  if err != nil {
    c.App.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  account.Password = string(hashedPassword)

  _, err = stmt.Exec(account.Username, account.Password)

  if err != nil {
    c.App.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  // response (prevent to return encoded password)
  response := map[string]bool{"success": true}

  c.App.Json(w, response, http.StatusOK)
}

// Remove a generated token
func (c *Accounts) Logout(w http.ResponseWriter, r *http.Request) {
  params := model.Token{}
  c.App.Parse(r, &params)

  c.App.DB.Exec("DELETE FROM `tokens` WHERE `username` = ? AND `token` = ?", params.Username, params.Token)

  c.App.Json(w, params, http.StatusOK)
}
