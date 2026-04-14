package types

const (
	TextOrientationHorizontal = int(0) // 横排
	TextOrientationVertical   = int(1) // 竖排
)

// 图片中的文字信息
type ImageText struct {
	Text            string  `json:"text" v:"required|length:1,128" dc:"要添加的文字内容"`                                       // 要添加的文字
	TextOrientation int     `json:"textOrientation" v:"required|in:0,1" dc:"文字布局, 0: 横排, 1: 竖排"`                        // 文字布局
	Size            int     `json:"size" v:"required|min:12|max:128" dc:"字体大小"`                                         // 字体大小，默认为32
	Space           float64 `json:"space" v:"required|min:0.1|max:32" dc:"文字间距"`                                        // 文字之间的间距，默认为1.35
	X               float64 `json:"x" v:"required|min:0.001|max:4096" dc:"文字的x坐标"`                                      // 文字的x坐标，默认为10
	Y               float64 `json:"y" v:"required|min:0.001|max:4096" dc:"文字的y坐标"`                                      // 文字的y坐标，默认为10
	Color           []uint8 `json:"color" v:"required" dc:"文字颜色, NRGBA(r, g, b, a), 每一项取值都是[0, 255], 其中a表示透明度"`         // 文字颜色，默认为黑色
	BorderColor     []uint8 `json:"borderColor" v:"required" dc:"文字描边颜色, NRGBA(r, g, b, a), 每一项取值都是[0, 255], 其中a表示透明度"` // 文字描边颜色，默认为白色
	BorderWidth     int     `json:"borderWidth" v:"required" dc:"文字描边宽度"`                                               // 文字描边宽度，默认为1
}
