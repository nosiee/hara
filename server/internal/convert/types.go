package convert

type ConversionImageOptions struct {
	Extension string `json:"extension"`
	Width     uint   `json:"width"`
	Height    uint   `json:"height"`
	Quality   uint   `json:"quality"`
}

type ConversionVideoOptions struct {
	Extension        string `json:"extension"`
	AspectRatio      string `json:"aspectRatio"`
	Resolution       string `json:"resolution"`
	VideoBitRate     string `json:"videoBitRate"`
	VideoMaxBitRate  int    `json:"setVideoMaxBitRate"`
	VideoMinBitRate  int    `json:"setVideoMinBitRate"`
	VideoCodec       string `json:"videoCodec"`
	VFrames          int    `json:"vFrames"`
	FrameRate        int    `json:"frameRate"`
	AudioRate        int    `json:"audioRate"`
	SkipAudio        bool   `json:"skipAudio"`
	SkipVideo        bool   `json:"skipVideo"`
	MaxKeyFrame      int    `json:"maxKeyFrame"`
	MinKeyFrame      int    `json:"minKeyFrame"`
	KeyframeInterval int    `json:"keyframeInterval"`
	AudioCodec       string `json:"audioCodec"`
	AudioBitRate     string `json:"audioBitRate"`
	Channels         int    `json:"channels"`
	BufferSize       int    `json:"bufferSize"`
	Preset           string `json:"preset"`
	Tune             string `json:"tune"`
	AudioProfile     string `json:"audioProfile"`
	VideoProfile     string `json:"videoProfile"`
	Duration         string `json:"duration"`
	SeekTime         string `json:"seekTime"`
	Strict           int    `json:"strict"`
	AudioFilter      string `json:"audioFilter"`
	VideoFilter      string `json:"videoFilter"`
	CompressionLevel int    `json:"compressionLevel"`
}
