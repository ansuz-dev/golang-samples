package routes

import (
  "github.com/gin-gonic/gin"
  "golang-samples/gin/models"
  "golang-samples/gin/services"
)

func GetAccountByID(ctx *gin.Context) {
  // get id param
  id := ctx.Param("id")

  account, err := services.GetAccountByID(id)
  if err != nil {
    ctx.AbortWithError(400, err)
    return
  }

  ctx.JSON(200, account)
}

func UpdateAccountByID(ctx *gin.Context) {
  // get id param
  id := ctx.Param("id")

  account, err := services.GetAccountByID(id)
  if err != nil {
    ctx.AbortWithError(400, err)
    return
  }

  newAccount := models.Account{}
  if err = ctx.BindJSON(&newAccount); err != nil {
    ctx.AbortWithError(400, err)
    return
  }

  account.Username = newAccount.Username
  err = services.SaveAccount(account)
  if err != nil {
    ctx.AbortWithError(400, err)
    return
  }

  ctx.Status(200)
  return
}
