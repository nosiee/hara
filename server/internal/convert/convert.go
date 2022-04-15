package convert

import (
	"fmt"
	"hara/internal/config"
	"reflect"

	"github.com/google/uuid"
	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
	"gopkg.in/gographics/imagick.v3/imagick"
)

type Converter struct {
	tscoder *transcoder.Transcoder
	vopt    VideoOptions
}

func NewConverter() *Converter {
	return &Converter{
		new(transcoder.Transcoder),
		make(VideoOptions, 28),
	}
}

func (c *Converter) Initialize() {
	c.vopt.AddFunc("AspectRatio", (*models.Mediafile).SetAspect)
	c.vopt.AddFunc("Resolution", (*models.Mediafile).SetResolution)
	c.vopt.AddFunc("VideoBitRate", (*models.Mediafile).SetVideoBitRate)
	c.vopt.AddFunc("VideoMaxBitRate", (*models.Mediafile).SetVideoMaxBitrate)
	c.vopt.AddFunc("VideoMinBitRate", (*models.Mediafile).SetVideoMinBitRate)
	c.vopt.AddFunc("VideoCodec", (*models.Mediafile).SetVideoCodec)
	c.vopt.AddFunc("VFrames", (*models.Mediafile).SetVframes)
	c.vopt.AddFunc("FrameRate", (*models.Mediafile).SetFrameRate)
	c.vopt.AddFunc("AudioRate", (*models.Mediafile).SetAudioRate)
	c.vopt.AddFunc("SkipVideo", (*models.Mediafile).SetSkipVideo)
	c.vopt.AddFunc("SkipAudio", (*models.Mediafile).SetSkipAudio)
	c.vopt.AddFunc("MaxKeyFrame", (*models.Mediafile).SetMaxKeyFrame)
	c.vopt.AddFunc("MinKeyFrame", (*models.Mediafile).SetMinKeyFrame)
	c.vopt.AddFunc("KeyframeInterval", (*models.Mediafile).SetKeyframeInterval)
	c.vopt.AddFunc("AudioCodec", (*models.Mediafile).SetAudioCodec)
	c.vopt.AddFunc("AudioBitRate", (*models.Mediafile).SetAudioBitRate)
	c.vopt.AddFunc("Channels", (*models.Mediafile).SetAudioChannels)
	c.vopt.AddFunc("BufferSize", (*models.Mediafile).SetBufferSize)
	c.vopt.AddFunc("Preset", (*models.Mediafile).SetPreset)
	c.vopt.AddFunc("Tune", (*models.Mediafile).SetTune)
	c.vopt.AddFunc("AudioProfile", (*models.Mediafile).SetAudioProfile)
	c.vopt.AddFunc("VideoProfile", (*models.Mediafile).SetVideoProfile)
	c.vopt.AddFunc("Duration", (*models.Mediafile).SetDuration)
	c.vopt.AddFunc("SeekTime", (*models.Mediafile).SetSeekTime)
	c.vopt.AddFunc("Strict", (*models.Mediafile).SetStrict)
	c.vopt.AddFunc("AudioFilter", (*models.Mediafile).SetAudioFilter)
	c.vopt.AddFunc("VideoFilter", (*models.Mediafile).SetVideoFilter)
	c.vopt.AddFunc("CompressionLevel", (*models.Mediafile).SetCompressionLevel)
}

func (c Converter) ConvertVideo(ifile string, options ConversionOptions) (ofilename string, err error) {
	fname := uuid.NewString()

	ofilename = fmt.Sprintf("%s.%s", fname, options.Extension)
	ofilepath := fmt.Sprintf("%s/%s", config.Values.OutputVideoPath, ofilename)

	if err = c.tscoder.Initialize(ifile, ofilepath); err != nil {
		return
	}

	o := reflect.ValueOf(options)
	for i := 0; i < o.NumField(); i++ {
		fname := o.Type().Field(i).Name
		fvalue := o.Field(i).Interface()

		c.vopt.CallFunc(fname, c.tscoder.MediaFile(), fvalue)
	}

	done := c.tscoder.Run(false)
	err = <-done

	return
}

func (c Converter) ConvertImage(ifile string, options ConversionOptions) (ofilename string, err error) {
	fname := uuid.NewString()

	ofilename = fmt.Sprintf("%s.%s", fname, options.Extension)
	ofilepath := fmt.Sprintf("%s/%s", config.Values.OutputImagePath, ofilename)

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	if err = mw.ReadImage(ifile); err != nil {
		return
	}

	if options.Width != 0 && options.Height != 0 {
		if err = mw.ResizeImage(options.Width, options.Height, imagick.FILTER_LANCZOS); err != nil {
			return
		}
	}

	if options.Quality != 0 {
		if err = mw.SetImageCompressionQuality(options.Quality); err != nil {
			return
		}
	}

	mw.WriteImage(ofilepath)
	return
}
