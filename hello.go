package main
import (
  "fmt"
  "net/http"
  //"log"
  //"github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func YourHandler(w http.ResponseWriter, r *http.Request) {



  w.Write([]byte("Gorilla!\n"))
}

func main() {
  fmt.Println("Starting Server...")

    // Connect to PSQL DB (GORM)
  db, err := gorm.Open("postgres", "sslmode=disable user=postgres dbname=frcc password=")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  //db.AutoMigrate(&Members{})

  //var member Member


  rows, err := db.Raw(`
    SELECT
      id,
      first_name,
      last_name,
      address,
      city,
      state,
      zip,
      home_phone,
      cell_phone,
      email,
      membership_date,
      status,
      birth_date,
      birth_year,
      "createdAt",
      "updatedAt"
    FROM "Members"
  `).Rows()
  defer rows.Close()
  //var members []Members
  for rows.Next() {
    var member Members
    rows.Scan(
      &member.id,
      &member.first_name,
      &member.last_name,
      &member.address,
      &member.city,
      &member.state,
      &member.zip,
      &member.home_phone,
      &member.cell_phone,
      &member.email,
      &member.membership_date,
      &member.status,
      &member.birth_date,
      &member.birth_year,
      &member.createdAt,
      &member.updatedAt)

    fmt.Println(member.id)

    fmt.Println("----")
  }


  //fmt.Println(member.id)









  //r := mux.NewRouter()
  // Routes consist of a path and a handler function.
  //r.HandleFunc("/", YourHandler)

  // Bind to a port and pass our router in
  //log.Fatal(http.ListenAndServe(":8000", r))
}






// Connect to PSQL DB (go-pg)
// db := pg.Connect(&pg.Options{
//   User:     "postgresss",
//   Password: "",
//   Database: "postgres",
// })
// defer db.Close()



// View All Tables
// rows, err := db.Raw("SELECT table_name from information_schema.tables").Rows()
// for rows.Next() {
//   var tablename string
//   rows.Scan(&tablename)
//   fmt.Println(tablename)
// }

/*
rows, err := db.Raw("SELECT id FROM \"Members\"").Rows()
if err != nil {
  w.Write([]byte("Error Retrieving Members"))
}
defer rows.Close()
// var id                int
// var first_name        string
// var last_name         string
// var address           string
// var city              string
// var state             string
// var zip               string
// var home_phone        string
// var cell_phone        string
// var email             string
// var membership_date   time.Time
// var status            string
// var birth_date        string
// var birth_year        string
// var createdAt         time.Time
// var updatedAt         time.Time
// for rows.Next() {
//   //var member Members
//   rows.Scan(&id, &first_name, &last_name, &address, &city, &state, &zip, &home_phone, &cell_phone, &email, &membership_date, &status, &birth_date, &birth_year, &createdAt, &updatedAt)
//   fmt.Println(id)
// }

w.Write([]byte("Success!"))
*/
