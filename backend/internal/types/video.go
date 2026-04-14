package types

// VideoInfo 视频信息
type VideoInfo struct {
	Title    string `json:"title" dc:"视频标题"`
	Author   string `json:"author" dc:"视频作者"`
	Width    int    `json:"width" dc:"视频宽度"`
	Height   int    `json:"height" dc:"视频高度"`
	Duration int    `json:"duration" dc:"视频时长，单位: 秒"`
}

type Subtitle struct {
	StartTime string `json:"startTime" v:"required|length:12,12" dc:"开始时间, 格式: 00:00:00,000"`
	EndTime   string `json:"endTime" v:"required|length:12,12" dc:"开始时间, 格式: 00:00:05,000"`
	Text      string `json:"Text" v:"required|length:1,1024" dc:"字幕内容，支持使用\n换行, 最多2行, 同时显示中英文字幕，如: 你好\nHello"`
}

// SubtitleStyle 字幕信息
type SubtitleStyle struct {
	Subtitles   []Subtitle `json:"subtitles" dc:"字幕内容"`
	Font        int        `json:"font" v:"required|in:1,2,3,4,5" dc:"字体类型"`
	FontSize    int        `json:"fontSize" v:"required|min:6|max:256" dc:"字体大小"`
	Color       Color      `json:"color" dc:"文字颜色"`
	BorderColor Color      `json:"borderColor" dc:"文字描边颜色"`
	BorderWidth int        `json:"borderWidth" v:"required|min:0|max:128" dc:"文字描边宽度"`
}

type SubtitleStyleEx struct {
	Text        string  `json:"Text" v:"required|length:1,64" dc:"中文字幕，如：你好"`
	TextEn      string  `json:"textEn" v:"length:1,512" dc:"英文字幕，可以为空，如：Hello"`
	StartTime   float64 `json:"startTime" v:"required|min:-1|max:6000" dc:"字幕开始时间, 单位：秒"`
	EndTime     float64 `json:"endTime" v:"required|min:-1|max:6000" dc:"字幕结束时间, 单位：秒"`
	Font        int     `json:"font" v:"required|in:1,2,3,4,5" dc:"字体类型"`
	FontSize    int     `json:"fontSize" v:"required|min:6|max:256" dc:"字体大小"`
	Color       string  `json:"color" v:"required|length:6,7" dc:"文字颜色，示例：F0DAC5或者#F0DAC5"`
	BorderColor string  `json:"borderColor" v:"required|length:6,7" dc:"文字描边颜色，示例：9C2B2B或者#9C2B2B"`
	BorderWidth int     `json:"borderWidth" v:"required|min:0|max:128" dc:"文字描边宽度"`
}

type AutoSubtitleStyle struct {
	Text        string `json:"Text" v:"required|length:1,10000" dc:"字幕内容"`
	TextEn      string `json:"textEn" v:"length:1,80000" dc:"英文字幕，可以为空，如：Hello"`
	Font        int    `json:"font" v:"required|in:1,2,3,4,5" dc:"字体类型"`
	FontSize    int    `json:"fontSize" v:"required|min:6|max:256" dc:"字体大小"`
	Color       string `json:"color" v:"required|length:6,7" dc:"文字颜色，示例：F0DAC5或者#F0DAC5"`
	BorderColor string `json:"borderColor" v:"required|length:6,7" dc:"文字描边颜色，示例：9C2B2B或者#9C2B2B"`
	BorderWidth int    `json:"borderWidth" v:"required|min:0|max:128" dc:"文字描边宽度"`
}

type TextStyle struct {
	Text        string  `json:"text" v:"required|length:1,256" dc:"文字内容"`
	X           string  `json:"x" v:"required|length:1,16" dc:"文字在视频中的位置, 示例: 100, 表示距离视频左侧100像素, (w-text_w)/2, 表示左右居中"`
	Y           string  `json:"y" v:"required|length:1,16" dc:"文字在视频中的位置, 示例: 100, 表示距离视频顶部100像素, h-100, 表示离底部100像素"`
	Font        int     `json:"font" v:"required|in:1,2,3,4,5" dc:"字体类型"`
	FontSize    int     `json:"fontSize" v:"required|min:6|max:256" dc:"字体大小"`
	Color       Color   `json:"color" dc:"文字颜色"`
	ShadowColor Color   `json:"shadowColor" dc:"文字阴影颜色"`
	ShadowX     int     `json:"shadowX" v:"required|min:0|max:512" dc:"文字阴影X轴偏移"`
	ShadowY     int     `json:"shadowY" v:"required|min:0|max:512" dc:"文字阴影Y轴偏移"`
	StartTime   float64 `json:"startTime" v:"required|min:-1|max:6000" dc:"文字开始显示的时间（秒）, -1表示一直显示"`
	EndTime     float64 `json:"endTime" v:"required|min:-1|max:6000" dc:"文字结束显示的时间（秒）, -1表示一直显示"`
}

