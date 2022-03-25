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
		make(VideoOptionsMap, 28),
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

func (conv *Converter) ConvertVideo(inputPath, outputPath string, options ConversionVideoOptions) error {
	if err := conv.tcoder.Initialize(inputPath, outputPath); err != nil {
		return err
	}

	o := reflect.ValueOf(options.Output)
	for i := 0; i < o.NumField(); i++ {
		fname := o.Type().Field(i).Name
		fvalue := o.Field(i).Interface()

		// TODO: Actually we process all fields including default values.
		// If for a string field it's pretty ez to check the value. For int, uint and bool it's not.
		// Find a way to do that.
		conv.vopt.CallFunc(fname, conv.tcoder.MediaFile(), fvalue)
	}

	done := conv.tcoder.Run(false)
	return <-done
}
