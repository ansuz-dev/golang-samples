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

func SaveAccount(account *models.Account) (err error) {
  err = DB.Save(account).Error
  return
}

func GetPostByID(id string) (post *models.Post, err error) {
  post = &models.Post{}
  err = DB.First(post, id).Error
  return
}

func GetPostsByAccount(accountId string) (posts *[]models.Post, err error) {
  posts = new([]models.Post)
  err = DB.Where("account_id = ?", accountId).Find(posts).Error
  return
}

func SavePost(post *models.Post) (err error) {
  err = DB.Save(post).Error
  return
}
