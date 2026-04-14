package types

type Input struct {
	Prompt        string `json:"prompt" v:"required|length:1,800" dc:"文本提示词, 支持中英文, 长度不超过800个字符"`
	ImageURL      string `json:"imageURL" dc:"图片URL, 支持JPG/JPEG/PNG/BMP格式, 图片大小不超过10MB, 分辨率: 360≤图像边长≤2000, 单位像素"`
	FirstFrameURL string `json:"firstFrameURL" dc:"视频第一帧URL, 支持JPG/JPEG/PNG/BMP格式, 图片大小不超过10MB, 分辨率: 360≤图像边长≤2000, 单位像素"`
	LastFrameURL  string `json:"lastFrameURL" dc:"视频最后一帧URL, 支持JPG/JPEG/PNG/BMP格式, 图片大小不超过10MB, 分辨率: 360≤图像边长≤2000, 单位像素"`
}

type Params struct {
	Resolution   string `json:"resolution" dc:"视频分辨率, 支持以下分辨率: 720p, 480p"`
	PromptExtend *bool  `json:"promptExtend" dc:"是否开启prompt智能改写。开启后使用大模型对输入prompt进行智能改写。对于较短的prompt生成效果提升明显, 但会增加耗时"`
	Duration     *int   `json:"duration" dc:"视频时长, 单位: 秒, 范围: 3~5"`
	Seed         *int   `json:"seed" dc:"随机数种子，用于控制模型生成内容的随机性, 范围: 0~2147483647"`
}
