package main

import (
  "fmt"
  "github.com/disintegration/gift"
  "image"
  "image/jpeg"
  _ "image/png"
  "os"
  "path/filepath"
  "strings"
)

// load image file to image.Image struct
func loadImage(filename string) (img image.Image, err error) {
  f, err := os.Open(filename)
  if err != nil {
    return
  }
  defer f.Close()
  img, _, err = image.Decode(f)
  if err != nil {
    return
  }
  return
}

// save image.Image struct to jpeg file
func saveImage(filename string, img image.Image) (err error) {
  f, err := os.Create(filename)
  if err != nil {
    return
  }
  defer f.Close()
  err = jpeg.Encode(f, img, &jpeg.Options{
    Quality: 80,
  })
  if err != nil {
    return
  }

  return
}

// get file name without extension
func filenameWithoutExt(fn string) string {
  return strings.TrimSuffix(fn, filepath.Ext(fn))
}

// resize image to given width
func resize(src string, w int) (dst string, err error) {
  g := gift.New(
    gift.Resize(w, 0, gift.LanczosResampling),
  )
  imageSrc, err := loadImage(src)
  if err != nil {
    return
  }

  imageDst := image.NewRGBA(g.Bounds(imageSrc.Bounds()))
  g.Draw(imageDst, imageSrc)

  dst = fmt.Sprintf(
    "%s/%s_%dw.jpg",
    filepath.Dir(src),
    filenameWithoutExt(filepath.Base(src)),
    w,
  )

  err = saveImage(dst, imageDst)

  return
}

// get image dimension
func getDimension(imagePath string) (int, int, error) {
  file, err := os.Open(imagePath)
  if err != nil {
    return 0, 0, err
  }

  image, _, err := image.DecodeConfig(file)
  if err != nil {
    return 0, 0, err
  }
  return image.Width, image.Height, nil
}

func main() {
  imagePath := "./sample.png"

  w, h, err := getDimension(imagePath)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Width: %d, Height: %d\n", w, h)

  dst, err := resize(imagePath, 1024)
  if err != nil {
    panic(err)
  }

  fmt.Println("Dst: ", dst)

  fmt.Println("Done !")
}
