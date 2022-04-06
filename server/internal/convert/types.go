package convert

type ConversionOptions struct {
	Extension        string `option:"ext"`
	Lifetime         uint64 `option:"lifetime"`
	Width            uint   `option:"width"`
	Height           uint   `option:"height"`
	Quality          uint   `option:"quality"`
	AspectRatio      string `option:"aration"`
	Resolution       string `option:"resolution"`
	VideoBitRate     string `option:"vbitrate"`
	VideoMaxBitRate  int    `option:"vmaxbitrate"`
	VideoMinBitRate  int    `option:"vminbitrate"`
	VideoCodec       string `option:"vcodec"`
	VFrames          int    `option:"vframes"`
	FrameRate        int    `option:"framerate"`
	AudioRate        int    `option:"arate"`
	SkipAudio        bool   `option:"skipaudio"`
	SkipVideo        bool   `option:"skipvideo"`
	MaxKeyFrame      int    `option:"maxkeyframe"`
	MinKeyFrame      int    `option:"minkeyframe"`
	KeyframeInterval int    `option:"kfinterval"`
	AudioCodec       string `option:"acodec"`
	AudioBitRate     string `option:"abitrate"`
	Channels         int    `option:"channels"`
	BufferSize       int    `option:"buffersize"`
	Preset           string `option:"preset"`
	Tune             string `option:"tune"`
	AudioProfile     string `option:"aprofile"`
	VideoProfile     string `option:"vprofile"`
	Duration         string `option:"duration"`
	SeekTime         string `option:"seektime"`
	Strict           int    `option:"strict"`
	AudioFilter      string `option:"afileter"`
	VideoFilter      string `option:"vfilter"`
	CompressionLevel int    `option:"clevel"`
}
