package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
    z := uint8((x-y)/2)
    return color.RGBA{z, z, 255, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
