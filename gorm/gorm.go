package main

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "time"
)

type Account struct {
  ID           int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  Username     string    `gorm:"type:VARCHAR(256);NOT NULL"`
  PasswordHash string    `gorm:"type:VARCHAR(64);NOT NULL"`
  State        string    `gorm:"type:VARCHAR(64);NOT NULL"`
  CreatedAt    time.Time `gorm:"NOT NULL"`
  UpdatedAt    time.Time `gorm:"NOT NULL"`
  Posts        []Post    `gorm:"foreignkey:AccountId"`
}

type Post struct {
  ID        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
  AccountId int       `gorm:"NOT NULL"`
  Content   string    `gorm:"NOT NULL"`
  CreatedAt time.Time `gorm:"NOT NULL"`
  UpdatedAt time.Time `gorm:"NOT NULL"`
}

func connect(connection string) *gorm.DB {

  // Create connection
  db, err := gorm.Open("mysql", connection)
  if err != nil {
    panic(err)
  }

  err = db.AutoMigrate(&Account{}, &Post{}).Error
  if err != nil {
    panic(err)
  }
  db.Model(&Post{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")

  return db
}

func main() {
  db := connect("dev:root@tcp(127.0.0.1:3306)/contacts?parseTime=true")
  defer db.Close()
  fmt.Println("Connected !")

  // create a new account
  // acc1 := Account{
  //   Username:     "user_1",
  //   PasswordHash: "password_1",
  //   State:        "ENABLED",
  // }
  // err := db.Create(&acc1).Error
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }

  // Get an account
  var readAcc Account
  err := db.Preload("Posts").First(&readAcc, 1).Error
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(readAcc)

  // Create a post
  post2 := Post{Content: "A new post for user 1!"}
  err = db.Model(&readAcc).Association("Posts").Append(post2).Error
  if err != nil {
    fmt.Println(err)
    return
  }

  // // Update the account
  // readAcc.Username = "user_updated"
  // readAcc.PasswordHash = "new_password"
  // err = db.Save(&readAcc).Error
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }

  // Get posts related to account
  // var posts []Post
  // err = db.Model(&readAcc).Related(&posts).Error
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }
  // fmt.Println(posts)

  // // Delete an account
  var post3 Post
  err := db.Where("id = ?", 3).First(&post3).Error
  if err != nil {
    fmt.Println(err)
    return
  }
  db.Delete(&post3)
}