type TextStyleEx struct {
	Text        string  `json:"text" v:"required|length:1,256" dc:"字幕内容"`
	StartTime   float64 `json:"startTime" v:"required|min:-1|max:6000" dc:"字幕开始显示的时间（秒）, -1表示一直显示"`
	EndTime     float64 `json:"endTime" v:"required|min:-1|max:6000" dc:"字幕结束显示的时间（秒）, -1表示一直显示"`
	X           string  `json:"x" v:"required|length:1,16" dc:"字幕在视频中的位置, 示例: 100, 表示距离视频左侧100像素, (w-text_w)/2, 表示左右居中"`
	Y           string  `json:"y" v:"required|length:1,16" dc:"字幕在视频中的位置, 示例: 100, 表示距离视频顶部100像素, h-100, 表示离底部100像素"`
	Font        int     `json:"font" v:"required|in:1,2,3,4,5" dc:"字体类型"`
	FontSize    int     `json:"fontSize" v:"required|min:6|max:256" dc:"字体大小"`
	Color       string  `json:"color" v:"required|length:6,7" dc:"字幕颜色，示例：F0DAC5或者#F0DAC5"`
	ShadowColor string  `json:"shadowColor" v:"required|length:6,7" dc:"字幕阴影颜色，示例：9C2B2B或者#9C2B2B"`
	ShadowX     int     `json:"shadowX" v:"required|min:0|max:512" dc:"字幕阴影X轴偏移"`
	ShadowY     int     `json:"shadowY" v:"required|min:0|max:512" dc:"字幕阴影Y轴偏移"`
}

type AutoTextStyle struct {
	Text        string `json:"text" v:"required|length:1,10000" dc:"字幕内容"`
	X           string `json:"x" v:"required|length:1,16" dc:"文字在视频中的位置, 示例: 100, 表示距离视频左侧100像素, (w-text_w)/2, 表示左右居中"`
	Y           string `json:"y" v:"required|length:1,16" dc:"文字在视频中的位置, 示例: 100, 表示距离视频顶部100像素, h-100, 表示离底部100像素"`
	Font        int    `json:"font" v:"required|in:1,2,3,4,5" dc:"字体类型"`
	FontSize    int    `json:"fontSize" v:"required|min:6|max:256" dc:"字体大小"`
	Color       string `json:"color" v:"required|length:6,7" dc:"文字颜色，示例：F0DAC5或者#F0DAC5"`
	ShadowColor string `json:"shadowColor" v:"required|length:6,7" dc:"文字阴影颜色，示例：9C2B2B或者#9C2B2B"`
	ShadowX     int    `json:"shadowX" v:"required|min:0|max:512" dc:"文字阴影X轴偏移"`
	ShadowY     int    `json:"shadowY" v:"required|min:0|max:512" dc:"文字阴影Y轴偏移"`
}

type PIPStyle struct {
	OverlayURL  string  `json:"overlayURL" v:"required|valid-url" dc:"视频文件路径，添加的画中画视频"`
	ColorKey    string  `json:"colorKey" v:"required|length:6,7" dc:"指定黑色转成透明色，格式：#F0DAC5"`
	ScaleWidth  int     `json:"scaleWidth" v:"required|min:0|max:10000" dc:"指定视频overlayURL的宽度，单位: 像素"`
	ScaleHeight int     `json:"scaleHeight" v:"required|min:0|max:10000" dc:"指定视频overlayURL的高度，单位: 像素"`
	X           string  `json:"x" v:"required|length:1,22" dc:"指定视频overlayURL的x坐标，单位: 像素"`
	Y           string  `json:"y" v:"required|length:1,22" dc:"指定视频overlayURL的y坐标，单位: 像素"`
	StartTime   float64 `json:"startTime" v:"required|min:0|max:10000" dc:"指定视频overlayURL的开始时间，单位: 秒"`
	EndTime     float64 `json:"endTime" v:"required|min:0|max:10000" dc:"指定视频overlayURL的结束时间，单位: 秒"`
}
