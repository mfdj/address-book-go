package main

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "fmt"
)

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
}
