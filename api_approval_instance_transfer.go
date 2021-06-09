// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// TransferApprovalInstance
//
// 对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDNyUjL1QjM14SN0ITN
func (r *ApprovalService) TransferApprovalInstance(ctx context.Context, request *TransferApprovalInstanceReq, options ...MethodOptionFunc) (*TransferApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalTransferApprovalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#TransferApprovalInstance mock enable")
		return r.cli.mock.mockApprovalTransferApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "TransferApprovalInstance",
		Method:                "POST",
		URL:                   "https://www.feishu.cn/approval/openapi/v2/instance/transfer",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(transferApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApprovalTransferApprovalInstance(f func(ctx context.Context, request *TransferApprovalInstanceReq, options ...MethodOptionFunc) (*TransferApprovalInstanceResp, *Response, error)) {
	r.mockApprovalTransferApprovalInstance = f
}

func (r *Mock) UnMockApprovalTransferApprovalInstance() {
	r.mockApprovalTransferApprovalInstance = nil
}

type TransferApprovalInstanceReq struct {
	ApprovalCode   string  `json:"approval_code,omitempty"`    // 审批定义 Code
	InstanceCode   string  `json:"instance_code,omitempty"`    // 审批实例 Code
	UserID         string  `json:"user_id,omitempty"`          // 操作用户
	TaskID         string  `json:"task_id,omitempty"`          // 任务 ID
	Comment        *string `json:"comment,omitempty"`          // 意见
	TransferUserID string  `json:"transfer_user_id,omitempty"` // 被转交人唯一 ID
	OpenID         *string `json:"open_id,omitempty"`          // 用户open_id <br>如果没有user_id，必须要有open_id
	TransferOpenID *string `json:"transfer_open_id,omitempty"` // 被转交人open_id <br>如果没有transfer_user_id，必须要有transfer_open_id
}

type transferApprovalInstanceResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码，非0表示失败
	Msg  string                        `json:"msg,omitempty"`  // 返回码的描述
	Data *TransferApprovalInstanceResp `json:"data,omitempty"`
}

type TransferApprovalInstanceResp struct{}
