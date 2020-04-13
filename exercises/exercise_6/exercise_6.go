package main

import (
  "compress/lzw"
  "io"
  "os"
)

/*
 * Compress `lorem.txt` using lzw algorithm
 */

func compress() (err error) {
  src, err := os.Open("lorem.txt")
  if err != nil {
    return
  }
  defer src.Close()

  dst, err := os.Create("lorem.txt.lzw")
  if err != nil {
    return
  }
  defer dst.Close()

  lzwWriter := lzw.NewWriter(dst, lzw.LSB, 8)
  defer lzwWriter.Close()

  _, err = io.Copy(lzwWriter, src)

  return
}

func decompress() (err error) {
  src, err := os.Open("lorem.txt.lzw")
  if err != nil {
    return
  }
  defer src.Close()

  dst, err := os.Create("lorem_de.txt")
  if err != nil {
    return
  }
  defer dst.Close()

  lzwReader := lzw.NewReader(src, lzw.LSB, 8)
  defer lzwReader.Close()

  _, err = io.Copy(dst, lzwReader)

  return
}

func main() {
  compress()
  decompress()
}
