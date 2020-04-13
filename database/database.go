package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "time"
)

func connect(connection string) *sql.DB {

  // Create connection
  db, err := sql.Open("mysql", connection)
  if err != nil {
    panic(err)
  }

  // check connection
  err = db.Ping()
  if err != nil {
    panic(err)
  }

  return db
}

func queryRows(db *sql.DB) {
  rows, err := db.Query("select id, username from Accounts")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  for rows.Next() {
    var (
      id       uint
      username string
    )
    err = rows.Scan(&id, &username)
    if err != nil {
      panic(err)
    }

    fmt.Printf("id = %d, username = %s\n", id, username)
  }
}

func queryRows2(db *sql.DB, stm *sql.Stmt) {
  rows, err := stm.Query(1, "user_2")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  for rows.Next() {
    var (
      id       uint
      username string
    )
    err = rows.Scan(&id, &username)
    if err != nil {
      panic(err)
    }

    fmt.Printf("id = %d, username = %s\n", id, username)
  }
}

func queryRow(db *sql.DB) {
  var (
    id       uint
    username string
  )

  row := db.QueryRow("select id, username from Accounts where id = ?", 3)

  fmt.Println("Passed", row)

  err := row.Scan(&id, &username)
  if err != nil {
    panic(err)
  }

  fmt.Printf("id = %d, username = %s\n", id, username)
}

func insertRow(db *sql.DB) {
  stm, err := db.
    Prepare("insert into Accounts(username, password_hash, state, created_date, updated_date) values(?, ?, ?, ?, ?)")
  if err != nil {
    panic(err)
  }
  defer stm.Close()

  row, err := stm.Exec("nguduong", "my_password_hash", "ENABLED", time.Now(), time.Now())
  if err != nil {
    panic(err)
  }

  fmt.Println(row)
}

func createTransaction(db *sql.DB) (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()

  stm, err := tx.
    Prepare("insert into Accounts(username, password_hash, state, created_date, updated_date) values(?, ?, ?, ?, ?)")
  if err != nil {
    return
  }

  for i := 0; i < 2; i++ {
    name := fmt.Sprintf("Account %d", i)
    _, err = stm.Exec(name, "password", "ENABLED", time.Now(), time.Now())
    if err != nil {
      return
    }
  }

  err = tx.Commit()
  if err != nil {
    return
  }

  stm.Close()
  return
}

func createTransactionFail(db *sql.DB) (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()

  stm, err := tx.
    Prepare("insert into Accounts(username, password_hash, state, created_date, updated_date) values(?, ?, ?, ?, ?)")
  if err != nil {
    return
  }

  for i := 0; i < 2; i++ {
    name := fmt.Sprintf("Account %d", i)
    _, err = stm.Exec(name, "password", "ENABLED", time.Now(), time.Now())
    if err != nil {
      return
    }
    // simulate break here
    if i == 1 {
      fmt.Println("Break here ! Rollback")
      return
    }
  }

  err = tx.Commit()
  if err != nil {
    return
  }

  stm.Close()
  return
}

func main() {
  // change dev:root to your local credential
  db := connect("dev:root@tcp(127.0.0.1:3306)/contacts")
  defer db.Close()
  fmt.Println("Connected !")

  stm, err := db.Prepare("select id, username from Accounts where id = ? and username = ?")
  if err != nil {
    panic(err)
  }
  defer stm.Close()

  queryRow(db)
  // insertRow(db)

  // createTransactionFail(db)
}
