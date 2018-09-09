
package main

import (
	"os"
	"io"
	"fmt"
	"os/exec"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"path/filepath"
	"errors"
)


type ImageHandler interface {
    ImageType() string
    Decode(reader io.Reader) (image.Image, error)
    Encode(newImgFile *os.File, newImage image.Image) error
    Convert(newImageTempPath string, quality uint) error
}

func NewImageHandler(file  *os.File) (ImageHandler, error) {
	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		return nil, errors.New("Can't get type of this file")
	}
	defer file.Seek(0, 0)
    switch http.DetectContentType(buff) {
		case "image/jpeg":
			return &JPEGHandler{}, nil
		case "image/png":
			return &PNGHandler{}, nil
		default:
			return nil, errors.New("Unsupported Image type")
    }
}

type JPEGHandler struct {
}

func (j *JPEGHandler) ImageType() string {
    return "image/png"
}

func (j *JPEGHandler) Decode(reader io.Reader) (image.Image, error) {
    return jpeg.Decode(reader)
}

func (j *JPEGHandler) Encode(newImgFile *os.File, newImage image.Image) error {
    return jpeg.Encode(newImgFile, newImage, nil)
}

func (j *JPEGHandler) Convert(newImageTempPath string, quality uint) error {
    args := []string{fmt.Sprintf("--max=%d", quality), newImageTempPath}
    cmd := exec.Command("jpegoptim", args...)
    err := cmd.Run()
    if err != nil {
        return errors.New("Jpegoptim command not working")
    }

    return nil
}

type PNGHandler struct {
}

func (p *PNGHandler) ImageType() string {
    return "image/png"
}

func (p *PNGHandler) Decode(reader io.Reader) (image.Image, error) {
    return png.Decode(reader)
}

func (p *PNGHandler) Encode(newImgFile *os.File, newImage image.Image) error {
    return png.Encode(newImgFile, newImage)
}

func (p *PNGHandler) Convert(newImageTempPath string, quality uint) error {
    args := []string{newImageTempPath, "-f", "--ext=.png"}

    if quality != 100 {
        var qualityMin = quality - 10
        qualityParameter := fmt.Sprintf("--quality=%[1]d-%[2]d", qualityMin, quality)
        args = append([]string{qualityParameter}, args...)
    }
    cmd := exec.Command("pngquant", args...)
    err := cmd.Run()
    if err != nil {
        return errors.New("Pngquant command not working")
    }
    return nil
}

func Optimize(path string, width uint, height uint) (string, error) {

	var file *os.File
	var newFileImg *os.File
	var img image.Image = nil
	var err error = nil
	var imageHandler ImageHandler
	var destPath string = ""

	file, _ = os.Open(path)

	defer file.Close()

	imageHandler, err = NewImageHandler(file)

	if err != nil {
		return "", err
	}

	img, err = imageHandler.Decode(file)

	if err != nil {
		return "", err
	}

	newImage := resize.Resize(width, height, img, resize.NearestNeighbor)

	destPath = filepath.Join("optimized", filepath.Base(path))

	newFileImg, err = os.Create(destPath)

	if err != nil {
		return "", err
	}

	defer newFileImg.Close()

	err = imageHandler.Encode(newFileImg, newImage)

	imageHandler.Convert(destPath, 100)

	if err != nil {
        return "", err
    }

	return destPath, nil
}
