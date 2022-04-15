package convert

import (
	"errors"
	"hara/internal/convert"
	"strings"
)

type MockConverter struct{}

func NewMockConverter() *MockConverter {
	return &MockConverter{}
}

func (c MockConverter) ConvertVideo(file string, options convert.ConversionOptions) (string, error) {
	switch strings.Split(file, "/")[2] {
	case "errorconvert.mp4":
		return "", errors.New("Test error")
	case "erroradd.mp4":
		return "erroradd.mp4", nil
	default:
		return "output.mp4", nil
	}
}

func (c MockConverter) ConvertImage(file string, options convert.ConversionOptions) (string, error) {
	switch strings.Split(file, "/")[2] {
	case "errorconvert.jpg":
		return "", errors.New("Test error")
	case "erroradd.jpg":
		return "erroradd.jpg", nil
	default:
		return "output.jpg", nil
	}
}
