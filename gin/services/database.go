package services

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "golang-samples/gin/models"
)

var DB *gorm.DB

func Connect(connection string) (err error) {

  // Create connection
  DB, err = gorm.Open("mysql", connection)
  if err != nil {
    return
  }

  err = DB.AutoMigrate(&models.Account{}, &models.Post{}).Error
  if err != nil {
    panic(err)
  }
  DB.Model(&models.Post{}).
    AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")

  return
}

func GetAccountByID(id string) (account *models.Account, err error) {
  account = &models.Account{}
  err = DB.First(account, id).Error
  return
}

func GetPostsByAccount(accountId string) (posts *[]models.Post, err error) {
  posts = new([]models.Post)
  err = DB.Find(posts).Where("account_id = ?", accountId).Error
  return
}
