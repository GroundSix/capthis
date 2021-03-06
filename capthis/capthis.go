/**
 * capthis v0.1-dev
 *
 * (c) Ground Six
 *
 * @package capthis
 * @version 0.1-dev
 * 
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package capthis

import (
    "fmt"
    "../vendor/imagick/imagick"
)

/**
 * @var imagick.MagickWand instance of MagickWand
 * @var imagick.PixelWand instance of PixelWand
 * @var imagick.DrawingWant instance of DrawingWand
 *
 * @var string text content to overlay the image
 * @var string font for the text
 * @var string name of the font fill color
 * @var string name of the font stroke color
 * @var string name of the new image to be outputted
 */
type Caption struct {
    magickWand   *imagick.MagickWand
    pixelWand    *imagick.PixelWand
    drawingWand  *imagick.DrawingWand

    name         string
    text         string
    font         string
    fontSize     float64
    fillColor    string
    strokeColor  string
    newImageName string
}

/**
 * Creates a new instance of imagick
 * with each wand for processing the
 * text over the image
 *
 * @return Caption
 */
func New() *Caption {
    caption := new(Caption)
    caption.magickWand = imagick.NewMagickWand()
    caption.pixelWand = imagick.NewPixelWand()
    caption.drawingWand = imagick.NewDrawingWand()

    return caption
}

/**
 * Process the image by adding the text
 * overlay according to data in the
 * defined struct
 *
 * @return nil
 */
func (c *Caption) ProcessImage() {
    c.magickWand.ReadImage(c.name)
    c.magickWand.SetImageMatte(false)
    c.drawingWand.SetFont(c.font)
    c.drawingWand.SetFontSize(c.fontSize)

    // Setting the fill color for the text
    c.pixelWand.SetColor(c.fillColor)
    c.drawingWand.SetFillColor(c.pixelWand)

    // Setting the stroke color for the text
    c.pixelWand.SetColor(c.strokeColor)
    c.drawingWand.SetFillColor(c.pixelWand)

    c.drawingWand.SetGravity(imagick.GRAVITY_CENTER)
    c.magickWand.AnnotateImage(c.drawingWand, 0, 0, 0, c.text)
    c.magickWand.WriteImage(c.newImageName)

    defer c.magickWand.Destroy()
    defer c.drawingWand.Destroy()
    defer c.pixelWand.Destroy()

    fmt.Println(c.name, "has been written to", c.newImageName)
}

/**
 * Sets the name of the initial image
 * to be processed
 *
 * @param string image name
 */
func (c *Caption) SetName(name string) {
    c.name = name
}

/**
 * Sets the text that will overlay
 * the image
 *
 * @param string image text
 */
func (c *Caption) SetText(text string) {
    c.text = text
}

/**
 * Sets the font for the text
 *
 * @param string text font
 */
func (c *Caption) SetFont(font string) {
    c.font = font
}

/**
 * Sets the font size for the text
 *
 * @param int text font size
 */
func (c *Caption) SetFontSize(fontSize float64) {
    c.fontSize = fontSize
}

/**
 * Sets the fill color for the text
 *
 * @param string text fill color
 */
func (c *Caption) SetFillColor(fillColor string) {
    c.fillColor = fillColor
}

/**
 * Sets the stroke color for the text
 *
 * @param string text stroke color
 */
func (c *Caption) SetStrokeColor(strokeColor string) {
    c.strokeColor = strokeColor
}

/**
 * Sets the name of the image that
 * will be outputted at the end
 * of processing
 *
 * @param string new image name
 */
func (c *Caption) SetNewImageName(newImageName string) {
    c.newImageName = newImageName
}

/**
 * Returns our instance of MagickWand
 *
 * @return imagick.MagickWand
 */
func (c Caption) MagickWand() *imagick.MagickWand  {
    return c.magickWand
}

/**
 * Returns our instance of PixelWand
 *
 * @return imagick.PixelWand
 */
func (c Caption) PixelWand() *imagick.PixelWand {
    return c.pixelWand
}

/**
 * Returns our instance of DrawingWand
 *
 * @return imagick.DrawingWand
 */
func (c Caption) DrawingWand() *imagick.DrawingWand {
    return c.drawingWand
}
