// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetSearchDataSourceList 获取创建的所有数据源信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/list
func (r *SearchService) GetSearchDataSourceList(ctx context.Context, request *GetSearchDataSourceListReq, options ...MethodOptionFunc) (*GetSearchDataSourceListResp, *Response, error) {
	if r.cli.mock.mockSearchGetSearchDataSourceList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#GetSearchDataSourceList mock enable")
		return r.cli.mock.mockSearchGetSearchDataSourceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "GetSearchDataSourceList",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/search/v2/data_sources",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getSearchDataSourceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockSearchGetSearchDataSourceList(f func(ctx context.Context, request *GetSearchDataSourceListReq, options ...MethodOptionFunc) (*GetSearchDataSourceListResp, *Response, error)) {
	r.mockSearchGetSearchDataSourceList = f
}

func (r *Mock) UnMockSearchGetSearchDataSourceList() {
	r.mockSearchGetSearchDataSourceList = nil
}

type GetSearchDataSourceListReq struct {
	View      *int64  `query:"view" json:"-"`       // 回包数据格式，0-全量数据；1-摘要数据。,**注**：摘要数据仅包含"id"，"name"，"state"。, 示例值：0, 可选值有: `0`：全量数据, `1`：摘要数据
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："PxZFma9OIRhdBlT/dOYNiu2Ro8F2WAhcby7OhOijfljZ"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`50`
}

type getSearchDataSourceListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetSearchDataSourceListResp `json:"data,omitempty"`
}

type GetSearchDataSourceListResp struct {
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                             `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetSearchDataSourceListRespItem `json:"items,omitempty"`      // 数据源中的数据记录
}

type GetSearchDataSourceListRespItem struct {
	ID            string `json:"id,omitempty"`              // 数据源的唯一标识
	Name          string `json:"name,omitempty"`            // data_source的展示名称
	State         int64  `json:"state,omitempty"`           // 数据源状态，0-未上线，1-已上线, 可选值有: `0`：未上线, `1`：已上线
	Description   string `json:"description,omitempty"`     // 对于数据源的描述
	CreateTime    string `json:"create_time,omitempty"`     // 创建时间，使用Unix时间戳，单位为“秒”
	UpdateTime    string `json:"update_time,omitempty"`     // 更新时间，使用Unix时间戳，单位为“秒”
	IsExceedQuota bool   `json:"is_exceed_quota,omitempty"` // 是否超限
}
