package convert

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

func ConvertVideo(fpath string, outputPath string, options *ConversionVideoOptions) (string, error) {
	tcoder := new(transcoder.Transcoder)
	dest := fmt.Sprintf("%s/%s", outputPath, options.Output.Name)

	if err := tcoder.Initialize(fpath, dest); err != nil {
		return "", err
	}

	tcoder.MediaFile().SetVideoBitRate("4M")
	tcoder.MediaFile().SetVideoCodec("libvpx")
	tcoder.MediaFile().SetAudioCodec("libvorbis")
	tcoder.MediaFile().SetSkipAudio(true)

	done := tcoder.Run(true)
	if err := <-done; err != nil {
		return "", err
	}

	return dest, nil
}
