package search

import (
	"image"

	"encoding/hex"
	
	"crypto/md5"
)

func hashIcon(icon image.Image) string {
	hash := md5.New()

	for x := 0; x < icon.Bounds().Dx(); x++ {
		for y := 0; y < icon.Bounds().Dy(); y++ {
			r, g, b, a := icon.At(x, y).RGBA()
			
			hash.Write([]byte{byte(r), byte(g), byte(b), byte(a)})
		}
	}

	return hex.EncodeToString(hash.Sum(nil))
}
