package capthis

import (
	"../vendor/imagick/imagick"
)

type Caption struct {
	magicWand 	 *imagick.MagickWand
	pixelWand 	 *imagick.PixelWand
	drawingWand  *imagick.DrawingWand

	text 		 string
	font 		 string
	fillColor 	 string
	strokeColor  string
	newImageName string
}

func New() *Caption {
	caption := new(Caption)
	caption.magicWand = imagick.NewMagickWand()
	caption.pixelWand = imagick.NewPixelWand()
	caption.drawingWand = imagick.NewDrawingWand()

	return caption
}

func (c *Caption) SetText(text string) {
	c.text = text
}

func (c *Caption) SetFont(font string) {
	c.font = font
}

func (c *Caption) SetFillColor(fillColor string) {
	c.fillColor = fillColor
}

func (c *Caption) SetStrokeColor(strokeColor string) {
	c.strokeColor = strokeColor
}

func (c *Caption) SetNewImageName(newImageName string) {
	c.newImageName = newImageName
}
