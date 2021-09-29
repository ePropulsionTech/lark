// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteBitableField 该接口用于在数据表中删除一个字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/delete
func (r *BitableService) DeleteBitableField(ctx context.Context, request *DeleteBitableFieldReq, options ...MethodOptionFunc) (*DeleteBitableFieldResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteBitableField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableField mock enable")
		return r.cli.mock.mockBitableDeleteBitableField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "DeleteBitableField",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteBitableFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableDeleteBitableField(f func(ctx context.Context, request *DeleteBitableFieldReq, options ...MethodOptionFunc) (*DeleteBitableFieldResp, *Response, error)) {
	r.mockBitableDeleteBitableField = f
}

func (r *Mock) UnMockBitableDeleteBitableField() {
	r.mockBitableDeleteBitableField = nil
}

type DeleteBitableFieldReq struct {
	AppToken string `path:"app_token" json:"-"` // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值："tblsRc9GRRXKqhvW"
	FieldID  string `path:"field_id" json:"-"`  // field id, 示例值："fldPTb0U2y"
}

type deleteBitableFieldResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteBitableFieldResp `json:"data,omitempty"`
}

type DeleteBitableFieldResp struct {
	FieldID string `json:"field_id,omitempty"` // field id
	Deleted bool   `json:"deleted,omitempty"`  // 删除标记
}
