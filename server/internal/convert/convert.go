package convert

import (
	"reflect"

	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
)

type Converter struct {
	tcoder *transcoder.Transcoder
	vopt   VideoOptionsMap
}

func NewConverter() *Converter {
	return &Converter{
		new(transcoder.Transcoder),
		make(VideoOptionsMap, 53),
	}
}

func (conv *Converter) Initialize() {
	conv.vopt.AddFunc("VideoBitRate", (*models.Mediafile).SetVideoBitRate)
	conv.vopt.AddFunc("AudioBitRate", (*models.Mediafile).SetAudioBitRate)
	conv.vopt.AddFunc("AudioCodec", (*models.Mediafile).SetAudioCodec)
	conv.vopt.AddFunc("VideoCodec", (*models.Mediafile).SetVideoCodec)

	conv.vopt.AddFunc("Channels", (*models.Mediafile).SetAudioChannels)
	conv.vopt.AddFunc("Resolution", (*models.Mediafile).SetResolution)
	conv.vopt.AddFunc("AspectRatio", (*models.Mediafile).SetAspect)

	conv.vopt.AddFunc("FrameRate", (*models.Mediafile).SetFrameRate)

	// TODO: volume, filters, trim and etc
}

func (conv *Converter) ConvertVideo(inputPath, outputPath string, options ConversionVideoOptions) error {
	if err := conv.tcoder.Initialize(inputPath, outputPath); err != nil {
		return err
	}

	o := reflect.ValueOf(options.Output)
	for i := 0; i < o.NumField(); i++ {
		fname := o.Type().Field(i).Name
		fvalue := o.Field(i).Interface()

		conv.vopt.CallFunc(fname, conv.tcoder.MediaFile(), fvalue)
	}

	done := conv.tcoder.Run(false)
	return <-done
}
