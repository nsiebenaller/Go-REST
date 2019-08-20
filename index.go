package main
import (
  "fmt"
  "os"
  "net/http"
  "log"
  "errors"
  "github.com/gorilla/mux"
  //"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "app/models"
  "app/routes"
  "github.com/joho/godotenv"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
  fmt.Println("Starting Server...")

  conn, err := getConnectionString()
  if(err != nil) {
    fmt.Println("Cannot get connection string")
    return
  }

  // Connect to PSQL DB (GORM)
  models.Initialize(conn)

  // Setup router
  r := mux.NewRouter()
  r.HandleFunc("/", routes.GetAllMembers)

  // Bind to a port and pass our router in
  log.Fatal(http.ListenAndServe(":8000", r))
}

func getConnectionString() (string, error) {

  conn := ""

  db_sslmode, ssl_exists := os.LookupEnv("DB_SSLMODE")
  if(ssl_exists) {
    conn = conn + "sslmode=" + db_sslmode
  } else {
    fmt.Println("Cannot find environment variable: 'DB_SSLMODE' ")
    return conn, errors.New("Cannot find environment variable: 'DB_SSLMODE' ")
  }

  db_user, user_exists := os.LookupEnv("DB_USER")
  if(user_exists) {
    conn = conn + " user=" + db_user
  } else {
    fmt.Println("Cannot find environment variable: 'DB_USER' ")
    return conn, errors.New("Cannot find environment variable: 'DB_USER' ")
  }

  db, db_exists := os.LookupEnv("DATABASE")
  if(db_exists) {
    conn = conn + " dbname=" + db
  } else {
    fmt.Println("Cannot find environment variable: 'DATABASE' ")
    return conn, errors.New("Cannot find environment variable: 'DATABASE' ")
  }

  db_pass, pass_exists := os.LookupEnv("DB_PASS")
  if(pass_exists) {
    conn = conn + " password=" + db_pass
  } else {
    fmt.Println("Cannot find environment variable: 'DB_PASS' ")
    return conn, errors.New("Cannot find environment variable: 'DB_PASS' ")
  }

  return conn, nil
}
