// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateBitableView 在数据表中新增一个视图
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/create
func (r *BitableService) CreateBitableView(ctx context.Context, request *CreateBitableViewReq, options ...MethodOptionFunc) (*CreateBitableViewResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableView mock enable")
		return r.cli.mock.mockBitableCreateBitableView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableView",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableCreateBitableView(f func(ctx context.Context, request *CreateBitableViewReq, options ...MethodOptionFunc) (*CreateBitableViewResp, *Response, error)) {
	r.mockBitableCreateBitableView = f
}

func (r *Mock) UnMockBitableCreateBitableView() {
	r.mockBitableCreateBitableView = nil
}

type CreateBitableViewReq struct {
	AppToken string  `path:"app_token" json:"-"`  // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string  `path:"table_id" json:"-"`   // table id, 示例值："tblsRc9GRRXKqhvW"
	ViewID   *string `json:"view_id,omitempty"`   // 视图Id, 示例值："vewTpR1urY"
	ViewName *string `json:"view_name,omitempty"` // 视图名字, 示例值："甘特视图1"
	ViewType *string `json:"view_type,omitempty"` // 视图类型, 示例值："gantt"
}

type createBitableViewResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableViewResp `json:"data,omitempty"`
}

type CreateBitableViewResp struct {
	View *CreateBitableViewRespApptableview `json:"view,omitempty"` // 视图
}

type CreateBitableViewRespApptableview struct {
	ViewID   string `json:"view_id,omitempty"`   // 视图Id
	ViewName string `json:"view_name,omitempty"` // 视图名字
	ViewType string `json:"view_type,omitempty"` // 视图类型
}
