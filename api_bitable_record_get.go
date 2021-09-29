// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBitableRecord 该接口用于根据 record_id 的值检索现有记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-record/get
func (r *BitableService) GetBitableRecord(ctx context.Context, request *GetBitableRecordReq, options ...MethodOptionFunc) (*GetBitableRecordResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableRecord != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableRecord mock enable")
		return r.cli.mock.mockBitableGetBitableRecord(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableRecord",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableRecordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableGetBitableRecord(f func(ctx context.Context, request *GetBitableRecordReq, options ...MethodOptionFunc) (*GetBitableRecordResp, *Response, error)) {
	r.mockBitableGetBitableRecord = f
}

func (r *Mock) UnMockBitableGetBitableRecord() {
	r.mockBitableGetBitableRecord = nil
}

type GetBitableRecordReq struct {
	AppToken string `path:"app_token" json:"-"` // bitable app token, 示例值："bascnCMII2ORej2RItqpZZUNMIe"
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值："tblxI2tWaxP5dG7p"
	RecordID string `path:"record_id" json:"-"` // 单条记录的 id, 示例值："recn0hoyXL"
}

type getBitableRecordResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableRecordResp `json:"data,omitempty"`
}

type GetBitableRecordResp struct {
	Record *GetBitableRecordRespRecord `json:"record,omitempty"` // 记录
}

type GetBitableRecordRespRecord struct {
	RecordID string                 `json:"record_id,omitempty"` // 记录 id
	Fields   map[string]interface{} `json:"fields,omitempty"`    // 记录字段
}
