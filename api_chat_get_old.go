// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetChatOld
//
// :::html
// <md-alert type="error">
// 为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get)
// </md-alert>
// :::
// 获取群名称、群主 ID、成员列表 ID 等群基本信息。
// - 需要启用机器人能力；机器人必须在群里
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN
//
// Deprecated
func (r *ChatService) GetChatOld(ctx context.Context, request *GetChatOldReq, options ...MethodOptionFunc) (*GetChatOldResp, *Response, error) {
	if r.cli.mock.mockChatGetChatOld != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#GetChatOld mock enable")
		return r.cli.mock.mockChatGetChatOld(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChatOld",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/chat/v4/info",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getChatOldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockChatGetChatOld(f func(ctx context.Context, request *GetChatOldReq, options ...MethodOptionFunc) (*GetChatOldResp, *Response, error)) {
	r.mockChatGetChatOld = f
}

func (r *Mock) UnMockChatGetChatOld() {
	r.mockChatGetChatOld = nil
}

type GetChatOldReq struct {
	ChatID string `query:"chat_id" json:"-"` // 群 ID
}

type getChatOldResp struct {
	Code int64           `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 返回码描述
	Data *GetChatOldResp `json:"data,omitempty"` // -
}

type GetChatOldResp struct {
	Avatar                   string                  `json:"avatar,omitempty"`                      // 群头像
	Description              string                  `json:"description,omitempty"`                 // 群描述
	I18nNames                map[string]string       `json:"i18n_names,omitempty"`                  // 群国际化名称（设置了国际化名称才会有这个字段）
	ChatID                   string                  `json:"chat_id,omitempty"`                     // 群 ID
	Members                  []*GetChatOldRespMember `json:"members,omitempty"`                     // 成员列表
	Name                     string                  `json:"name,omitempty"`                        // 群名称，类型为group时有效
	Type                     string                  `json:"type,omitempty"`                        // 群类型，group表示群聊，p2p表示单聊
	OwnerUserID              string                  `json:"owner_user_id,omitempty"`               // 群主的 user_id（机器人是群主的时候没有这个字段）
	OwnerOpenID              string                  `json:"owner_open_id,omitempty"`               // 群主的 open_id （机器人是群主的时候没有这个字段）
	OnlyOwnerEdit            bool                    `json:"only_owner_edit,omitempty"`             // 是否仅群主可编辑群信息，群信息包括头像、名称、描述、公告
	OnlyOwnerAdd             bool                    `json:"only_owner_add,omitempty"`              // 是否仅群主可以添加人
	ShareAllowed             bool                    `json:"share_allowed,omitempty"`               // 是否允许分享群
	AddMemberVerify          bool                    `json:"add_member_verify,omitempty"`           // 是否开启入群验证
	OnlyOwnerAtAll           bool                    `json:"only_owner_at_all,omitempty"`           // 是否仅群主@all
	SendMessagePermission    string                  `json:"send_message_permission,omitempty"`     // 允许谁发送消息<br>all: 所有人<br>  owner： 仅群主<br>selected_member: 指定成员
	JoinMessageVisibility    MessageVisibility       `json:"join_message_visibility,omitempty"`     // 成员入群通知<br>all：所有人 <br>owner：仅群主 <br>not_anyone：不通知任何人"
	LeaveMessageVisibility   MessageVisibility       `json:"leave_message_visibility,omitempty"`    // 成员退群通知<br>all：所有人 <br>owner：仅群主 <br>not_anyone：不通知任何人
	GroupEmailEnabled        bool                    `json:"group_email_enabled,omitempty"`         // 是否开启群邮件
	SendGroupEmailPermission string                  `json:"send_group_email_permission,omitempty"` // 发送群邮件的权限<br>owner：仅群主  <br>group_member：群组内成员<br>tenant_member：团队成员 <br>all：所有人
}

type GetChatOldRespMember struct {
	OpenID string `json:"open_id,omitempty"` // 某成员的open_id
	UserID string `json:"user_id,omitempty"` // 某成员的user_id
}
