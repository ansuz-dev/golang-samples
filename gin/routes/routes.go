package routes

import (
  "fmt"
  "github.com/gin-gonic/gin"
  // "net/http"
  // "path/filepath"
)

func middleware1() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    // Get "m_2" from middleware2
    if m2, exists := ctx.Get("m_2"); exists {
      fmt.Println("In middleware 1", m2)
    } else {
      fmt.Println("middleware 2 is not executed")
    }

    ctx.Set("m_1", "passed m1")
    ctx.Next()
  }
}

func middleware2() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    // Get "m_1" from middleware1
    if m1, exists := ctx.Get("m_1"); exists {
      fmt.Println("In middleware 2", m1)
    }

    ctx.Set("m_2", "passed m2")
    ctx.Next()
  }
}

func middleware3() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    ctx.Next()

    // Get "m_1" from middleware1
    if m1, exists := ctx.Get("m_1"); exists {
      fmt.Println("In middleware 3", m1)
    }
    // Get "m_2" from middleware2
    if m2, exists := ctx.Get("m_2"); exists {
      fmt.Println("In middleware 3", m2)
    }

    ctx.Set("m_3", "passed m3")
  }
}

func Create() (g *gin.Engine) {

  g = gin.Default()

  v1 := g.Group("/v1")
  {
    // Use middlewares in "/v1" API group
    v1.Use(middleware1())
    v1.Use(middleware2())
    v1.Use(middleware3())

    v1.GET("", func(ctx *gin.Context) {
      if m1, exists := ctx.Get("m_1"); exists {
        fmt.Println(m1)
      }
      if m2, exists := ctx.Get("m_2"); exists {
        fmt.Println(m2)
      }
      if m3, exists := ctx.Get("m_3"); exists {
        fmt.Println(m3)
      }
      ctx.String(200, "API v1.0")
    })

    // cookie := v1.Group("/cookie")
    // {
    //   cookie.GET("/set", func(ctx *gin.Context) {
    //     ctx.SetCookie("token", "my_token", 0, "/", "", false, true)
    //     ctx.Status(200)
    //   })

    //   cookie.GET("/get", func(ctx *gin.Context) {
    //     token, err := ctx.Cookie("token")
    //     if err != nil {
    //       if err == http.ErrNoCookie {
    //         ctx.String(200, "Token not found")
    //       } else {
    //         ctx.AbortWithError(500, err)
    //       }
    //       return
    //     }
    //     ctx.String(200, token)
    //   })

    //   cookie.GET("/remove", func(ctx *gin.Context) {
    //     ctx.SetCookie("token", "", -1, "/", "", false, true)
    //     ctx.Status(200)
    //   })
    // }

    // upload := v1.Group("/upload")
    // {
    //   // curl -F 'file=@/c/Users/aoshi/Documents/hoa_don_007.pdf' http://127.0.0.1:3000/v1/upload/single
    //   upload.POST("/single", func(ctx *gin.Context) {
    //     file, err := ctx.FormFile("file")
    //     if err != nil {
    //       ctx.AbortWithError(400, err)
    //       return
    //     }

    //     filename := filepath.Base(file.Filename)
    //     if err := ctx.SaveUploadedFile(file, filename); err != nil {
    //       ctx.AbortWithError(400, err)
    //       return
    //     }

    //     ctx.Status(200)
    //   })

    //   // curl -F 'files=@/c/Users/aoshi/Documents/hoa_don_007.pdf' -F 'files=@/c/Users/aoshi/Documents/hoa_don_008.pdf' http://127.0.0.1:3000/v1/upload/multiple
    //   upload.POST("/multiple", func(ctx *gin.Context) {
    //     // Multipart form
    //     form, err := ctx.MultipartForm()
    //     if err != nil {
    //       ctx.AbortWithError(400, err)
    //       return
    //     }
    //     files := form.File["files"]

    //     for _, file := range files {
    //       filename := filepath.Base(file.Filename)
    //       if err := ctx.SaveUploadedFile(file, filename); err != nil {
    //         ctx.AbortWithError(400, err)
    //         return
    //       }
    //     }

    //     ctx.Status(200)
    //   })
    // }

    account := v1.Group("/account")
    {
      // add account handlers
      account.GET("/:id", GetAccountByID)
      account.PUT("/:id", UpdateAccountByID)
    }

    post := v1.Group("/post")
    {
      // add post handlers
      post.POST("", CreatePost)
      post.GET("", GetPostByAccount)
      post.PUT("/:id", UpdatePostByID)
      post.DELETE("/:id", nil)
    }
  }

  return
}
