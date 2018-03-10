package main

import (
  "app"
  "controller"
  "middleware"
)

func main() {
  app := app.App{}
  app.Initialize()

  // middlewares
  cors := middleware.Cors{app}
  security := middleware.Security{app}
  app.Router.Use(cors.Middleware)
  app.Router.Use(security.Middleware)

  // controllers
  users    := controller.Users{app}
  accounts := controller.Accounts{app}

  // routes
  app.Router.HandleFunc("/users", users.Index).Methods("GET", "OPTIONS")
  app.Router.HandleFunc("/users/{id}", users.View).Methods("GET", "OPTIONS")

  app.Router.HandleFunc("/token", accounts.Token).Methods("POST", "OPTIONS")
  app.Router.HandleFunc("/token", accounts.Logout).Methods("DELETE")
  app.Router.HandleFunc("/accounts", accounts.New).Methods("POST", "OPTIONS")

  // start app
  app.Run()
}
