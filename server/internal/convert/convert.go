package convert

import (
	"fmt"
	"reflect"

	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
	"gopkg.in/gographics/imagick.v3/imagick"
)

var (
	tscoder = new(transcoder.Transcoder)
	vopt    = make(VideoOptionsMap, 28)
)

func init() {
	vopt.AddFunc("AspectRatio", (*models.Mediafile).SetAspect)
	vopt.AddFunc("Resolution", (*models.Mediafile).SetResolution)
	vopt.AddFunc("VideoBitRate", (*models.Mediafile).SetVideoBitRate)
	vopt.AddFunc("VideoMaxBitRate", (*models.Mediafile).SetVideoMaxBitrate)
	vopt.AddFunc("VideoMinBitRate", (*models.Mediafile).SetVideoMinBitRate)
	vopt.AddFunc("VideoCodec", (*models.Mediafile).SetVideoCodec)
	vopt.AddFunc("VFrames", (*models.Mediafile).SetVframes)
	vopt.AddFunc("FrameRate", (*models.Mediafile).SetFrameRate)
	vopt.AddFunc("AudioRate", (*models.Mediafile).SetAudioRate)
	vopt.AddFunc("SkipVideo", (*models.Mediafile).SetSkipVideo)
	vopt.AddFunc("SkipAudio", (*models.Mediafile).SetSkipAudio)
	vopt.AddFunc("MaxKeyFrame", (*models.Mediafile).SetMaxKeyFrame)
	vopt.AddFunc("MinKeyFrame", (*models.Mediafile).SetMinKeyFrame)
	vopt.AddFunc("KeyframeInterval", (*models.Mediafile).SetKeyframeInterval)
	vopt.AddFunc("AudioCodec", (*models.Mediafile).SetAudioCodec)
	vopt.AddFunc("AudioBitRate", (*models.Mediafile).SetAudioBitRate)
	vopt.AddFunc("Channels", (*models.Mediafile).SetAudioChannels)
	vopt.AddFunc("BufferSize", (*models.Mediafile).SetBufferSize)
	vopt.AddFunc("Preset", (*models.Mediafile).SetPreset)
	vopt.AddFunc("Tune", (*models.Mediafile).SetTune)
	vopt.AddFunc("AudioProfile", (*models.Mediafile).SetAudioProfile)
	vopt.AddFunc("VideoProfile", (*models.Mediafile).SetVideoProfile)
	vopt.AddFunc("Duration", (*models.Mediafile).SetDuration)
	vopt.AddFunc("SeekTime", (*models.Mediafile).SetSeekTime)
	vopt.AddFunc("Strict", (*models.Mediafile).SetStrict)
	vopt.AddFunc("AudioFilter", (*models.Mediafile).SetAudioFilter)
	vopt.AddFunc("VideoFilter", (*models.Mediafile).SetVideoFilter)
	vopt.AddFunc("CompressionLevel", (*models.Mediafile).SetCompressionLevel)
}

func ConvertVideo(ifile string, options ConversionVideoOptions) error {
	// TODO: load outputpath from config(?)
	ofile := fmt.Sprintf("%s/%s", "output", options.Name)

	if err := tscoder.Initialize(ifile, ofile); err != nil {
		return err
	}

	o := reflect.ValueOf(options)
	for i := 0; i < o.NumField(); i++ {
		fname := o.Type().Field(i).Name
		fvalue := o.Field(i).Interface()

		vopt.CallFunc(fname, tscoder.MediaFile(), fvalue)
	}

	done := tscoder.Run(false)
	return <-done
}

func ConvertImage(ifile string, options ConversionImageOptions) (err error) {
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

	// TODO: load outputpath from config(?)
	ofile := fmt.Sprintf("%s/%s", "output", options.Name)
	mw.WriteImage(ofile)

	return
}
