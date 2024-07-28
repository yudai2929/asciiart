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
	"os"
)

const (
	aaSet = "MWN$@%#&B89EGA6mK5HRkbYT43V0JL7gpaseyxznocv?jIftr1li*=-~^`':;,. "
)

type mimeType string

const (
	mimeTypePNG  mimeType = "png"
	mimeTypeJPEG mimeType = "jpeg"
)

// EncodingType represents the encoding type of the base64 encoded image.
type EncodingType string

const (
	// StdEncoding represents the standard encoding type.
	StdEncoding EncodingType = "std"
	// URLEncoding represents the URL encoding type.
	URLEncoding EncodingType = "url"
)

// Generate generates ASCII art from an image file.
func Generate(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	return toASCII(img), nil
}

// GenerateFromBase64 generates ASCII art from a base64 encoded image.
func GenerateFromBase64(encodedString string, encodingType EncodingType) (string, error) {
	base64Str, imageType, err := extractBase64Data(encodedString)
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

	img, err := toImage(bytes.NewReader(imageBytes), imageType)
	if err != nil {
		return "", err
	}

	return toASCII(img), nil
}

func extractBase64Data(encodedString string) (string, mimeType, error) {
	const prefix = "data:image/"
	const suffix = ";base64,"

	if len(encodedString) < len(prefix)+len(suffix) {
		return "", "", fmt.Errorf("invalid base64 encoded string")
	}

	mimeStr := encodedString[len(prefix) : len(prefix)+3]
	if !isValidMimeType(mimeStr) {
		return "", "", fmt.Errorf("unsupported image type: %s", mimeStr)
	}
	base64Str := encodedString[len(prefix)+3+len(suffix):]
	return base64Str, mimeType(mimeStr), nil
}

func isValidMimeType(mimeStr string) bool {
	switch mimeStr {
	case string(mimeTypePNG), string(mimeTypeJPEG):
		return true
	default:
		return false
	}
}

func toImage(reader io.Reader, imageType mimeType) (img image.Image, err error) {
	switch imageType {
	case mimeTypePNG:
		img, err = png.Decode(reader)
	case mimeTypeJPEG:
		img, err = jpeg.Decode(reader)
	default:
		return nil, fmt.Errorf("unsupported image type: %s", imageType)
	}
	return img, err
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
