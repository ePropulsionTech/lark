// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CompleteTask 该接口用于将接任务状态修改为已完成
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/complete
func (r *TaskService) CompleteTask(ctx context.Context, request *CompleteTaskReq, options ...MethodOptionFunc) (*CompleteTaskResp, *Response, error) {
	if r.cli.mock.mockTaskCompleteTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CompleteTask mock enable")
		return r.cli.mock.mockTaskCompleteTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CompleteTask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/complete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(completeTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockTaskCompleteTask(f func(ctx context.Context, request *CompleteTaskReq, options ...MethodOptionFunc) (*CompleteTaskResp, *Response, error)) {
	r.mockTaskCompleteTask = f
}

func (r *Mock) UnMockTaskCompleteTask() {
	r.mockTaskCompleteTask = nil
}

type CompleteTaskReq struct {
	TaskID string `path:"task_id" json:"-"` // 任务 ID, 示例值："bb54ab99-d360-434f-bcaa-a4cc4c05840e"
}

type completeTaskResp struct {
	Code int64             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *CompleteTaskResp `json:"data,omitempty"`
}

type CompleteTaskResp struct{}
