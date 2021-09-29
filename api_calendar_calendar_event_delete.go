// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteCalendarEvent
//
// 该接口用于以当前身份（应用 / 用户）删除日历上的一个日程。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份必须对日历有 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
// 当前身份必须是日程的组织者。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/delete
func (r *CalendarService) DeleteCalendarEvent(ctx context.Context, request *DeleteCalendarEventReq, options ...MethodOptionFunc) (*DeleteCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarDeleteCalendarEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#DeleteCalendarEvent mock enable")
		return r.cli.mock.mockCalendarDeleteCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "DeleteCalendarEvent",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockCalendarDeleteCalendarEvent(f func(ctx context.Context, request *DeleteCalendarEventReq, options ...MethodOptionFunc) (*DeleteCalendarEventResp, *Response, error)) {
	r.mockCalendarDeleteCalendarEvent = f
}

func (r *Mock) UnMockCalendarDeleteCalendarEvent() {
	r.mockCalendarDeleteCalendarEvent = nil
}

type DeleteCalendarEventReq struct {
	NeedNotification *bool  `query:"need_notification" json:"-"` // 删除日程是否给日程参与人发送bot通知，默认为true, 示例值：false, 可选值有: `true`：true, `false`：false
	CalendarID       string `path:"calendar_id" json:"-"`        // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID          string `path:"event_id" json:"-"`           // 日程ID, 示例值："xxxxxxxxx_0"
}

type deleteCalendarEventResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarEventResp `json:"data,omitempty"`
}

type DeleteCalendarEventResp struct{}
