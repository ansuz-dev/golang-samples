package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "os"
)

func main() {
  // createEmptyFile("empty.txt")
  // getFileInfo("lorem.txt")
  // rename("lorem.txt", "lorem_new.txt")
  // deleteFile("empty.txt")
  // openFile("lorem.txt")
  // existed := exists("empty.txt")
  // fmt.Println(existed)

  // createHardLink("hl_lorem.txt", "hl_lorem_2.txt")
  // getSymlinkInfo("hl_lorem.txt")

  // This function only work on Unix system
  // createSymLink("lorem.txt", "sl_lorem.txt")

  // copyFile("lorem.txt", "lorem_dst.txt")

  buff := readFile2("lorem.txt", 6000)
  fmt.Println("Length of buff:", len(buff))

  // writeFile2("empty.txt")
}

func createEmptyFile(fileName string) {
  file, err := os.Create(fileName)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  fmt.Println(file)
}

func getFileInfo(fileName string) {
  fileInfo, err := os.Stat(fileName)
  if err != nil {
    panic(err)
  }

  fmt.Println("File name: ", fileInfo.Name())
  fmt.Println("File size: ", fileInfo.Size())
  fmt.Println("Permissions: ", fileInfo.Mode())
  fmt.Println("Last modified: ", fileInfo.ModTime())
}

func rename(oldName string, newName string) {
  err := os.Rename(oldName, newName)
  if err != nil {
    panic(err)
  }
}

func deleteFile(fileName string) {
  err := os.Remove(fileName)
  if err != nil {
    panic(err)
  }
}

func openFile(fileName string) {
  // Open file to read only
  // file, err := os.Open(fileName)
  // if err != nil {
  //   panic(err)
  // }
  // file.Close()

  // Open file with more options
  file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
  if err != nil {
    panic(err)
  }
  file.Close()
}

// Check file is existing or not
func exists(fileName string) bool {
  _, err := os.Stat(fileName)
  if err != nil {
    if os.IsNotExist(err) {
      return false
    }

    panic(err)
  }

  return true
}

func createHardLink(fileName string, linkName string) {
  err := os.Link(fileName, linkName)
  if err != nil {
    panic(err)
  }
}

func createSymLink(fileName string, linkName string) {
  err := os.Symlink(fileName, linkName)
  if err != nil {
    panic(err)
  }
}

func getSymlinkInfo(linkName string) {
  fileInfo, err := os.Lstat(linkName)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%#v\n", fileInfo)
}

func copyFile(src string, dst string) {
  srcFile, err := os.Open(src)
  if err != nil {
    panic(err)
  }
  defer srcFile.Close()

  dstFile, err := os.Create(dst)
  if err != nil {
    panic(err)
  }
  defer dstFile.Close()

  bytes, err := io.Copy(dstFile, srcFile)
  if err != nil {
    panic(err)
  }

  fmt.Printf("Copied %d bytes\n", bytes)

  // Flush memory to disk
  err = dstFile.Sync()
  if err != nil {
    panic(err)
  }
}

func readFile(src string, bytes int64) []byte {
  buff := make([]byte, bytes)

  file, err := os.Open(src)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  bytesRead, err := file.Read(buff)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes read: ", bytesRead)

  return buff
}

func readFile2(src string, bytes int64) []byte {
  buff := make([]byte, bytes)

  file, err := os.Open(src)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  bytesRead, err := io.ReadFull(file, buff)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes read: ", bytesRead)

  return buff
}

func writeFile(src string) {
  file, err := os.OpenFile(src, os.O_RDWR, 0666)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  buff := []byte{65, 67, 68}
  writtenBytes, err := file.Write(buff)
  if err != nil {
    panic(err)
  }
  fmt.Println("Byte written:", writtenBytes)
}

func writeFile2(src string) {
  buff := []byte{69, 70, 71}

  err := ioutil.WriteFile(src, buff, 0666)
  if err != nil {
    panic(err)
  }
}
