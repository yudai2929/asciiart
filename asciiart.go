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

	"golang.org/x/image/draw"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

const (
	defaultAASet = "MWN$@%#&B89EGA6mK5HRkbYT43V0JL7gpaseyxznocv?jIftr1li*=-~^`':;,. "
	defaultWidth = 100
)

// EncodingType represents the encoding type of the base64 encoded image.
type EncodingType string

const (
	// StdEncoding represents the standard encoding type.
	StdEncoding EncodingType = "std"
	// URLEncoding represents the URL encoding type.
	URLEncoding EncodingType = "url"
)

type options struct {
	width int
	aaSet string
}

// Option represents the option to customize the ASCII art.
type Option func(*options)

// WithWidth sets the width of the ASCII art.
func WithWidth(width int) Option {
	return func(o *options) {
		o.width = width
	}
}

// WithAASet sets the ASCII art set.
func WithAASet(aaSet string) Option {
	return func(o *options) {
		o.aaSet = aaSet
	}
}

// Generate generates ASCII art from an image.
func Generate(reader io.Reader, opts ...Option) (string, error) {
	o := options{
		width: defaultWidth,
		aaSet: defaultAASet,
	}

	for _, opt := range opts {
		opt(&o)
	}

	if o.width < 50 || o.width > 200 {
		return "", fmt.Errorf("width must be between 50 and 200")
	}

	if len(o.aaSet) < 4 || len(o.aaSet) > 64 {
		return "", fmt.Errorf("ASCII art set must be between 4 and 64 characters")
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	// Calculate the height based on the newWidth and original aspect ratio
	newWidth, newHeight := calculateDimensions(img, o.width)
	resizedImg := resizeImage(img, newWidth, newHeight)

	return toASCII(resizedImg, o.aaSet), nil
}

// GenerateFromBase64 generates ASCII art from a base64 encoded image.
func GenerateFromBase64(encodedString string, encodingType EncodingType, opts ...Option) (string, error) {
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

	return Generate(bytes.NewReader(imageBytes), opts...)
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

func resizeImage(img image.Image, newWidth, newHeight int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return dst
}

func calculateDimensions(img image.Image, newWidth int) (int, int) {
	bounds := img.Bounds()
	originalWidth := bounds.Dx()
	originalHeight := bounds.Dy()

	if newWidth == 0 {
		return originalWidth, originalHeight
	}

	newHeight := int(float64(newWidth) * float64(originalHeight) / float64(originalWidth))

	return newWidth, newHeight
}

func toASCII(image image.Image, aaSet string) string {
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
