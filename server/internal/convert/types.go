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
	Name         string `json:"name,omitempty"`
	Extension    string `json:"extension"`
	VideoBitRate string `json:"videoBitRate,omitempty"`
	AudioBitRate string `json:"audioBitRate,omitempty"`
	AudioCodec   string `json:"audioCodec,omitempty"`
	VideoCodec   string `json:"videoCodec,omitempty"`
	Channels     int    `json:"channels,omitempty"`
	Volume       string `json:"volume,omitempty"`
	Resolution   string `json:"resolution,omitempty"`
	AspectRatio  string `json:"aspectRatio,omitempty"`
	FrameRate    int    `json:"frameRate,omitempty"`
	TrimStart    string `json:"trimStart,omitempty"`
	TrimEnd      string `json:"trimEnd,omitempty"`
}
