
package main

import (
	"os"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"net/http"
	"path/filepath"
	"image"
	_ "errors"
)


func getType(file  *os.File) (string) {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return "ciao"
	}
	file.Seek(0, 0)
	return http.DetectContentType(buff)
}



func Optimize(path string, width uint, height uint) (string, error) {

	var file *os.File
	var newFileImg *os.File

	file, _ = os.Open(path)

	defer file.Close()

	var imageType string = getType(file)
	var img image.Image = nil

	if imageType == "image/jpeg" {
		img, _ = jpeg.Decode(file)
	}
	if imageType == "image/png" {
		img, _ = png.Decode(file)
	}

	newImage := resize.Resize(width, height, img, resize.NearestNeighbor)

	var destPath string = filepath.Join("optimized", filepath.Base(path))

	newFileImg, _ = os.Create(destPath)
	defer newFileImg.Close()

	if imageType == "image/jpeg" {
		jpeg.Encode(newFileImg, newImage, nil)
	}
	if imageType == "image/png" {
		png.Encode(newFileImg, newImage)
	}

	return destPath, nil
}
