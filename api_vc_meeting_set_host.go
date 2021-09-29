// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SetVCHostMeeting 设置会议的主持人
//
// 发起设置主持人的操作者必须具有相应的权限（如果操作者为用户，必须是会中当前主持人）；该操作使用CAS并发安全机制，需传入会中当前主持人，如果操作失败可使用返回的最新数据重试
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/set_host
func (r *VCService) SetVCHostMeeting(ctx context.Context, request *SetVCHostMeetingReq, options ...MethodOptionFunc) (*SetVCHostMeetingResp, *Response, error) {
	if r.cli.mock.mockVCSetVCHostMeeting != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#SetVCHostMeeting mock enable")
		return r.cli.mock.mockVCSetVCHostMeeting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "SetVCHostMeeting",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/meetings/:meeting_id/set_host",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(setVCHostMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockVCSetVCHostMeeting(f func(ctx context.Context, request *SetVCHostMeetingReq, options ...MethodOptionFunc) (*SetVCHostMeetingResp, *Response, error)) {
	r.mockVCSetVCHostMeeting = f
}

func (r *Mock) UnMockVCSetVCHostMeeting() {
	r.mockVCSetVCHostMeeting = nil
}

type SetVCHostMeetingReq struct {
	UserIDType  *IDType                         `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`,, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MeetingID   string                          `path:"meeting_id" json:"-"`     // 会议ID（视频会议的唯一标识，视频会议开始后才会产生）, 示例值："6911188411932033028"
	HostUser    *SetVCHostMeetingReqHostUser    `json:"host_user,omitempty"`     // 将要设置的主持人
	OldHostUser *SetVCHostMeetingReqOldHostUser `json:"old_host_user,omitempty"` // 当前主持人（CAS并发安全：如果和会中当前主持人不符则会设置失败，可使用返回的最新数据重新设置）
}

type SetVCHostMeetingReqHostUser struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int64  `json:"user_type,omitempty"` // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type SetVCHostMeetingReqOldHostUser struct {
	ID       *string `json:"id,omitempty"`        // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType *int64  `json:"user_type,omitempty"` // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type setVCHostMeetingResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *SetVCHostMeetingResp `json:"data,omitempty"`
}

type SetVCHostMeetingResp struct {
	HostUser *SetVCHostMeetingRespHostUser `json:"host_user,omitempty"` // 会中当前主持人
}

type SetVCHostMeetingRespHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户ID
	UserType int64  `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}
