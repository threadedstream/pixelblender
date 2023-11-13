package simple

import (
	"errors"
	"image"
	"image/jpeg"
	"os"

	"github.com/threadedstream/pixelblender/internal"
)

func Main(args ...string) error {
	if len(args) < 3 {
		return errors.New("./pixelblender image_1.jpg image_2.jpg op")
	}
	imageOneStream, err := os.OpenFile(args[0], os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	imageTwoStream, err := os.OpenFile(args[1], os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	imageOne, err := jpeg.Decode(imageOneStream)
	if err != nil {
		return err
	}

	imageTwo, err := jpeg.Decode(imageTwoStream)
	if err != nil {
		return err
	}

	if imageOne.Bounds().Size() != imageTwo.Bounds().Size() {
		// TODO(gildarov): automatically resize an image upon conflict
		return errors.New("image sizes must be equal")
	}

	height := imageOne.Bounds().Dy()
	width := imageOne.Bounds().Dx()

	resultImage, err := internal.GetProperImage(imageOne.ColorModel(), image.Rect(0, 0, width, height))
	if err != nil {
		return err
	}

	opFunc, err := internal.GetOperatorFunc(args[2])
	if err != nil {
		return err
	}

	internal.MultiplyImages(imageOne, imageTwo, resultImage, width, height, opFunc)

	output, err := os.OpenFile("result.jpeg", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	if err = jpeg.Encode(output, resultImage, &jpeg.Options{
		Quality: 100,
	}); err != nil {
		return err
	}
	return nil
}
