package models
import (
  "time"
  //"github.com/jinzhu/gorm"
)

type Members struct {
  //gorm.Model
  Id                int       `json:"id"`
  First_name        string    `json:"first_name"`
  Last_name         string    `json:"last_name"`
  Address           string    `json:"address"`
  City              string    `json:"city"`
  State             string    `json:"state"`
  Zip               string    `json:"zip"`
  Home_phone        string    `json:"home_phone"`
  Cell_phone        string    `json:"cell_phone"`
  Email             string    `json:"email"`
  Membership_date   time.Time `json:"membership_date"`
  Status            string    `json:"status"`
  Birth_date        string    `json:"birth_date"`
  Birth_year        string    `json:"birth_year"`
  CreatedAt         time.Time `json:"createdAt"`
  UpdatedAt         time.Time `json:"updatedAt"`
}
