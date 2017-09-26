package main

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
   "./models"
)

var db *sql.DB

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
      adrs := new(models.Address)
      err := rows.Scan(&adrs.PersonId, &adrs.Street, &adrs.City, &adrs.State, &adrs.Zip)
      if err != nil {
         panic(err.Error())
         return
      }
      fmt.Println(adrs)
    }
}
