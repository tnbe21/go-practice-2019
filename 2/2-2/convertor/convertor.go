package convertor

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

type convertor interface {
	_convert(src string) error
}

type jpgConvertor struct{}

type gifConvertor struct{}

func Convert(src, dst string) error {
	convertor, err := createConvertor(dst)
	if err != nil {
		return err
	}
	return convertor._convert(src)
}

func createConvertor(dst string) (convertor, error) {
	switch dst {
	case "jpg":
		return &jpgConvertor{}, nil
	case "gif":
		return &gifConvertor{}, nil
	default:
		return nil, fmt.Errorf("dst must be jpg or gif.")
	}
}

func (*jpgConvertor) _convert(src string) error {
	srcImg, err := getDecodedImg(src)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(strings.TrimRight(src, ".png") + ".jpg")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	jpeg.Encode(dstFile, srcImg, nil)
	return nil
}

func (*gifConvertor) _convert(src string) error {
	srcImg, err := getDecodedImg(src)
	if err != nil {
		return err
	}

	dstFile, err := os.Create(strings.TrimRight(src, ".png") + ".gif")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	gif.Encode(dstFile, srcImg, nil)
	return nil
}

func getDecodedImg(src string) (image.Image, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()

	srcImg, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, err
	}

	return srcImg, nil
}
