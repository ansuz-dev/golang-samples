package routes

import (
  "github.com/gin-gonic/gin"
  "golang-samples/gin/models"
  "golang-samples/gin/services"
)

// curl -X POST -H "Content-Type: application/json" -d '{"accountId":1,"content":"This is a post"}' http://127.0.0.1:3000/v1/post

func CreatePost(ctx *gin.Context) {
  post := models.Post{}
  if err := ctx.BindJSON(&post); err != nil {
    ctx.AbortWithError(400, err)
  }

  // TODO: Create a new post

  ctx.JSON(200, post)
}

func GetPostByAccount(ctx *gin.Context) {
  // get account ID in query string
  accountId := ctx.Query("accountId")

  posts, err := services.GetPostsByAccount(accountId)
  if err != nil {
    ctx.AbortWithError(400, err)
    return
  }
  ctx.JSON(200, posts)
}

func HandlerFunc(ctx *gin.Context) {
  // TODO: handle here
}
