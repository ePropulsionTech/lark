// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteMailGroup 删除一个邮件组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/delete
func (r *MailService) DeleteMailGroup(ctx context.Context, request *DeleteMailGroupReq, options ...MethodOptionFunc) (*DeleteMailGroupResp, *Response, error) {
	if r.cli.mock.mockMailDeleteMailGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailGroup mock enable")
		return r.cli.mock.mockMailDeleteMailGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "DeleteMailGroup",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteMailGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMailDeleteMailGroup(f func(ctx context.Context, request *DeleteMailGroupReq, options ...MethodOptionFunc) (*DeleteMailGroupResp, *Response, error)) {
	r.mockMailDeleteMailGroup = f
}

func (r *Mock) UnMockMailDeleteMailGroup() {
	r.mockMailDeleteMailGroup = nil
}

type DeleteMailGroupReq struct {
	MailGroupID string `path:"mailgroup_id" json:"-"` // 邮件组ID或者邮件组地址, 示例值："xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
}

type deleteMailGroupResp struct {
	Code int64                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailGroupResp `json:"data,omitempty"`
}

type DeleteMailGroupResp struct{}
