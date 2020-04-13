package main

import (
  "archive/zip"
  "fmt"
  "io/ioutil"
  "os"
)

func archive(dst string, files []string) {
  dstFile, err := os.Create(dst)
  if err != nil {
    panic(err)
  }
  defer dstFile.Close()

  zipWriter := zip.NewWriter(dstFile)
  for _, file := range files {
    fileWriter, err := zipWriter.Create(file)
    if err != nil {
      panic(err)
    }

    // read all from file
    bytes, err := ioutil.ReadFile(file)
    if err != nil {
      panic(err)
    }

    // write bytes to writer
    writtenBytes, err := fileWriter.Write(bytes)
    if err != nil {
      panic(err)
    }

    fmt.Println("Written bytes:", writtenBytes)
  }

  err = zipWriter.Close()
  if err != nil {
    panic(err)
  }
}

func main() {
  dst := "test.zip"
  files := []string{
    "lorem.txt",
    "lorem_2.txt",
    "lorem_3.txt",
  }

  archive(dst, files)
}
