// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteSearchDataSourceItem 删除数据记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/delete
func (r *SearchService) DeleteSearchDataSourceItem(ctx context.Context, request *DeleteSearchDataSourceItemReq, options ...MethodOptionFunc) (*DeleteSearchDataSourceItemResp, *Response, error) {
	if r.cli.mock.mockSearchDeleteSearchDataSourceItem != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#DeleteSearchDataSourceItem mock enable")
		return r.cli.mock.mockSearchDeleteSearchDataSourceItem(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "DeleteSearchDataSourceItem",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id/items/:item_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteSearchDataSourceItemResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockSearchDeleteSearchDataSourceItem(f func(ctx context.Context, request *DeleteSearchDataSourceItemReq, options ...MethodOptionFunc) (*DeleteSearchDataSourceItemResp, *Response, error)) {
	r.mockSearchDeleteSearchDataSourceItem = f
}

func (r *Mock) UnMockSearchDeleteSearchDataSourceItem() {
	r.mockSearchDeleteSearchDataSourceItem = nil
}

type DeleteSearchDataSourceItemReq struct {
	DataSourceID string `path:"data_source_id" json:"-"` // 数据源的ID, 示例值："service_ticket"
	ItemID       string `path:"item_id" json:"-"`        // 数据记录的ID, 示例值："01010111"
}

type deleteSearchDataSourceItemResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteSearchDataSourceItemResp `json:"data,omitempty"`
}

type DeleteSearchDataSourceItemResp struct{}
