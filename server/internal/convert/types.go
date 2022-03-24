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
	Name          string `json:"name"`
	VideoBitrate  string `json:"video-bitrate"`
	AudioBitrate  string `json:"audio-bitrate"`
	AudioCodec    string `json:"audio-codec"`
	VideoCodec    string `json:"video-codec"`
	AudioChannels string `json:"channels"`
	SampleRate    string `json:"sample-rate"`
	Volume        string `json:"volume"`
	Resolution    string `json:"resolution"`
	AspectRatio   string `json:"aspect-ratio"`
	FPS           string `json:"fps"`
	TrimStart     string `json:"trim-start"`
	TrimEnd       string `json:"trim-end"`
}
