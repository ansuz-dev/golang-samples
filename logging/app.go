package main

import (
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "github.com/sirupsen/logrus/hooks/writer"
  "io/ioutil"
  "os"
  "time"
)

var consoleLog *logrus.Logger
var consoleLog2 *logrus.Logger
var fileLog *logrus.Logger
var hookLog *logrus.Logger

func addLogger(logger *logrus.Logger) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    startTime := time.Now()

    // Process request
    ctx.Next()

    endTime := time.Now()
    latencyTime := endTime.Sub(startTime)

    logger.Infof(
      "| %3d | %5v | %10s | %s | %s |",
      ctx.Writer.Status(),
      latencyTime,
      ctx.ClientIP(),
      ctx.Request.Method,
      ctx.Request.RequestURI,
    )
  }
}

func logConsole2() gin.HandlerFunc {
  consoleLog2 := logrus.New()

  // Write log to console
  consoleLog2.Out = os.Stdout

  // Log levels:
  // - Trace
  // - Debug
  // - Info
  // - Warn
  // - Error
  // - Fatal
  // - Panic
  consoleLog2.SetLevel(logrus.DebugLevel)

  return func(ctx *gin.Context) {
    startTime := time.Now()

    // Process request
    ctx.Next()

    endTime := time.Now()
    latencyTime := endTime.Sub(startTime)

    consoleLog2.WithFields(logrus.Fields{
      "status_code":  ctx.Writer.Status(),
      "latency_time": latencyTime,
      "client_ip":    ctx.ClientIP(),
      "method":       ctx.Request.Method,
      "req_uri":      ctx.Request.RequestURI,
    }).Info()
  }
}

func main() {
  g := gin.New()

  consoleLog := logrus.New()
  consoleLog.SetOutput(os.Stdout)

  logfile, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }
  defer logfile.Close()

  fileLog := logrus.New()
  fileLog.SetOutput(logfile)

  errorfile, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }
  defer errorfile.Close()

  debugfile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }
  defer debugfile.Close()

  hookLog := logrus.New()
  hookLog.SetOutput(ioutil.Discard)

  hookLog.AddHook(&writer.Hook{
    Writer: errorfile,
    LogLevels: []logrus.Level{
      logrus.PanicLevel,
      logrus.FatalLevel,
      logrus.ErrorLevel,
      logrus.WarnLevel,
    },
  })

  hookLog.AddHook(&writer.Hook{
    Writer: debugfile,
    LogLevels: []logrus.Level{
      logrus.InfoLevel,
      logrus.DebugLevel,
      logrus.TraceLevel,
    },
  })

  hookLog.Error("This error should be logged in file error.log")
  hookLog.Info("This info should be logged in file debug.log")

  g.Use(addLogger(consoleLog))
  g.Use(logConsole2())
  g.Use(addLogger(fileLog))

  g.Run("127.0.0.1:3000")

}
