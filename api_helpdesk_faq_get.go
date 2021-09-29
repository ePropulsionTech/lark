// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHelpdeskFAQ 该接口用于获取服务台知识库详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/get
func (r *HelpdeskService) GetHelpdeskFAQ(ctx context.Context, request *GetHelpdeskFAQReq, options ...MethodOptionFunc) (*GetHelpdeskFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskFAQ != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskFAQ mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskFAQ",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs/:id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskGetHelpdeskFAQ(f func(ctx context.Context, request *GetHelpdeskFAQReq, options ...MethodOptionFunc) (*GetHelpdeskFAQResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskFAQ = f
}

func (r *Mock) UnMockHelpdeskGetHelpdeskFAQ() {
	r.mockHelpdeskGetHelpdeskFAQ = nil
}

type GetHelpdeskFAQReq struct {
	ID string `path:"id" json:"-"` // 知识库ID, 示例值："6856395634652479491"
}

type getHelpdeskFAQResp struct {
	Code int64               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskFAQResp `json:"data,omitempty"`
}

type GetHelpdeskFAQResp struct {
	FAQ *GetHelpdeskFAQRespFAQ `json:"faq,omitempty"` // 知识库详情
}

type GetHelpdeskFAQRespFAQ struct {
	FAQID          string                           `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                           `json:"id,omitempty"`              // 知识库旧版ID，请使用faq_id
	HelpdeskID     string                           `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                           `json:"question,omitempty"`        // 问题
	Answer         string                           `json:"answer,omitempty"`          // 答案
	AnswerRichtext string                           `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int64                            `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int64                            `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory              `json:"categories,omitempty"`      // 分类
	Tags           []string                         `json:"tags,omitempty"`            // 关联词列表
	ExpireTime     int64                            `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *GetHelpdeskFAQRespFAQUpdateUser `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *GetHelpdeskFAQRespFAQCreateUser `json:"create_user,omitempty"`     // 创建用户
}

type GetHelpdeskFAQRespFAQUpdateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}

type GetHelpdeskFAQRespFAQCreateUser struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
}
