package internal

import (
	"errors"
	"image"
	"image/color"
)

type PBImage interface {
	image.Image
	Set(x, y int, color color.Color)
}

type TransformerFunc func(a, b uint32) uint8

var (
	xorFunc = func(a, b uint32) uint8 {
		return uint8(a ^ b)
	}
	orFunc = func(a, b uint32) uint8 {
		return uint8(a | b)
	}
	andFunc = func(a, b uint32) uint8 {
		return uint8(a & b)
	}
)

// MultiplyImages "multiplies" two images with using operator func fn
func MultiplyImages(imageOne, imageTwo image.Image, resultImage PBImage, width, height int, fn TransformerFunc) {
	for y := height; y >= 0; y-- {
		for x := width; x >= 0; x-- {
			r1, g1, b1, a1 := imageOne.At(x, y).RGBA()
			r2, g2, b2, a2 := imageTwo.At(x, y).RGBA()
			resultImage.Set(x, y, &color.RGBA{
				R: fn(r1, r2),
				G: fn(g1, g2),
				B: fn(b1, b2),
				A: fn(a1, a2),
			})
		}
	}
}

func GetOperatorFunc(op string) (TransformerFunc, error) {
	switch op {
	default:
		return nil, errors.New("operation not supported: " + op)
	case "xor":
		return xorFunc, nil
	case "or":
		return orFunc, nil
	case "and":
		return andFunc, nil
	}
}

func GetProperImage(model color.Model, rect image.Rectangle) (PBImage, error) {
	switch model {
	default:
		return image.NewRGBA64(rect), nil
	case color.Alpha16Model:
		return image.NewAlpha16(rect), nil
	case color.AlphaModel:
		return image.NewAlpha(rect), nil
	case color.Gray16Model:
		return image.NewGray16(rect), nil
	case color.GrayModel:
		return image.NewGray(rect), nil
	case color.CMYKModel:
		return image.NewCMYK(rect), nil
	case color.NRGBA64Model:
		return image.NewNRGBA64(rect), nil
	case color.NRGBAModel:
		return image.NewNRGBA(rect), nil
	case color.RGBA64Model:
		return image.NewRGBA64(rect), nil
	case color.RGBAModel:
		return image.NewRGBA(rect), nil
	}
}
