package main

import (
   "./storage"
   "fmt"
   "./models"
)

func main() {
   db := storage.InitDb("goldie:gopher@/address_book_go")
   defer db.Close()

   rows, err := db.Query("SELECT person_id, street, city, state, zip FROM address")
   if err != nil {
      panic(err.Error())
   }
   defer rows.Close() 
   
   for rows.Next() {
      adrs := new(models.Address)
      err := rows.Scan(&adrs.PersonId, &adrs.Street, &adrs.City, &adrs.State, &adrs.Zip)
      if err != nil {
         panic(err.Error())
      }
      fmt.Println(adrs)
    }
}
