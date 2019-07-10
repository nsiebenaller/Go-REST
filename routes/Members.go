package main
import (
  "time"
)

type Members struct {
  gorm.Model
  id                int
  first_name        string
  last_name         string
  address           string
  city              string
  state             string
  zip               string
  home_phone        string
  cell_phone        string
  email             string
  membership_date   time.Time
  status            string
  birth_date        string
  birth_year        string
  createdAt         time.Time
  updatedAt         time.Time
}
