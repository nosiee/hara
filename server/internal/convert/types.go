package convert

type ConversionVideoOptions struct {
	Type   string             `json:"type"`
	Input  InputVideoOptions  `json:"input"`
	Output OutputVideoOptions `json:"output"`
}

type InputVideoOptions struct {
	Extension string `json:"extension"`
}

type OutputVideoOptions struct {
	Name             string `json:"name,omitempty"`
	Extension        string `json:"extension"`
	AspectRatio      string `json:"aspectRatio,omitempty"`
	Resolution       string `json:"resolution,omitempty"`
	VideoBitRate     string `json:"videoBitRate,omitempty"`
	VideoMaxBitRate  int    `json:"setVideoMaxBitRate,omitempty"`
	VideoMinBitRate  int    `json:"setVideoMinBitRate,omitempty"`
	VideoCodec       string `json:"videoCodec,omitempty"`
	VFrames          int    `json:"vFrames,omitempty"`
	FrameRate        int    `json:"frameRate,omitempty"`
	AudioRate        int    `json:"audioRate,omitempty"`
	SkipAudio        bool   `json:"skipAudio,omitempty"`
	SkipVideo        bool   `json:"skipVideo,omitempty"`
	MaxKeyFrame      int    `json:"maxKeyFrame,omitempty"`
	MinKeyFrame      int    `json:"minKeyFrame,omitempty"`
	KeyframeInterval int    `json:"keyframeInterval,omitempty"`
	AudioCodec       string `json:"audioCodec,omitempty"`
	AudioBitRate     string `json:"audioBitRate,omitempty"`
	Channels         int    `json:"channels,omitempty"`
	BufferSize       int    `json:"bufferSize,omitempty"`
	Preset           string `json:"preset,omitempty"`
	Tune             string `json:"tune,omitempty"`
	AudioProfile     string `json:"audioProfile,omitempty"`
	VideoProfile     string `json:"videoProfile,omitempty"`
	Duration         string `json:"duration,omitempty"`
	SeekTime         string `json:"seekTime,omitempty"`
	Strict           int    `json:"strict,omitempty"`
	AudioFilter      string `json:"audioFilter,omitempty"`
	VideoFilter      string `json:"videoFilter,omitempty"`
	CompressionLevel int    `json:"compressionLevel,omitempty"`
}
