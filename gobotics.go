package main

import (
    "fmt"
    "database/sql"
    _"github.com/mattn/go-sqlite3"
)


func load_sql_file(db *sql.DB){
  stmt, _ := db.Prepare("create table if not exists client_devices( " +
      " name text, platform text, mac_address text, ip_address varchar(15) );" )
  _, err := stmt.Exec()
  if err != nil { panic(err) }
}

func main() {
  db, _ := sql.Open("sqlite3", "./foo.db")

  load_sql_file(db)

  fmt.Printf("%s", db)
  println("Welcome to gobotics")
}
