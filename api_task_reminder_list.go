// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetTaskReminderList 返回提醒时间列表，支持分页，最大值为50
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/list
func (r *TaskService) GetTaskReminderList(ctx context.Context, request *GetTaskReminderListReq, options ...MethodOptionFunc) (*GetTaskReminderListResp, *Response, error) {
	if r.cli.mock.mockTaskGetTaskReminderList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#GetTaskReminderList mock enable")
		return r.cli.mock.mockTaskGetTaskReminderList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTaskReminderList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/reminders",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getTaskReminderListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockTaskGetTaskReminderList(f func(ctx context.Context, request *GetTaskReminderListReq, options ...MethodOptionFunc) (*GetTaskReminderListResp, *Response, error)) {
	r.mockTaskGetTaskReminderList = f
}

func (r *Mock) UnMockTaskGetTaskReminderList() {
	r.mockTaskGetTaskReminderList = nil
}

type GetTaskReminderListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：50, 最大值：`50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："「填写上次返回的page_token」"
	TaskID    string  `path:"task_id" json:"-"`     // 任务 ID, 示例值："0d38e26e-190a-49e9-93a2-35067763ed1f"
}

type getTaskReminderListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetTaskReminderListResp `json:"data,omitempty"`
}

type GetTaskReminderListResp struct {
	Items     []*GetTaskReminderListRespItem `json:"items,omitempty"`      // 返回提醒时间设置列表
	PageToken string                         `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
}

type GetTaskReminderListRespItem struct {
	ID                 string `json:"id,omitempty"`                   // 提醒时间设置的 ID（在删除时候需要使用这个）
	RelativeFireMinute int64  `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间（如提前 30 分钟，截止时间后 30 分钟，则为 -30）
}
