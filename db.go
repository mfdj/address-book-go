package main

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
)

var db *sql.DB

type Address struct {
   PersonId int
   Street   string
   City     string
   State    string
   Zip      string
}

func main() {
   db, err := sql.Open("mysql", "goldie:gopher@/address_book_go")
   defer db.Close()

   if err != nil {
       panic(err.Error())
   }

   err = db.Ping()
   if err != nil {
      panic(err.Error())
   }

   rows, err := db.Query("SELECT person_id, street, city, state, zip FROM address")
   if err != nil {
      return
   }
   defer rows.Close()
   
   for rows.Next() {
      adrs := new(Address)
      err := rows.Scan(&adrs.PersonId, &adrs.Street, &adrs.City, &adrs.State, &adrs.Zip)
      if err != nil {
         panic(err.Error())
         return
      }
      fmt.Println(adrs)
    }
}
