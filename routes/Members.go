package routes
import (
  //"fmt"
  "net/http"
  "encoding/json"
  //"github.com/jinzhu/gorm"
  "app/models"
)

func GetAllMembers(w http.ResponseWriter, r *http.Request) {
  rows, err := models.DB.Raw(`
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
  if err != nil {
    panic(err.Error())
  }
  var members []models.Members
  for rows.Next() {
    var member models.Members
    rows.Scan(
      &member.Id,
      &member.First_name,
      &member.Last_name,
      &member.Address,
      &member.City,
      &member.State,
      &member.Zip,
      &member.Home_phone,
      &member.Cell_phone,
      &member.Email,
      &member.Membership_date,
      &member.Status,
      &member.Birth_date,
      &member.Birth_year,
      &member.CreatedAt,
      &member.UpdatedAt)
    members = append(members, member)
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(members)

  // js, err := json.Marshal(members)
  // if err != nil {
  //   http.Error(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }
  // w.Header().Set("Content-Type", "application/json")
  // w.Write(js)
}
