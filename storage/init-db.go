package storage

import(
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func InitDb(dataSourceName string) *sql.DB {
   db, err := sql.Open("mysql", dataSourceName)

   if err != nil {
      panic(err.Error())
   }

   err = db.Ping()
   if err != nil {
      panic(err.Error())
   }

   return db
}
