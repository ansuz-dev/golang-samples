package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

// curl -F 'file=@/c/Users/aoshi/Documents/hoa_don_007.pdf' http://127.0.0.1:3000/upload
func upload(w http.ResponseWriter, r *http.Request) {
  fmt.Println("File Upload Endpoint Hit")

  r.ParseMultipartForm(10 << 20)
  file, handler, err := r.FormFile("file")
  if err != nil {
    fmt.Println("Error Retrieving the File")
    fmt.Println(err)
    return
  }
  defer file.Close()
  fmt.Printf("Uploaded File: %+v\n", handler.Filename)
  fmt.Printf("File Size: %+v\n", handler.Size)
  fmt.Printf("MIME Header: %+v\n", handler.Header)

  dstFile, err := os.Create(handler.Filename)
  if err != nil {
    fmt.Println(err)
  }
  defer dstFile.Close()

  // read all of the contents of the uploaded file into a byte array
  fileBytes, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
  }
  // write this byte array to the destination file
  dstFile.Write(fileBytes)
  // return that we have successfully uploaded the file!
  fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:3000",
  }

  http.HandleFunc("/upload", upload)

  server.ListenAndServe()
}
