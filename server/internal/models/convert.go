package models

import "hara/internal/convert"

type Converter interface {
	ConvertImage(string, convert.ConversionOptions) (string, error)
	ConvertVideo(string, convert.ConversionOptions) (string, error)
}
