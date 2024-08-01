package asciiart

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

const (
	aaSet = "MWN$@%#&B89EGA6mK5HRkbYT43V0JL7gpaseyxznocv?jIftr1li*=-~^`':;,. "
)

// EncodingType represents the encoding type of the base64 encoded image.
type EncodingType string

const (
	// StdEncoding represents the standard encoding type.
	StdEncoding EncodingType = "std"
	// URLEncoding represents the URL encoding type.
	URLEncoding EncodingType = "url"
)

// Generate generates ASCII art from an image.
func Generate(reader io.Reader) (string, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	return toASCII(img), nil
}

// GenerateFromBase64 generates ASCII art from a base64 encoded image.
func GenerateFromBase64(encodedString string, encodingType EncodingType) (string, error) {
	base64Str, err := extractBase64Data(encodedString)
	if err != nil {
		return "", err
	}

	var base64Decoder *base64.Encoding
	switch encodingType {
	case StdEncoding:
		base64Decoder = base64.StdEncoding
	case URLEncoding:
		base64Decoder = base64.URLEncoding
	default:
		return "", fmt.Errorf("unsupported encoding type: %s", encodingType)
	}

	imageBytes, err := base64Decoder.DecodeString(base64Str)
	if err != nil {
		return "", err
	}

	return Generate(bytes.NewReader(imageBytes))
}

func extractBase64Data(encodedString string) (string, error) {
	const prefix = "data:image/"
	if !strings.HasPrefix(encodedString, prefix) {
		return "", fmt.Errorf("invalid base64 encoded string")
	}

	const suffix = ";base64,"
	base64Idx := strings.Index(encodedString, suffix)
	if base64Idx == -1 {
		return "", fmt.Errorf("invalid base64 encoded string")
	}

	return encodedString[base64Idx+len(suffix):], nil
}

func toASCII(image image.Image) string {
	bounds := image.Bounds()
	var asciiImage string

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			curPixel := image.At(x, y)
			grayColor := color.GrayModel.Convert(curPixel).(color.Gray)
			idx := int(grayColor.Y) * len(aaSet) / 256
			asciiImage += string(aaSet[idx])
		}
		asciiImage += "\n"
	}
	return asciiImage
}
