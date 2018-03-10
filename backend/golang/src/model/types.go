package model

type User struct {
  Id         string   `json:"id"`
  Username   string   `json:"username"`
  Name       string   `json:"name"`
}

type Account struct {
  Username   string   `json:"username"`
  Password   string   `json:"password"`
}

type Token struct {
  Username   string   `json:"username"`
  Token      string   `json:"token"`
}
