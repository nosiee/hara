package convert

import (
	"fmt"
	"reflect"

	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
)

type Converter struct {
	tcoder     *transcoder.Transcoder
	vopt       VideoOptionsMap
	outputPath string
}

func NewConverter(outputPath string) *Converter {
	return &Converter{
		new(transcoder.Transcoder),
		make(VideoOptionsMap, 28),
		outputPath,
	}
}

func (conv *Converter) Initialize() {
	conv.vopt.AddFunc("AspectRatio", (*models.Mediafile).SetAspect)
	conv.vopt.AddFunc("Resolution", (*models.Mediafile).SetResolution)
	conv.vopt.AddFunc("VideoBitRate", (*models.Mediafile).SetVideoBitRate)
	conv.vopt.AddFunc("VideoMaxBitRate", (*models.Mediafile).SetVideoMaxBitrate)
	conv.vopt.AddFunc("VideoMinBitRate", (*models.Mediafile).SetVideoMinBitRate)
	conv.vopt.AddFunc("VideoCodec", (*models.Mediafile).SetVideoCodec)
	conv.vopt.AddFunc("VFrames", (*models.Mediafile).SetVframes)
	conv.vopt.AddFunc("FrameRate", (*models.Mediafile).SetFrameRate)
	conv.vopt.AddFunc("AudioRate", (*models.Mediafile).SetAudioRate)
	conv.vopt.AddFunc("SkipVideo", (*models.Mediafile).SetSkipVideo)
	conv.vopt.AddFunc("SkipAudio", (*models.Mediafile).SetSkipAudio)
	conv.vopt.AddFunc("MaxKeyFrame", (*models.Mediafile).SetMaxKeyFrame)
	conv.vopt.AddFunc("MinKeyFrame", (*models.Mediafile).SetMinKeyFrame)
	conv.vopt.AddFunc("KeyframeInterval", (*models.Mediafile).SetKeyframeInterval)
	conv.vopt.AddFunc("AudioCodec", (*models.Mediafile).SetAudioCodec)
	conv.vopt.AddFunc("AudioBitRate", (*models.Mediafile).SetAudioBitRate)
	conv.vopt.AddFunc("Channels", (*models.Mediafile).SetAudioChannels)
	conv.vopt.AddFunc("BufferSize", (*models.Mediafile).SetBufferSize)
	conv.vopt.AddFunc("Preset", (*models.Mediafile).SetPreset)
	conv.vopt.AddFunc("Tune", (*models.Mediafile).SetTune)
	conv.vopt.AddFunc("AudioProfile", (*models.Mediafile).SetAudioProfile)
	conv.vopt.AddFunc("VideoProfile", (*models.Mediafile).SetVideoProfile)
	conv.vopt.AddFunc("Duration", (*models.Mediafile).SetDuration)
	conv.vopt.AddFunc("SeekTime", (*models.Mediafile).SetSeekTime)
	conv.vopt.AddFunc("Strict", (*models.Mediafile).SetStrict)
	conv.vopt.AddFunc("AudioFilter", (*models.Mediafile).SetAudioFilter)
	conv.vopt.AddFunc("VideoFilter", (*models.Mediafile).SetVideoFilter)
	conv.vopt.AddFunc("CompressionLevel", (*models.Mediafile).SetCompressionLevel)
}

func (conv *Converter) ConvertVideo(inputPath string, options ConversionVideoOptions) error {
	outputPath := fmt.Sprintf("%s/%s", conv.outputPath, options.Output.Name)
	if err := conv.tcoder.Initialize(inputPath, outputPath); err != nil {
		return err
	}

	o := reflect.ValueOf(options.Output)
	for i := 0; i < o.NumField(); i++ {
		fname := o.Type().Field(i).Name
		fvalue := o.Field(i).Interface()

		// NOTE: (vc_default_fiels): Default values for integer and boolean fields have no effect on conversion
		// So we can pass them to the as arguments.
		conv.vopt.CallFunc(fname, conv.tcoder.MediaFile(), fvalue)
	}

	done := conv.tcoder.Run(false)
	return <-done
}
