// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetOKRPeriodList 获取OKR周期列表
//
// 使用tenant_access_token需要额外申请权限<md-perm
// href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">以应用身份访问OKR信息</md-perm>
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/list
func (r *OKRService) GetOKRPeriodList(ctx context.Context, request *GetOKRPeriodListReq, options ...MethodOptionFunc) (*GetOKRPeriodListResp, *Response, error) {
	if r.cli.mock.mockOKRGetOKRPeriodList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] OKR#GetOKRPeriodList mock enable")
		return r.cli.mock.mockOKRGetOKRPeriodList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "GetOKRPeriodList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/periods",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getOKRPeriodListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockOKRGetOKRPeriodList(f func(ctx context.Context, request *GetOKRPeriodListReq, options ...MethodOptionFunc) (*GetOKRPeriodListResp, *Response, error)) {
	r.mockOKRGetOKRPeriodList = f
}

func (r *Mock) UnMockOKRGetOKRPeriodList() {
	r.mockOKRGetOKRPeriodList = nil
}

type GetOKRPeriodListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标志page_token, 示例值："xaasdasdax"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小，默认10, 示例值：10, 默认值: `10`
}

type getOKRPeriodListResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetOKRPeriodListResp `json:"data,omitempty"`
}

type GetOKRPeriodListResp struct {
	PageToken string                      `json:"page_token,omitempty"` // 分页标志
	HasMore   bool                        `json:"has_more,omitempty"`   // 是否有更多
	Items     []*GetOKRPeriodListRespItem `json:"items,omitempty"`      // 数据项
}

type GetOKRPeriodListRespItem struct {
	ID     string `json:"id,omitempty"`      // id
	ZhName string `json:"zh_name,omitempty"` // 中文名称
	EnName string `json:"en_name,omitempty"` // 英文名称
	Status int64  `json:"status,omitempty"`  // 启用状态, 可选值有: `0`：正常状态, `1`：暂不处理, `2`：标记失效, `3`：隐藏周期
}
