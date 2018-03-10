package app

import (
  "os"
  "log"
  "fmt"
  "model"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "database/sql"
  "github.com/gorilla/mux"
)

const (
  APP_DSN = "APP_DSN"
  APP_ADDR = "APP_ADDR"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

// Initialize application router and database
func (a *App) Initialize() {
  var err error

  dsn := getEnvVar(APP_DSN)
  a.DB, err = model.NewDB(dsn)

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  a.Router = mux.NewRouter()
}

// Start to listen HTTP requests
func (a *App) Run() {
  addr := getEnvVar(APP_ADDR)

  log.Print(fmt.Sprintf("Server started listen on %v", addr))
  log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a App) Parse(req *http.Request, t interface{}) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    //panic(err)
  }
  err = json.Unmarshal(body, &t)
  if err != nil {
    //panic(err)
  }
}

func (a App) Error(w http.ResponseWriter, message string, code int) {
  a.Json(w, map[string]string{"error": message}, code)
}

func (a App) NotFound(w http.ResponseWriter) {
  a.Json(w, map[string]string{"error": "Not found"}, http.StatusNotFound)
}

func (a App) Json(w http.ResponseWriter, payload interface{}, code int) {
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}

func getEnvVar(varname string) string {
  value := os.Getenv(varname)
  if value == "" {
    log.Fatal(fmt.Sprintf("Please fill the %v environment variable", varname))
    os.Exit(1)
  }
  return value
}
