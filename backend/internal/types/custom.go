package types

import "time"

// BotCustom 智能体定制信息
type BotCustom struct {
	ID           uint      `json:"id" dc:"定制ID"`
	ContactType  string    `json:"contactType" dc:"联系方式类型，phone或wechat"`
	ContactValue string    `json:"contactValue" dc:"联系方式值"`
	VideoLink    string    `json:"videoLink" dc:"参考视频地址"`
	Description  string    `json:"description" dc:"定制描述信息"`
	Budget       string    `json:"budget" dc:"价格区间"`
	Status       string    `json:"status" dc:"Custom requirement processing status: Pending, Processing, Completed"`
	CreatedAt    time.Time `json:"createAt" dc:"创建时间"`
	UpdatedAt    time.Time `json:"updateAt" dc:"更新时间"`
}

// CreateBotCustomParams 创建智能体定制参数
type CreateBotCustomParams struct {
	ContactType  string `json:"contactType" v:"required|in:phone,wechat" dc:"联系方式类型，phone或wechat"`
	ContactValue string `json:"contactValue" v:"required|length:1,128" dc:"联系方式值"`
	VideoLink    string `json:"videoLink" v:"required|length:1,1024" dc:"参考视频地址"`
	Description  string `json:"description" v:"required|length:1,1024" dc:"定制描述信息"`
	Budget       string `json:"budget" v:"required|length:1,128" dc:"价格区间"`
}
