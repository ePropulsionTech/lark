// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateSheetFilterView 筛选范围的设置参考：[筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/create
func (r *DriveService) CreateSheetFilterView(ctx context.Context, request *CreateSheetFilterViewReq, options ...MethodOptionFunc) (*CreateSheetFilterViewResp, *Response, error) {
	if r.cli.mock.mockDriveCreateSheetFilterView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheetFilterView mock enable")
		return r.cli.mock.mockDriveCreateSheetFilterView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateSheetFilterView",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createSheetFilterViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCreateSheetFilterView(f func(ctx context.Context, request *CreateSheetFilterViewReq, options ...MethodOptionFunc) (*CreateSheetFilterViewResp, *Response, error)) {
	r.mockDriveCreateSheetFilterView = f
}

func (r *Mock) UnMockDriveCreateSheetFilterView() {
	r.mockDriveCreateSheetFilterView = nil
}

type CreateSheetFilterViewReq struct {
	SpreadSheetToken string  `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string  `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b**12"
	FilterViewID     *string `json:"filter_view_id,omitempty"`   // 筛选视图 id, 示例值："pH9hbVcCXA"
	FilterViewName   *string `json:"filter_view_name,omitempty"` // 筛选视图名字, 示例值："筛选视图 1"
	Range            *string `json:"range,omitempty"`            // 筛选视图的筛选范围, 示例值："0b**12!C1:H14"
}

type createSheetFilterViewResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateSheetFilterViewResp `json:"data,omitempty"`
}

type CreateSheetFilterViewResp struct {
	FilterView *CreateSheetFilterViewRespFilterView `json:"filter_view,omitempty"` // 创建的筛选视图的 id 、name、range
}

type CreateSheetFilterViewRespFilterView struct {
	FilterViewID   string `json:"filter_view_id,omitempty"`   // 筛选视图 id
	FilterViewName string `json:"filter_view_name,omitempty"` // 筛选视图名字
	Range          string `json:"range,omitempty"`            // 筛选视图的筛选范围
}
