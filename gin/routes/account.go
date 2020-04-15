package routes

import (
  "github.com/gin-gonic/gin"
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
