package capthis

import (
    "fmt"
    "testing"
    "./capthis"
)

/**
 * capthis.go
 */
func TestNew(t *testing.T) {
    caption := capthis.New()
    if (fmt.Sprintf("%T", caption) != "*capthis.Caption") {
        t.Error("Expected type *capthis.Caption")
    }

    if (fmt.Sprintf("%T", caption.MagickWand()) != "*imagick.MagickWand") {
        t.Error("Expected type *imagick.MagickWand")
    }

    if (fmt.Sprintf("%T", caption.PixelWand()) != "*imagick.PixelWand") {
        t.Error("Expected type *imagick.PixelWand")
    }

    if (fmt.Sprintf("%T", caption.DrawingWand()) != "*imagick.DrawingWand") {
        t.Error("Expected type *imagick.DrawingWand")
    }
}
