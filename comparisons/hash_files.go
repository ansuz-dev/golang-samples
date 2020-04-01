package main

import (
  "crypto/sha256"
  "encoding/hex"
  "fmt"
  "io"
  "os"
  "runtime"
  "sync"
  "time"
)

func fileHash(file string) (hash string, err error) {
  hasher := sha256.New()
  f, err := os.Open(file)
  if err != nil {
    return
  }
  defer f.Close()

  if _, err = io.Copy(hasher, f); err != nil {
    return
  }

  return hex.EncodeToString(hasher.Sum(nil)), nil
}

func hashNTimes(n int) {
  var wg sync.WaitGroup

  for i := 0; i < n; i++ {
    wg.Add(1)
    go func(wg *sync.WaitGroup, index int) {
      defer wg.Done()

      _, err := fileHash("./cat.jpg")
      if err != nil {
        return
      }
    }(&wg, i)
  }

  wg.Wait()
}

func main() {
  nprocs := runtime.GOMAXPROCS(0)
  fmt.Printf("n CPU = %d\n", nprocs)

  t0 := time.Now()
  hashNTimes(1000)
  elapsed := time.Since(t0)
  fmt.Printf("Execution time = %d nanoseconds\n", elapsed.Nanoseconds())
}
