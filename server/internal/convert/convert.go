package convert

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"hara/internal/config"
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

func ConvertVideo(ifile string, options ConversionVideoOptions) (ofilename string, err error) {
	ofilename = fmt.Sprintf("%s.%s", getRandomFileName(), options.Extension)
	ofilepath := fmt.Sprintf("%s/%s", config.Values.OutputVideoPath, ofilename)

	if err = tscoder.Initialize(ifile, ofilepath); err != nil {
		return
	}

	o := reflect.ValueOf(options)
	for i := 0; i < o.NumField(); i++ {
		fname := o.Type().Field(i).Name
		fvalue := o.Field(i).Interface()

		vopt.CallFunc(fname, tscoder.MediaFile(), fvalue)
	}

	done := tscoder.Run(false)
	err = <-done

	return
}

func ConvertImage(ifile string, options ConversionImageOptions) (ofilename string, err error) {
	ofilename = fmt.Sprintf("%s.%s", getRandomFileName(), options.Extension)
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

func getRandomFileName() string {
	u := make([]byte, 8)
	_, _ = rand.Read(u)

	u[7] = (u[7] | 0x80) & 0xBF
	u[5] = (u[5] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}
