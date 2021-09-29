// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetSheetFilterView 获取指定筛选视图 id 的名字和筛选范围。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/get
func (r *DriveService) GetSheetFilterView(ctx context.Context, request *GetSheetFilterViewReq, options ...MethodOptionFunc) (*GetSheetFilterViewResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetFilterView != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetFilterView mock enable")
		return r.cli.mock.mockDriveGetSheetFilterView(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSheetFilterView",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetFilterViewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetSheetFilterView(f func(ctx context.Context, request *GetSheetFilterViewReq, options ...MethodOptionFunc) (*GetSheetFilterViewResp, *Response, error)) {
	r.mockDriveGetSheetFilterView = f
}

func (r *Mock) UnMockDriveGetSheetFilterView() {
	r.mockDriveGetSheetFilterView = nil
}

type GetSheetFilterViewReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值："pH9hbVcCXA"
}

type getSheetFilterViewResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetSheetFilterViewResp `json:"data,omitempty"`
}

type GetSheetFilterViewResp struct {
	FilterView *GetSheetFilterViewRespFilterView `json:"filter_view,omitempty"` // 筛选视图信息，包括 id、name、range
}

type GetSheetFilterViewRespFilterView struct {
	FilterViewID   string `json:"filter_view_id,omitempty"`   // 筛选视图 id
	FilterViewName string `json:"filter_view_name,omitempty"` // 筛选视图名字
	Range          string `json:"range,omitempty"`            // 筛选视图的筛选范围
}
