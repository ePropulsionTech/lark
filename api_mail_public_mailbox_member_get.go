// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetPublicMailboxMember 获取公共邮箱单个成员信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-member/get
func (r *MailAPI) GetPublicMailboxMember(ctx context.Context, request *GetPublicMailboxMemberReq) (*GetPublicMailboxMemberResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/members/:member_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getPublicMailboxMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "GetPublicMailboxMember", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetPublicMailboxMemberReq struct {
	UserIDType      *IDType `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	PublicMailboxID string  `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值："xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	MemberID        string  `path:"member_id" json:"-"`         // 公共邮箱内成员唯一标识, 示例值："xxxxxxxxxxxxxxx"
}

type getPublicMailboxMemberResp struct {
	Code int                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetPublicMailboxMemberResp `json:"data,omitempty"` //
}

type GetPublicMailboxMemberResp struct {
	MemberID string       `json:"member_id,omitempty"` // 公共邮箱内成员唯一标识
	UserID   string       `json:"user_id,omitempty"`   // 租户内用户的唯一标识（当成员类型是USER时有值）
	Type     MailUserType `json:"type,omitempty"`      // 成员类型, 可选值有: `USER`：内部用户
}
