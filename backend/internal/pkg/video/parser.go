package video

import (
	"cozeos/internal/types"
	"encoding/json"
	"fmt"
	"strconv"
)

// FFmpegOutput 用于解析 FFmpeg 的 JSON 输出
type FFprobeOutput struct {
	Streams []struct {
		CodecType string `json:"codec_type"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Tags      struct {
			Title  string `json:"title"`
			Artist string `json:"artist"`
		} `json:"tags"`
	} `json:"streams"`
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

// ParseVideoInfo 从 FFmpeg 的 JSON 输出中解析视频信息
func ParseVideoInfo(ffprobeOutput []byte) (*types.VideoInfo, error) {
	var output FFprobeOutput
	if err := json.Unmarshal(ffprobeOutput, &output); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ffmpeg output: %v", err)
	}

	var videoInfo types.VideoInfo

	// 遍历流信息，找到视频流
	for _, stream := range output.Streams {
		if stream.CodecType == "video" {
			videoInfo.Width = stream.Width
			videoInfo.Height = stream.Height
			videoInfo.Title = stream.Tags.Title
			videoInfo.Author = stream.Tags.Artist
			break
		}
	}

	// 解析视频时长
	duration, err := strconv.ParseFloat(output.Format.Duration, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse duration: %v", err)
	}
	videoInfo.Duration = int(duration)

	return &videoInfo, nil
}
